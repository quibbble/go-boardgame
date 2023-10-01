package bgn

const (
	GameTag      string = "Game"
	TeamsTag     string = "Teams"
	VariantTag   string = "Variant"
	SeedTag      string = "Seed"
	CompletedTag string = "Completed"
	CreatedAtTag string = "CreatedAt"
	UpdatedAtTag string = "UpdatedAt"
)

// RequiredTags are the list of tags any game must include
var RequiredTags = []string{
	GameTag,  // the name of the game being played
	TeamsTag, // list of teams playing the game
}
