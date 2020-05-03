package model

const CommandCriticalError string = "CRITICAL_ERROR"
const CommandError string = "ERROR"
const CommandConnected string = "CONNECTED"
const CommandDisconnected string = "DISCONNECTED"
const CommandPlayerList string = "PLAYER_LIST"
const CommandSetGamestate string = "SET_GAMESTATE"
const CommandCountdown string = "COUNTDOWN"

type Command struct {
	Command   string
	Arguments string
}
