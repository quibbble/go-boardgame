package go_boardgame

// BoardGame is a representation of a board game allowing one to perform actions on the game as well as to retrieve game data
type BoardGame interface {
	// Do performs an action on the game
	Do(action *BoardGameAction) error

	// GetSnapshot retrieves the current game state from 'team' view
	// Entering nothing returns a complete snapshot i.e. all hands, resources, etc.
	// Entering more than one team will error
	GetSnapshot(team ...string) (*BoardGameSnapshot, error)
}

// AdvancedBoardGame provides additional functionality that is not required to play a game but can be useful in many cases
type AdvancedBoardGame interface {
	BoardGame

	// GetNotation returns a simplified notation allowing for easy storage and retrieval of games
	// This notation can be defined as desired but below is a simple example of a potential implementation
	// "'num teams':'seed':'create option details':'team num','action num','details','details',...;..."
	// EX: "2:123::0,0,0,F;1,1;" 2 teams, seed is 123, no extra options, team 0 played action 0 with details 0 and F, team 1 played action 1
	GetNotation() string
}
