package model

const CommandError string = "ERROR"
const CommandConnected string = "CONNECTED"

type Command struct {
	Command   string `json:"Command"`
	Arguments string `json:"Arguments"`
}
