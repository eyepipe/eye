package jwt_claims

import "github.com/golang-jwt/jwt/v5"

type GenerateUploadVerificationClaims struct {
	UploadUUID    string `json:"upload_uuid"`
	ServerHashHex string `json:"server_hash_hex"`
	jwt.RegisteredClaims
}
