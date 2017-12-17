package crawler

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseDir(t *testing.T) {
	_, _ = parseDir("/Users/aashishkarki/Desktop/terraform-provider-aws-master/website/docs/d", "d")
}

func Test_parseDir_full(t *testing.T) {
	for _, dir := range dDirs {
		commands, err := parseDir(dir, "d")
		assert.NoError(t, err)
		b, err := json.Marshal(commands)
		assert.NoError(t, err)
		t.Log(string(b))
	}
}

func Test_parse_d_Dir(t *testing.T) {
	for _, dir := range rDir {
		commands, err := parseDir(dir, "r")
		assert.NoError(t, err)
		b, err := json.Marshal(commands)
		assert.NoError(t, err)
		t.Log(string(b))
	}
}
