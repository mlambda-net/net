package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Token(t *testing.T) {

	token := NewToken("aaaa")

	result, err := token.Create(map[string]interface{}{
		"user_id":    "roy",
		"authorized": true,
	})
	assert.Nil(t, err)

	claims, err := token.Claims(result)
	assert.Nil(t, err)
	assert.Equal(t, "roy", claims.Get("user_id"))
	assert.Equal(t, true, claims.Get("authorized").(bool))

}
