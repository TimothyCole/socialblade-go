package socialblade

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	key = os.Getenv("SOCIALBLADE_KEY")

	email = os.Getenv("SOCIALBLADE_EMAIL")
	token = os.Getenv("SOCIALBLADE_TOKEN")
}

func TestAuth(t *testing.T) {
	sb, err := Auth(key)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		assert.Equal(t, sb.user, false)
		assert.Equal(t, sb.key, key)

		t.Log("Authed as Third-Party\n - Key:", obfuscated(sb.key, 10))
	}
}

func TestAuthAsUser(t *testing.T) {
	sb, err := AuthAsUser(email, token)
	assert.Nil(t, err)
	if assert.NotNil(t, sb) {
		assert.Equal(t, sb.user, true)
		assert.Equal(t, sb.email, email)
		assert.Equal(t, sb.token, token)

		t.Log("Authed as User\n - Email:", sb.email, "\n - Token:", obfuscated(sb.token, 10))
	}
}

func obfuscated(v string, l int) string {
	if l >= len(v) {
		return v
	}

	var m = len(v[l:]) - l
	return v[0:l] + strings.Repeat("*", m) + v[m+l:]
}
