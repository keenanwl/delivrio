package utils

import (
	"context"
	"delivrio.io/go/ent"
)

// DefaultProductImg Preparing for images residing in different locations (S3)
func DefaultProductImg(ctx context.Context, p *ent.Product) (string, error) {
	img := ""

	firstImage, err := p.QueryProductImage().
		First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return "", err
	} else if firstImage != nil {
		img = firstImage.URL
	}

	return img, nil
}
