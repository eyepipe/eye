package encryption_manager

import (
	"github.com/eyepipe/eye/internal/pkg/eye_api"
	"github.com/eyepipe/eye/internal/pkg/services/encryption_service"
	"github.com/eyepipe/eye/internal/pkg/services/signature_reader_service"
	"resty.dev/v3"
)

type Manager struct {
	service         encryption_service.IService
	api             *eye_api.Client
	downloader      *resty.Client
	signatureReader *signature_reader_service.SignatureReaderService
}

func NewManager(service encryption_service.IService) *Manager {
	return &Manager{
		service:         service,
		api:             eye_api.NewClient(),
		signatureReader: signature_reader_service.NewService(),
		// TODO move redirection policy under the downloader.R()
		downloader: resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()),
	}
}
