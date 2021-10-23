package bgerr

// Status represents common board game status errors
type Status int

const (
	StatusToFewTeams = iota
	StatusToManyTeams
	StatusInvalidOption
	StatusInvalidTeam
	StatusInvalidActionType
	StatusInvalidActionDetails
	StatusWrongTurn
	StatusGameOver
)

var statusText = map[int]string{
	StatusToFewTeams:           "To Few Teams",
	StatusToManyTeams:          "To Many Teams",
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
