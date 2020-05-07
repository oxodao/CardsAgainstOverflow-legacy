package model

const CommandCriticalError string = "CRITICAL_ERROR"
const CommandError string = "ERROR"
const CommandConnected string = "CONNECTED"
const CommandDisconnected string = "DISCONNECTED"
const CommandPlayerList string = "PLAYER_LIST"
const CommandSetGamestate string = "SET_GAMESTATE"
const CommandCountdown string = "COUNTDOWN"
const CommandJudgeSelection string = "JUDGE_SELECTION"
const CommandGotSettings string = "GOT_SETTINGS"

const CommandReroll string = "REROLL"

const CommandWizzRefilled string = "WIZZ_REFILLED"

const CommandHasPlayed string = "HAS_PLAYED"

type Command struct {
	Command   string
	Arguments string
}
