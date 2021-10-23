package bgerr

// Status represents common board game status errors
type Status int

const (
	StatusTooFewTeams = iota
	StatusTooManyTeams
	StatusInvalidOption
	StatusInvalidTeam
	StatusInvalidActionType
	StatusInvalidActionDetails
	StatusWrongTurn
	StatusGameOver
)

var statusText = map[int]string{
	StatusTooFewTeams:          "Too Few Teams",
	StatusTooManyTeams:         "Too Many Teams",
	StatusInvalidOption:        "Invalid Option",
	StatusInvalidTeam:          "Invalid Team",
	StatusInvalidActionType:    "Invalid Game Action Type",
	StatusInvalidActionDetails: "Invalid Action Details",
	StatusWrongTurn:            "Wrong Turn",
	StatusGameOver:             "Game Over",
}

// StatusText returns the text for a board game status code
// returns empty string if the code is invalid
func StatusText(code int) string {
	return statusText[code]
}
