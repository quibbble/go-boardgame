package bgerr

// Status represents common board game status errors
type Status int

const (
	StatusTooFewTeams = iota
	StatusTooManyTeams
	StatusInvalidOption
	StatusUnknownTeam
	StatusUnknownActionType
	StatusInvalidActionDetails
	StatusInvalidAction
	StatusWrongTurn
	StatusGameOver
	StatusBGNDecodingFailure
	StatusBGNEncodingFailure
)

var statusText = map[int]string{
	StatusTooFewTeams:          "Too Few Teams",
	StatusTooManyTeams:         "Too Many Teams",
	StatusInvalidOption:        "Invalid Option",
	StatusUnknownTeam:          "Unknown Team",
	StatusUnknownActionType:    "Unknown Action Type",
	StatusInvalidActionDetails: "Invalid Action Details",
	StatusInvalidAction:        "Invalid Action",
	StatusWrongTurn:            "Wrong Turn",
	StatusGameOver:             "Game Over",
	StatusBGNDecodingFailure:   "BGN Decoding Failure",
	StatusBGNEncodingFailure:   "BGN Encoding Failure",
}

// StatusText returns the text for a board game status code
// returns empty string if the code is invalid
func StatusText(code int) string {
	return statusText[code]
}
