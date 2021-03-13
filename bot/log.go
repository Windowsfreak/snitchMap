package bot

import (
	"encoding/hex"
	"github.com/Lukaesebrot/mojango"
	"github.com/Windowsfreak/go-mc/bot/path"
	. "github.com/Windowsfreak/go-mc/bot/repository"
	"github.com/Windowsfreak/go-mc/chat"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const ENTER = "entered"        // * PDcolt2 entered snitch at  [world -4104 90 10665]
const LOGOUT = "logged out in" // * skraeyling logged out in snitch at  [world -4087 90 10692]
const LOGIN = "logged in at"   // * lasoogna logged in to snitch at  [world 4145 2 2435]
// "§f[2024 93 4941]       0.00      MLS "
// "§f[948 223 2940]       0.00      nvs1 skycurtain"
// " * Page 289 is empty"

var discoveryBegin = time.Time{}
var mojangoClient = mojango.New()
var repository Repository

func InitializeRepository() {
	var err error
	repository, err = NewRepository()
	if err != nil {
		log.Fatal(err)
	}
}

func StartSnitchDiscovery() {
	discoveryBegin = time.Now()
	go func(myTime time.Time) {
		time.Sleep(time.Second * 1200)
		EndSnitchDiscovery(myTime.Add(-20 * time.Second))
	}(discoveryBegin)
}
func DiscoverSnitch(msg chat.Message) *SnitchEvent {
	if strings.Contains(msg.ClearString(), "snitch at") && strings.HasPrefix(msg.ClearString(), " * ") {
		return HandleHitSnitch(msg)
	}
	return nil
}
func EndSnitchDiscovery(time time.Time) error {
	err := repository.KillSnitchesOlderThan(time)
	if err != nil {
		log.Println(err)
	}
	return err
}

var regexpNew = regexp.MustCompile(`^([^ ]+) is brand new!$`)
var regexpChat = regexp.MustCompile(`^\[([^\]]+)\] ([^ :]+): (.*)$`)
var regexpHitSnitch = regexp.MustCompile(`^§b \* (?P<Person>\S*) (?P<Action>(?:entered|logged in to|logged out in)) snitch at (?P<Name>\S*) \[(?P<World>\S+) (?P<X>\S+) (?P<Y>\S+) (?P<Z>\S+)\]$`)
var regexpPage = regexp.MustCompile(`^§f\[(?P<X>\S+) (?P<Y>\S+) (?P<Z>\S+)\]\s+(?P<Cull>\d+\.\d*)\s+(?P<Group>\S+)\s(?P<Name>\S*)\n$`)
var regexpEmptyPage = regexp.MustCompile(`^ \* Page \d+ is empty$`)
var NextPage = 1
var LastJaMessage time.Time

func HandleSnitchPageMsg(msg chat.Message, c *Client) (lastPage bool) {
	isSnitchPage := false
	//println("Msg Text \"" + msg.Text + "\"")
	for i := range msg.Extra {
		if strings.HasPrefix(msg.Extra[i].Text, "§f[") {
			if HandleSnitchPageEntry(msg.Extra[i]) {
				isSnitchPage = true
			}
		} else if regexpEmptyPage.MatchString(msg.Extra[i].Text) {
			log.Println("End snitch discovery")
			isSnitchPage = false
			EndSnitchDiscovery(discoveryBegin.Add(-20 * time.Second))
			lastPage = true
			break
		} else {
			//println("Msg \"" + msg.Extra[i].Text + "\"")
		}
	}
	if isSnitchPage {
		LastJaMessage = time.Now()
		go func() {
			time.Sleep(time.Second * 3)
			NextPage = NextPage + 1
			err := c.Chat("/jalistlong " + strconv.Itoa(NextPage))
			if err != nil {
				log.Println(err)
			}
		}()
	}
	return
}
func RecordNew(m chat.Message) (username string) {
	info := regexpNew.FindStringSubmatch(m.ClearString())
	if info == nil || len(info) < 1 {
		log.Println("Msg No Match \"" + m.ClearString() + "\"")
		return
	}
	err := repository.StoreChat("", info[1], "is brand new!")
	if err != nil {
		log.Println(err)
	}
	username = info[1]
	return
}
func RecordChat(m chat.Message) (group string, username string, message string) {
	info := regexpChat.FindStringSubmatch(m.ClearString())
	if info == nil || len(info) < 3 {
		log.Println("Msg No Match \"" + m.ClearString() + "\"")
		return
	}
	group = info[1]
	username = info[2]
	message = info[3]
	err := repository.StoreChat(info[1], info[2], info[3])
	if err != nil {
		log.Println(err)
	}
	return
}
func RecordPlayerLoggingIn(uuid uuid.UUID, name string) *SnitchPlayer {
	player, err := repository.FetchUserByName(name)
	if err != nil {
	}
	if player == nil {
		player = &SnitchPlayer{}
	}
	player.UUID = uuid
	player.Name = name
	player.Login = time.Now()
	err = repository.StoreUser(player)
	if err != nil {
		log.Println(err)
	}
	err = repository.StoreChat("", player.Name, "joined")
	if err != nil {
		log.Println(err)
	}
	return player
}
func RecordPlayerLoggingOut(uuid uuid.UUID) *SnitchPlayer {
	name, err := repository.GetName(uuid)
	if err == nil {
		player, err := repository.FetchUserByName(name)
		if err != nil {
		}
		if player == nil {
			player = &SnitchPlayer{}
		}
		player.UUID = uuid
		player.Logout = time.Now()
		err = repository.StoreUser(player)
		if err != nil {
			log.Println(err)
		}
		err = repository.StoreChat("", player.Name, "left")
		if err != nil {
			log.Println(err)
		}
		return player
	}
	return nil
}
func HandleSnitchPageEntry(msg chat.Message) bool {
	info := regexpPage.FindStringSubmatch(msg.Text)
	if info == nil || len(info) < 5 {
		return false
	}
	x, _ := strconv.Atoi(info[1])
	y, _ := strconv.Atoi(info[2])
	z, _ := strconv.Atoi(info[3])
	cull := info[4]
	group := info[5]
	name := info[6]
	now := time.Now()
	v3 := path.V3{X: x, Y: y, Z: z}
	snitch, err := repository.FetchSnitch(&v3)
	if err != nil {
	}
	if snitch == nil {
		snitch = &Snitch{}
	}
	snitch.Name = name
	snitch.Group = group
	if snitch.V3 == nil {
		snitch.V3 = &v3
	}
	snitch.Read = now
	snitch.Dead = false
	cullf, err := strconv.ParseFloat(cull, 32)
	snitch.Cull = float32(cullf)
	err = repository.StoreSnitch(snitch)
	if err != nil {
		log.Println(err)
	}
	return true
}
func HandleHitSnitch(msg chat.Message) (snitchEvent *SnitchEvent) {
	info := regexpHitSnitch.FindStringSubmatch(msg.Text)
	if info == nil || len(info) < 7 {
		log.Println("Msg No Match \""+msg.Text+"\"", hex.Dump([]byte(msg.Text)))
		return
	}
	person := info[1]
	action := info[2]
	name := info[3]
	// world := info[4]
	now := time.Now()
	x, _ := strconv.Atoi(info[5])
	y, _ := strconv.Atoi(info[6])
	z, _ := strconv.Atoi(info[7])
	v3 := path.V3{X: x, Y: y, Z: z}
	snitch, err := repository.FetchSnitch(&v3)
	if err != nil {
	}
	if snitch == nil {
		snitch = &Snitch{}
	}
	snitch.Name = name
	if snitch.V3 == nil {
		snitch.V3 = &v3
	}
	snitch.Seen = now
	snitch.Dead = false
	snitch.Hits += 1
	err = repository.StoreSnitch(snitch)
	if err != nil {
		log.Println(err)
	}
	player, err := repository.FetchUserByName(person)
	if err != nil {
	}
	if player == nil {
		player = &SnitchPlayer{}
		id, err := mojangoClient.FetchUUID(person)
		if err == nil {
			player.UUID, err = uuid.Parse(id)
		} else {
			log.Println(err)
		}
	}
	player.Name = person
	player.Interval = now.Sub(player.Seen)
	player.PrevV3 = player.V3
	player.V3 = snitch.V3
	player.Seen = now
	player.Hits += 1
	err = repository.StoreUser(player)
	if err != nil {
		log.Println(err)
	}
	snitchEvent = &SnitchEvent{
		Person: player,
		Snitch: snitch,
		Action: action,
		Time:   now,
	}
	err = repository.StoreEvent(snitchEvent)
	if err != nil {
		log.Println(err)
	}
	return
}
