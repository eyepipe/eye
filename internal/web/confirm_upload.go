package web

import (
	"encoding/hex"
	"fmt"

	"github.com/eyepipe/eye/internal/lib/uuidv7"
	"github.com/eyepipe/eye/internal/pkg/services/export_service"
	"github.com/eyepipe/eye/pkg/proto"
	"github.com/gofiber/fiber/v3"
)

func (w *Web) ConfirmUpload(c fiber.Ctx) error {
	req := new(proto.ConfirmUploadRequestV1)
	err := c.Bind().JSON(req)
	if err != nil {
		return fmt.Errorf("failed to bind request: %w", err)
	}

	jwt, err := w.jwt.DecodeUploadVerificationJWT(req.Token)
	if err != nil {
		return fmt.Errorf("failed to decode jwt: %w", err)
	}

	jwtUUID, err := uuidv7.Decode(jwt.UploadUUID)
	if err != nil {
		return fmt.Errorf("failed to decode uuid: %w", err)
	}

	store := w.stores.MustGet(jwtUUID.Shard)
	upload, err := store.FindUpload(c.Context(), jwtUUID.String())
	if err != nil {
		return fmt.Errorf("failed to find upload: %w", err)
	}

	pubBytes, err := hex.DecodeString(req.PubKeyHex)
	if err != nil {
		return fmt.Errorf("failed to decode public key: %w", err)
	}

	pubKey, err := export_service.SPKIToPublicKey(pubBytes)
	if err != nil {
		return fmt.Errorf("failed to import public key: %w", err)
	}

	sig, err := hex.DecodeString(req.SigHex)
	if err != nil {
		return fmt.Errorf("failed to decode sig: %w", err)
	}

	serverHash, err := hex.DecodeString(jwt.ServerHashHex)
	if err != nil {
		return fmt.Errorf("failed to decode serverhash: %w", err)
	}

	signer := upload.GetSignerAlgo().ToSigner()
	err = signer.Verify(serverHash, sig, pubKey)
	if err != nil {
		return fmt.Errorf("failed to verify signature: %w", err)
	}

	err = store.UpdateUploadSignatureHex(c.Context(), jwt.UploadUUID, hex.EncodeToString(sig))
	if err != nil {
		return fmt.Errorf("failed to update upload signature: %w", err)
	}

	return c.JSON(proto.ConfirmUploadResponseV1{
		URL: fmt.Sprintf("%s/v1/downloads/%s", w.GetServerBaseURL(c), jwtUUID.String()),
	})
}
