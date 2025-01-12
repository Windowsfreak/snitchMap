// This file is automatically generated by gen_packetIDs.go. DO NOT EDIT.

package data

//go:generate go run gen_packetIDs.go

// PktID represents a packet ID used in the minecraft protocol.
type PktID int32

// Valid PktID values.
const (
	// Clientbound packets for connections in the login state.
	EncryptionBeginClientbound PktID = 0x01
	Success                    PktID = 0x02
	Compress                   PktID = 0x03
	LoginPluginRequest         PktID = 0x04
	Disconnect                 PktID = 0x00
	// Serverbound packets for connections in the login state.
	LoginStart                 PktID = 0x00
	EncryptionBeginServerbound PktID = 0x01
	LoginPluginResponse        PktID = 0x02

	// Clientbound packets for connections in the play state.
	KickDisconnect             PktID = 0x1a
	UnlockRecipes              PktID = 0x31
	Animation                  PktID = 0x06
	WorldEvent                 PktID = 0x21
	ScoreboardDisplayObjective PktID = 0x3b
	CustomPayloadClientbound   PktID = 0x18
	NamedSoundEffect           PktID = 0x19
	HeldItemSlotClientbound    PktID = 0x3a
	ChatClientbound            PktID = 0x0f
	AbilitiesClientbound       PktID = 0x2c
	EntityDestroy              PktID = 0x32
	SetPassengers              PktID = 0x43
	KeepAliveClientbound       PktID = 0x1f
	SpawnEntityExperienceOrb   PktID = 0x01
	OpenHorseWindow            PktID = 0x1e
	RemoveEntityEffect         PktID = 0x33
	RelEntityMove              PktID = 0x26
	SelectAdvancementTab       PktID = 0x37
	OpenSignEntity             PktID = 0x2a
	Map                        PktID = 0x24
	FacePlayer                 PktID = 0x33
	EntityEquipment            PktID = 0x3f
	ResourcePackSend           PktID = 0x34
	NbtQueryResponse           PktID = 0x54
	ScoreboardObjective        PktID = 0x42
	StopSound                  PktID = 0x52
	OpenWindow                 PktID = 0x13
	Camera                     PktID = 0x39
	Advancements               PktID = 0x4d
	UpdateTime                 PktID = 0x47
	Login                      PktID = 0x23
	PositionClientbound        PktID = 0x2F
	UpdateViewPosition         PktID = 0x2f
	EntitySoundEffect          PktID = 0x50
	Respawn                    PktID = 0x35
	BlockChange                PktID = 0x0b
	BlockBreakAnimation        PktID = 0x08
	Title                      PktID = 0x48
	EntityTeleport             PktID = 0x4c
	EntityEffect               PktID = 0x4f
	TileEntityData             PktID = 0x09
	SpawnPosition              PktID = 0x46
	WorldBorder                PktID = 0x38
	Experience                 PktID = 0x40
	PlayerlistHeader           PktID = 0x4a
	PlayerlistItem             PktID = 0x2e
	WindowItems                PktID = 0x14
	EntityUpdateAttributes     PktID = 0x4e
	EntityHeadRotation         PktID = 0x36
	VehicleMoveClientbound     PktID = 0x29
	MapChunk                   PktID = 0x20
	EntityLook                 PktID = 0x28
	Teams                      PktID = 0x44
	UpdateViewDistance         PktID = 0x41
	Explosion                  PktID = 0x1c
	MultiBlockChange           PktID = 0x10
	PlayerInfo                 PktID = 0x32
	CraftRecipeResponse        PktID = 0x2b
	TransactionClientbound     PktID = 0x11
	TradeList                  PktID = 0x26
	CloseWindowClientbound     PktID = 0x12
	TabCompleteClientbound     PktID = 0x0e
	SetCooldown                PktID = 0x17
	BlockAction                PktID = 0x0a
	NamedEntitySpawn           PktID = 0x05
	SpawnEntityPainting        PktID = 0x04
	UpdateLight                PktID = 0xFE
	CombatEvent                PktID = 0x2d
	SpawnEntityLiving          PktID = 0x03
	ScoreboardScore            PktID = 0x45
	DeclareCommands            PktID = 0x10
	UpdateHealth               PktID = 0x41
	EntityMetadata             PktID = 0x3c
	AttachEntity               PktID = 0x3d
	Tags                       PktID = 0x5b
	EntityStatus               PktID = 0x1b
	AcknowledgePlayerDigging   PktID = 0x14
	Collect                    PktID = 0x4b
	WorldParticles             PktID = 0x22
	Entity                     PktID = 0x2a
	UnloadChunk                PktID = 0x1d
	Difficulty                 PktID = 0x0d
	CraftProgressBar           PktID = 0x14
	BossBar                    PktID = 0x0c
	DeclareRecipes             PktID = 0x5a
	GameStateChange            PktID = 0x1e
	Statistics                 PktID = 0x07
	EntityVelocity             PktID = 0x3e
	SetSlot                    PktID = 0x16
	OpenBook                   PktID = 0x2c
	SoundEffect                PktID = 0x49
	EntityMoveLook             PktID = 0x27
	SpawnEntity                PktID = 0x00
	// Serverbound packets for connections in the play state.
	EnchantItem                PktID = 0x06
	CustomPayloadServerbound   PktID = 0x09
	SelectTrade                PktID = 0x23
	SetCreativeSlot            PktID = 0x1b
	UpdateSign                 PktID = 0x1c
	WindowClick                PktID = 0x07
	PositionLook               PktID = 0x0e
	UpdateCommandBlock         PktID = 0x26
	QueryBlockNbt              PktID = 0x01
	Flying                     PktID = 0x0c
	KeepAliveServerbound       PktID = 0x0b
	ClientCommand              PktID = 0x03
	BlockPlace                 PktID = 0x1f
	EntityAction               PktID = 0x15
	PositionServerbound        PktID = 0x0d
	ResourcePackReceive        PktID = 0x18
	Spectate                   PktID = 0x1e
	TeleportConfirm            PktID = 0x00
	GenerateStructure          PktID = 0x0f
	SetDifficulty              PktID = 0x02
	CloseWindowServerbound     PktID = 0x08
	Look                       PktID = 0x0f
	AdvancementTab             PktID = 0x19
	SetBeaconEffect            PktID = 0x24
	AbilitiesServerbound       PktID = 0x13
	ChatServerbound            PktID = 0x02
	DisplayedRecipe            PktID = 0x1e
	RecipeBook                 PktID = 0x17
	UpdateJigsawBlock          PktID = 0x29
	TransactionServerbound     PktID = 0x05
	SteerVehicle               PktID = 0x16
	NameItem                   PktID = 0x20
	PickItem                   PktID = 0x18
	UpdateStructureBlock       PktID = 0x2a
	TabCompleteServerbound     PktID = 0x01
	HeldItemSlotServerbound    PktID = 0x1a
	SteerBoat                  PktID = 0x11
	Settings                   PktID = 0x04
	UseItem                    PktID = 0x20
	CraftRecipeRequest         PktID = 0x12
	UpdateCommandBlockMinecart PktID = 0x27
	BlockDig                   PktID = 0x14
	EditBook                   PktID = 0x0c
	UseEntity                  PktID = 0x0a
	VehicleMoveServerbound     PktID = 0x10
	ArmAnimation               PktID = 0x1d
	LockDifficulty             PktID = 0x11
	QueryEntityNbt             PktID = 0x0d

	// Clientbound packets used to respond to ping/status requests.
	ServerInfo      PktID = 0x00
	PingClientbound PktID = 0x01
	// Serverbound packets used to ping or read server status.
	PingStart       PktID = 0x00
	PingServerbound PktID = 0x01
)
