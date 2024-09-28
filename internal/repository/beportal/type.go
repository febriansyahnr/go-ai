package beportal_repository

import (
	"github.com/febriansyahnr/go-ai/config"
	"github.com/febriansyahnr/go-ai/internal/repository"
)

type BEPortal struct {
	config *config.Config
	secret *config.Secret
}

func New() *BEPortal {
	return &BEPortal{}
}

var _ repository.IDisbursement = (*BEPortal)(nil)
