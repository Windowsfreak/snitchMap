package bot

// Settings of client
type Settings struct {
	Locale             string //地区
	ViewDistance       int    //视距
	ChatMode           int    //聊天模式
	ChatColors         bool   //聊天颜色
	DisplayedSkinParts uint8  //皮肤显示
	MainHand           int    //主手
	ReceiveMap         bool   //接收地图数据
	Brand              string // The brand string presented to the server.
}

/*
	Used by Settings.DisplayedSkinParts.
	For each bits set if shows match part.
*/
const (
	_ = 1 << iota
	Jacket
	LeftSleeve
	RightSleeve
	LeftPantsLeg
	RightPantsLeg
	Hat
)

//DefaultSettings are the default settings of client
var DefaultSettings = Settings{
	Locale:             "de_DE", // ^_^
	ViewDistance:       0,
	ChatMode:           0,
	DisplayedSkinParts: Jacket | LeftSleeve | RightSleeve | LeftPantsLeg | RightPantsLeg | Hat,
	MainHand:           1,
	ReceiveMap:         false,
	Brand:              "snitchBot EXPERIMENTAL by Windowsfreak",
}
