package jwt_token

import (
	"testing"

	"github.com/eyepipe/eye/internal/lib/jwt_token/jwt_claims"
	"github.com/stretchr/testify/require"
)

func TestToken_GenerateUploadVerificationJWT(t *testing.T) {
	t.Parallel()

	t.Run("it works", func(t *testing.T) {
		token := MustNew(t)

		claims := &jwt_claims.GenerateUploadVerificationClaims{
			UploadUUID:    "UploadUUID",
			ServerHashHex: "ServerHashHex",
		}
		got, err := token.GenerateUploadVerificationJWT(claims)
		require.NoError(t, err)
		require.NotEmpty(t, got)

		parsed, err := token.DecodeUploadVerificationJWT(got)
		require.NoError(t, err)
		require.Equal(t, claims.UploadUUID, parsed.UploadUUID)
		require.Equal(t, claims.ServerHashHex, parsed.ServerHashHex)

		// fake
		_, err = token.DecodeUploadVerificationJWT("foo")
		require.Error(t, err)
	})
}
