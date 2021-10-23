package go_boardgame

// BoardGameOptions are the options used to create a new game
type BoardGameOptions struct {
	// Teams is the list of teams that will be playing the game
	Teams []string

	// MoreOptions allows for additional game options to be passed that are unique to each game
	MoreOptions interface{}
}

// BoardGameAction represents an action that is performed on the game state
type BoardGameAction struct {
	// Team is the team performing the action
	Team string

	// ActionType is the key that determines what action to perform
	ActionType string

	// MoreDetails allows for additional action details to be passed that are unique to the action type
	MoreDetails interface{}
}

// BoardGameSnapshot represents the current state of the game
type BoardGameSnapshot struct {
	// Turn is the turn of the current team
	Turn string

	// Teams is a list of all teams playing the game
	Teams []string

	// Winners is a list of teams that have won the game
	Winners []string

	// MoreData allows for additional game data to be returned that is unique to each game
	MoreData interface{}

	// Actions is a list of past game actions that have lead to the current game state
	Actions []*BoardGameAction
}
