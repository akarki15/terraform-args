package crawler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseDir(t *testing.T) {
	_, _ = parseDir("testfiles/d")
}

func Test_parseDir_full(t *testing.T) {
	for _, dir := range dirs {
		commands, err := parseDir(dir)
		assert.NoError(t, err)
		b, err := json.Marshal(commands)
		assert.NoError(t, err)
		t.Log(string(b))
	}
}
