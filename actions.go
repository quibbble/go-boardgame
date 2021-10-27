package go_boardgame

// common action types most games will implement - not required to be implemented but typically helpful to do so
const (
	ActionReset = "Reset" // completely restarts the game with a fresh game state
	ActionUndo  = "Undo"  // undoes the most recent action
)
