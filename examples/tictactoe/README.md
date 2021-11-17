# Tic-Tac-Toe Example

This is an example implementation of Tic-Tac-Toe using the [go-boardgame](https://github.com/quibbble/go-boardgame) package. This is by no means the only or necessarily best way to implement a game but is a good place to look when getting started.

## Layout 

### builder.go

The builder object that implements bg.BoardGameBuilder lives here. This allows one to create a new game given a set of options and contains other information such as the game key.

### tictactoe.go

The game object that implements bg.BoardGame lives here. This allows a user to perform an action on the game state as well as view the current state of the game.

### state.go

The state object which handles all the core game logic lives here.

### models.go

The models such as action types, action details, and snapshot details live here.

### bgn.go

All Board Game Notation logic lives here. This is not necessary to play a game and may be ignored if desired. However, if implemented this will allow users to make use of [bgn](https://github.com/quibbble/go-boardgame/tree/main/pkg/bgn) for easy game storage and retrieval.

## Usage

### Create Game

```go
builder := &Builder{}
game, err := builder.Create(&bg.BoardGameOptions{
    Teams: []string{"TeamA", "TeamB"},
    Seed: 123,
})
```

### Play Game

```go
err := game.Do(&bg.BoardGameAction{
    Team: "TeamA",
    ActionType: "MarkLocation",
    MoreDetails: MarkLocationActionDetails{
        Row: 0,
        Column: 0,
    },
})
```

### View Game

```go
snapshot, err := game.GetSnapshot("TeamA")
```
