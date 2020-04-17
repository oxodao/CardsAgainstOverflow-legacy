package model

const CommandCriticalError string = "CRITICAL_ERROR"
const CommandError string = "ERROR"
const CommandConnected string = "CONNECTED"
const CommandDisconnected string = "DISCONNECTED"
const CommandPlayerList string = "PLAYER_LIST"
const CommandStarted string = "GAME_STARTED"
const CommandUpdateCards string = "UPDATE_CARDS"
const CommandSendVoteCard string = "VOTING_HAND"

type Command struct {
	Command   string `json:"Command"`
	Arguments string `json:"Arguments"`
}
