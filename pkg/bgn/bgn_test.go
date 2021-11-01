package bgn

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"text/scanner"
)

func Test_BGN_Parse(t *testing.T) {
	tests := []struct {
		name        string
		bgn         string
		shouldError bool
	}{
		{
			name:        "empty string should error",
			bgn:         "",
			shouldError: true,
		},
		{
			name:        "random string should error",
			bgn:         "12al12wt31nwauhwaj7haw",
			shouldError: true,
		},
		{
			name:        "missing Teams tag should error",
			bgn:         "[Seed \"123\"][Game \"Carcassonne\"]",
			shouldError: true,
		},
		{
			name:        "tags only should succeed",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]",
			shouldError: false,
		},
		{
			name:        "nested right bracket should error",
			bgn:         "[[Teams] \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]",
			shouldError: true,
		},
		{
			name:        "nested left bracket should error",
			bgn:         "[Teams]] \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]",
			shouldError: true,
		},
		{
			name:        "missing action key should error",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]0",
			shouldError: true,
		},
		{
			name:        "missing team index should error",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]a",
			shouldError: true,
		},
		{
			name:        "large team index should succeed",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]10a",
			shouldError: false,
		},
		{
			name:        "too many dashes should error",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"]10a-1.2-1.1",
			shouldError: true,
		},
		{
			name:        "complex bgn should succeed",
			bgn:         "[Teams \"red, blue\"][Seed \"123\"][Game \"Carcassonne\"][Date \"11-01-2021\"]0a-1.2 {Comment} 0b-1.2.T.R 1c 1a-1.3 1b-1.3.K.B",
			shouldError: false,
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.bgn)
		sc := scanner.Scanner{}
		sc.Init(r)
		_, err := ParseGame(&sc)
		assert.Equal(t, test.shouldError, err != nil, test.name)
	}
}
