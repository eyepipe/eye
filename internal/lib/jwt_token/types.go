package jwt_token

import (
	"github.com/eyepipe/eye/internal/lib/jwt_token/jwt_claims"
	"github.com/golang-jwt/jwt/v5"
)

//go:generate go-enum --file types.go --values

const (
	SubjectUploadVerification = "upload-verification"
	AudFrontend               = "frontend"
)

// Algo
// ENUM(
//
//	HS256
//	HS384
//	HS512
//
// )
type Algo string

type IToken interface {
	GenerateUploadVerificationJWT(claims *jwt_claims.GenerateUploadVerificationClaims) (string, error)
	DecodeUploadVerificationJWT(token string) (*jwt_claims.GenerateUploadVerificationClaims, error)
}

type DecodeOpts struct {
	Subject  string
	Audience string
}

func (a Algo) ToSigningMethod() jwt.SigningMethod {
	switch a {
	case AlgoHS256:
		return jwt.SigningMethodHS256
	case AlgoHS384:
		return jwt.SigningMethodHS384
	case AlgoHS512:
		return jwt.SigningMethodHS512
	default:
		return nil
	}
}
