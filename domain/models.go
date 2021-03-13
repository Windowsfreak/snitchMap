package domain

type Error struct {
	ErrorMessage string `json:"errorMessage"`
}

type PreSharedKeyRequest struct {
	PreSharedKey string
}

type GetEventsAfterRequest struct {
	PreSharedKeyRequest
	Rowid int64
	Limit int64
}

type GetEventsByUserAfterRequest struct {
	GetEventsAfterRequest
	Username string
}

type GetByRegionRequest struct {
	X1 int64
	Z1 int64
	X2 int64
	Z2 int64
}

type GetEventsByRegionAfterRequest struct {
	GetEventsAfterRequest
	GetByRegionRequest
}

type SetSnitchAlertByRegionRequest struct {
	PreSharedKeyRequest
	GetByRegionRequest
	Alert bool
}

type GetByTimeRequest struct {
	PreSharedKeyRequest
	Time int64
}

type GetByUsernameRequest struct {
	PreSharedKeyRequest
	Username string
}

type GetByTimeAndUsernameRequest struct {
	GetByTimeRequest
	Username string
}

type GetByTimeAndRegionRequest struct {
	GetByTimeRequest
	GetByRegionRequest
}

type User struct {
	User     string
	Login    int64
	Logout   int64
	Seen     int64
	X        int64
	Y        int64
	Z        int64
	Angle    float32
	Speed    float32
	Interval int64
	Hits     int64
}

type Event struct {
	Rowid int64
	X     int64
	Y     int64
	Z     int64
	User  string
	Event string
	Time  int64
	Nl    string
	Name  string
	Alert bool
}

type Snitch struct {
	X     int64
	Y     int64
	Z     int64
	Nl    string
	Name  string
	Seen  int64
	Read  int64
	Dead  bool
	Hits  int64
	Cull  float32
	Alert bool
}

type Chat struct {
	Rowid int64
	Time  int64
	Nl    string
	User  string
	Text  string
}

type ConfigStruct struct {
	ServerHost        string `yaml:"ServerHost"`
	ServerPort        int    `yaml:"ServerPort"`
	ServerAddr        string `yaml:"ServerAddr"`
	PreSharedKey      string `yaml:"PreSharedKey"`
	DiscordToken      string `yaml:"DiscordToken"`
	Username          string `yaml:"Username"`
	Password          string `yaml:"Password"`
	DebugChannel      string `yaml:"DebugChannel"`
	GlobalChatChannel string `yaml:"GlobalChatChannel"`
	SnitchChannel     string `yaml:"SnitchChannel"`
	GeneralChannel    string `yaml:"GeneralChannel"`
	Https             bool   `yaml:"Https"`
}

type UserResponse struct {
	Time  int64
	Users []User
}
