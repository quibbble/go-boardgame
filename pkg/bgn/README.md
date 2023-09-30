# Board Game Notation - BGN

[![GoDoc](https://godoc.org/github.com/quibbble/go-boardgame/pkg/bgn?status.svg)](https://godoc.org/github.com/quibbble/go-boardgame/pkg/bgn)

Many different standards exist for simplified human and machine readable notation of games. Chess is a common example with notations such as Portable Game Notation (PGN) or Forsyth–Edwards Notation (FEN). While useful for chess, these notations have difficulties being applied to other games. Currently no single standard exists that can be applied to any game hence the introduction of Board Game Notation (BGN) as a potential solution to this problem.

## Format

BGN is structured into two distinct sections, tags and actions.

### Tags

Tags are key value pairs used to describe the initial game state as well as to store additional meta data about the game.

#### Tag Format
```
[Key "Value"]
```

#### Tag Example
```
[Game "Carcassonne"]
```

#### Tag Requirements
There are two tags necessary for any game, Game and Teams.
- Game: the name of the game represented. Ex: `[Game "Carcassonne"]`
- Teams: the list of teams playing the game. Ex: `[Teams "TeamA, TeamB"]`

### Actions

Actions are an ordered list of actions teams take to create the current game state. Actions require the team, the action done, and any additional details needed to perform the action.

#### Action Format
```
{team index}{action character}&{action detail 1}.{action detail 2}...
```

#### Action Examples

```
0a&1.2 // team at index 0 of list in Teams Tag does action a with details 1 and 2
1b     // team at index 1 of list in Teams Tag does action b
```

### BGN Example
```
[Game "Carcassonne"]
[Teams "TeamA, TeamB"]
[Seed "123"]
[Completed "False"]
[Date "10-31-2021"]

0c 0a&1.2 0b&1.2.k.b 1c 1c 1c 1a&0.1 1b&0.1.m {you can add
comments like so} 0a&2.2 0b&2.2.t.l
```

## Usage

This package provides both a method of creating as well as parsing BGN text.

### Create BGN

```go
bgn := &Game{
    Tags: map[string]string{"Game": "Carcassonne", "Teams": "TeamA, TeamB", "Seed": "123"}
    Actions: []Action{
        {
            TeamIndex: 0,
            ActionType: 'a',
            Details: []string{"1", "2"}
        },
        {
            TeamIndex: 1,
            ActionType: 'c',
        },
    }
}
raw := bgn.String()
```

### Parse BGN
```go
raw := "[Game \"Carcassonne\"][Teams \"TeamA, TeamB\"][Seed \"123\"]0c 0a&1.2"
bgn, err := Parse(raw)
```
