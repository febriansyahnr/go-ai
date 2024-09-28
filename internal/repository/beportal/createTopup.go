package beportal_repository

import (
	"context"
	"net/http"

	"github.com/febriansyahnr/go-ai/constant"
	httputil "github.com/febriansyahnr/go-ai/util/http"
)

// CreateTopup implements repository.IDisbursement.
func (b *BEPortal) CreateTopup(ctx context.Context, header constant.TMapString, paymentID string) (string, error) {
	url := b.config.BePortal.BaseURL + TopupVAUrl
	payload := constant.TMapAny{
		"paymentMethodId": paymentID,
	}

	httputil.RequestHitAPI(ctx, http.MethodPost, url, payload, header)
	panic("unimplemented")
}
