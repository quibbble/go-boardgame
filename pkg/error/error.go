package error

import (
	"fmt"
)

// BoardGameError provides a simple method of returning common board game errors to reduce duplicate code
type BoardGameError struct {
	Err    error
	Status int
}

func (i *BoardGameError) Error() string {
	return fmt.Sprintf("%s: %s", StatusText(i.Status), i.Err.Error())
}
