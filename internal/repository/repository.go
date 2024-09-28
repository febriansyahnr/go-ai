package repository

import (
	"context"

	"github.com/febriansyahnr/go-ai/constant"
)

type IDisbursement interface {
	CreateTopup(ctx context.Context, header constant.TMapString, paymentID string) (string, error)
}
