package jwt_token

import (
	"fmt"
	"time"

	"github.com/eyepipe/eye/internal/lib/jwt_token/jwt_claims"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	issuer string
	secret []byte
	algo   jwt.SigningMethod // @example jwt.SigningMethodHS512
}

func New(issuer string, algo Algo, secret []byte) (*Token, error) {
	method := algo.ToSigningMethod()
	switch {
	case !algo.IsValid():
		return nil, fmt.Errorf("invalid algo: <%s>", algo.String())
	case method == nil:
		return nil, fmt.Errorf("blank signing method for alog <%s>", algo.String())
	}

	return &Token{
		issuer: issuer,
		secret: secret,
		algo:   method,
	}, nil
}

func (t *Token) GenerateUploadVerificationJWT(claims *jwt_claims.GenerateUploadVerificationClaims) (string, error) {
	token := jwt.NewWithClaims(t.algo,
		jwt_claims.GenerateUploadVerificationClaims{
			UploadUUID:    claims.UploadUUID,
			ServerHashHex: claims.ServerHashHex,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    t.issuer,
				Subject:   SubjectUploadVerification,
				Audience:  []string{AudFrontend},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	)

	return token.SignedString(t.secret)
}

func (t *Token) DecodeUploadVerificationJWT(token string) (*jwt_claims.GenerateUploadVerificationClaims, error) {
	out := new(jwt_claims.GenerateUploadVerificationClaims)
	err := t.decode(out, token, DecodeOpts{
		Subject:  SubjectUploadVerification,
		Audience: AudFrontend,
	})
	return out, err
}

func (t *Token) decode(dest jwt.Claims, token string, opts DecodeOpts) error {
	keyFn := func(token *jwt.Token) (any, error) {
		return t.secret, nil
	}

	got, err := jwt.ParseWithClaims(token, dest, keyFn,
		jwt.WithIssuedAt(),
		jwt.WithIssuer(t.issuer),
		jwt.WithSubject(opts.Subject),
		jwt.WithAudience(opts.Audience),
		jwt.WithValidMethods([]string{t.algo.Alg()}),
	)
	switch {
	case err != nil:
		return fmt.Errorf("failed to parse token: %w", err)
	case !got.Valid:
		return fmt.Errorf("invalid token")
	default:
		return nil
	}
}
