package tictactoe

// Action types
const (
	// ActionMarkLocation allows players to mark and X or O on the board
	ActionMarkLocation = "MarkLocation"
)

// MarkLocationActionDetails is the action details for marking an X or O on the board
type MarkLocationActionDetails struct {
	Row, Column int
}

// TicTacToeSnapshotData is the game data unique to TicTacToe
type TicTacToeSnapshotData struct {
	Board [size][size]string
}
