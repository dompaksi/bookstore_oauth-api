package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstant(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.isExpired(), "brand new access token should not be nil")
	assert.EqualValues(t, "", at.AccessToken)
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestGetNewAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.isExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.isExpired(), "access token created three hours from now should NOT be expired")
}