package environment

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestEnvVars(t *testing.T) {
	Load("../")

	assert.MatchRegex(t, os.Getenv("URL"), `localhost|((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}`)
	assert.MatchRegex(t, os.Getenv("PORT"), `^:\d+$`)
	assert.MatchRegex(t, os.Getenv("LOG_TO_FILE"), `TRUE|FALSE`)
	assert.MatchRegex(t, os.Getenv("RELEASE_MODE"), `TEST|PRODUCTION`)
}
