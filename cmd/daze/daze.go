package main

import (
	"bytes"
	"github.com/Windowsfreak/go-mc/bot/repository"
	"github.com/Windowsfreak/go-mc/domain"
	"github.com/Windowsfreak/go-mc/yggdrasil"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Windowsfreak/go-mc/bot"
	"github.com/Windowsfreak/go-mc/chat"
	_ "github.com/Windowsfreak/go-mc/data/lang/en-us"
	pk "github.com/Windowsfreak/go-mc/net/packet"
	"github.com/bwmarrin/discordgo"
)

var c *bot.Client
var dg *discordgo.Session

var shouldStayconnected = true
var messages = make(chan msg, 131072)

type msg struct {
	ChannelId string
	Message   string
}

func main() {
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&domain.Config)
	if err != nil {
		log.Fatal(err)
	}

	dg, err = discordgo.New("Bot " + domain.Config.DiscordToken)
	if err != nil {
		log.Println("error creating Discord session,", err)
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuilds + discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Println("error opening Discord connection,", err)
	}

	go func() {
		for shouldStayconnected {
			msg := <-messages
			_, err := dg.ChannelMessageSend(msg.ChannelId, msg.Message)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}()

	bot.InitializeRepository()
	var errors = 0
	for shouldStayconnected {
		if errors > 10 {
			messages <- msg{domain.Config.DebugChannel, EscapeDiscordMsg("Too many errors! If you want me to reconnect, please type stayconnected")}
			shouldStayconnected = false
			for !shouldStayconnected {
				time.Sleep(time.Second * 5)
			}
			messages <- msg{domain.Config.DebugChannel, EscapeDiscordMsg("Convinced. I'm back!")}
			errors = 0
		}
		authenticate, err := yggdrasil.Authenticate(domain.Config.Username, domain.Config.Password)
		if err != nil {
			log.Println(err.Error())
			messages <- msg{domain.Config.DebugChannel, EscapeDiscordMsg("Auth Servers don't seem to like me. Retrying...")}
			time.Sleep(time.Second * 15)
			errors++
			continue
		}

		c = bot.NewClient()

		c.Auth.UUID, c.Name = authenticate.SelectedProfile()
		c.AsTk = authenticate.AccessToken()

		//Login
		err = c.JoinServer(domain.Config.ServerHost, domain.Config.ServerPort)
		if err != nil {
			log.Println(err.Error())
			messages <- msg{domain.Config.DebugChannel, EscapeDiscordMsg("Server not available. Retrying...")}
			time.Sleep(time.Second * 10)
			errors++
			continue
		}
		log.Println("Login success")

		//Register event handlers
		c.Events.GameStart = onGameStart
		c.Events.ChatMsg = onChatMsg
		c.Events.Disconnect = onDisconnect
		c.Events.PluginMessage = onPluginMessage
		c.Events.PlayerJoin = onPlayerJoin
		c.Events.PlayerLeave = onPlayerLeave
		c.Events.Die = onDeath

		//JoinGame
		err = c.HandleGame()
		errors = 0
		if err != nil {
			log.Println(err.Error())
		}
		if shouldStayconnected {
			messages <- msg{domain.Config.DebugChannel, EscapeDiscordMsg("Disconnected. Reconnecting...")}
			log.Println("Reconnecting in 5 seconds...")
			time.Sleep(time.Second * 5)
		}
	}
}

func onPlayerJoin(snitchPlayer *repository.SnitchPlayer) error {
	if snitchPlayer == nil {
		snitchPlayer = &repository.SnitchPlayer{
			Name: "(unknown)",
		}
	}
	messages <- msg{domain.Config.GlobalChatChannel, EscapeDiscordMsg(snitchPlayer.Name + " joined")}
	return nil
}

func onPlayerLeave(snitchPlayer *repository.SnitchPlayer) error {
	if snitchPlayer == nil {
		snitchPlayer = &repository.SnitchPlayer{
			Name: "(unknown)",
		}
	}
	messages <- msg{domain.Config.GlobalChatChannel, EscapeDiscordMsg(snitchPlayer.Name + " left")}
	return nil
}

func onGameStart() error {
	log.Println("Game start")
	messages <- msg{domain.Config.DebugChannel, EscapeDiscordMsg("SnitchMap bot: Online")}
	return nil //if err isn't nil, HandleGame() will return it.
}

func onChatMsg(m chat.Message, pos byte) error {
	if pos == 0 && len(m.Extra) > 0 && strings.Contains(m.ClearString(), "Snitch List for") {
		if bot.HandleSnitchPageMsg(m, c) {
			messages <- msg{domain.Config.DebugChannel, EscapeDiscordMsg("I'm done updating all the snitches!")}
		}
	} else if pos == 1 && strings.HasPrefix(m.ClearString(), "[") {
		group, username, message := bot.RecordChat(m)
		if username != "" {
			messages <- msg{domain.Config.GlobalChatChannel, EscapeDiscordMsg("[" + group + "] " + username + ": " + message)}
		}
	} else if pos == 1 && strings.Contains(m.ClearString(), "is brand new!") {
		username := bot.RecordNew(m)
		messages <- msg{domain.Config.GlobalChatChannel, EscapeDiscordMsg(username + " is brand new!")}
	} else if pos == 0 && len(m.Text) > 0 {
		snitchEvent := bot.DiscoverSnitch(m)
		text := ""
		if snitchEvent != nil && snitchEvent.Snitch != nil && snitchEvent.Snitch.Alert == true {
			text += " @everyone"
		}
		if snitchEvent != nil && snitchEvent.Snitch != nil {
			messages <- msg{domain.Config.SnitchChannel, EscapeDiscordMsg("["+snitchEvent.Snitch.Group+"] "+
				snitchEvent.Person.Name+" "+snitchEvent.Action+
				" ["+strconv.Itoa(snitchEvent.Snitch.V3.X)+
				" "+strconv.Itoa(snitchEvent.Snitch.V3.Y)+
				" "+strconv.Itoa(snitchEvent.Snitch.V3.Z)+"] "+snitchEvent.Snitch.Name) + text}
		}
	}
	log.Println(m.String()) // output chat message without any format code (like color or bold)
	return nil
}

func onDeath() error {
	log.Println("Died and Respawned")
	c.Respawn() // If we exclude Respawn(...) then the player won't press the "Respawn" button upon death
	messages <- msg{domain.Config.DebugChannel, EscapeDiscordMsg("SnitchMap bot: Reviving a dead player...")}
	return nil
}

func onDisconnect(c chat.Message) error {
	log.Println("Disconnect:", c)
	messages <- msg{domain.Config.DebugChannel, EscapeDiscordMsg("SnitchMap bot: Offline")}
	return nil
}

func onPluginMessage(channel string, data []byte) error {
	switch channel {
	case "minecraft:brand":
		var brand pk.String
		if err := brand.Decode(bytes.NewReader(data)); err != nil {
			return err
		}
		log.Println("Server brand is:", brand)

	default:
		log.Println("PluginMessage", channel, string(data))
	}
	return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "shutdown" {
		if m.ChannelID != domain.Config.DebugChannel {
			c, err := s.State.Channel(m.ChannelID)
			if err != nil {
				messages <- msg{m.ChannelID, "Hey " + m.Author.Mention() + ", I'm configured not to accept this command from this channel!"}
			} else {
				messages <- msg{m.ChannelID, "Hey " + m.Author.Mention() + ", I'm configured not to accept this command from channel " + c.Mention() + "!"}
			}
			return
		}
		messages <- msg{domain.Config.DebugChannel, "SnitchMap bot: Offline"}
		shouldStayconnected = false
		log.Fatal("discord wants me to go")
	}

	if m.Content == "stayconnected" {
		if m.ChannelID != domain.Config.DebugChannel {
			c, err := s.State.Channel(m.ChannelID)
			if err != nil {
				messages <- msg{m.ChannelID, "Hey " + m.Author.Mention() + ", I'm configured not to accept this command from this channel!"}
			} else {
				messages <- msg{m.ChannelID, "Hey " + m.Author.Mention() + ", I'm configured not to accept this command from channel " + c.Mention() + "!"}
			}
			return
		}
		if shouldStayconnected {
			messages <- msg{domain.Config.DebugChannel, "Hey " + m.Author.Mention() + ", I'm already online!"}
		} else {
			messages <- msg{domain.Config.DebugChannel, "Get ready!"}
		}
		shouldStayconnected = true
	}

	if m.Content == "jalistlong" {
		if m.ChannelID != domain.Config.DebugChannel {
			c, err := s.State.Channel(m.ChannelID)
			if err != nil {
				messages <- msg{m.ChannelID, "Hey " + m.Author.Mention() + ", I'm configured not to accept this command from this channel!"}
			} else {
				messages <- msg{m.ChannelID, "Hey " + m.Author.Mention() + ", I'm configured not to accept this command from channel " + c.Mention() + "!"}
			}
			if err != nil {
				log.Println(err.Error())
			}
			return
		}
		if bot.LastJaMessage.After(time.Now().Add(-20 * time.Second)) {
			messages <- msg{m.ChannelID, "Hey " + m.Author.Mention() + ", I'm still updating the snitches. Please be patient before doing it again!"}
			return
		}
		bot.NextPage = 1
		bot.StartSnitchDiscovery()
		err := c.Chat("/jalistlong")
		if err != nil {
			log.Println(err.Error())
		}
		messages <- msg{m.ChannelID, "Hey " + m.Author.Mention() + ", I'm now updating the snitches. Please be patient!"}
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		messages <- msg{m.ChannelID, "Pong!"}
	}
}
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"", "`", "\\`", "@", "\\@​\u200b", "#", "\\#​\u200b")

func EscapeDiscordMsg(message string) string {
	return quoteEscaper.Replace(message)
}
