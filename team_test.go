package socialblade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeam(t *testing.T) {
	sb := AuthAsAnon()
	_, err := sb.Team()
	assert.NotNil(t, err)
}
