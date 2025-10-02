package common

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/shared-utils/pulid"
	"fmt"
)

// Own package to prevent import cycle error

type ReturnOrderDeliveryOptionsColliIDs struct {
	ReturnColliID    pulid.ID `json:"return_colli_id"`
	DeliveryOptionID pulid.ID `json:"delivery_option_id"`
}

type CreateShipment struct {
	Shipment pulid.ID
	Labels   []string
}

func GroupPackagesBySenderReceiver(ctx context.Context, allPackages []*ent.Colli) ([][]*ent.Colli, error) {
	output := make([][]*ent.Colli, 0)
	var lastSender *ent.Address
	var lastRecipient *ent.Address
	var lastParcelShop pulid.ID
	var err error

	for i, p := range allPackages {
		if i == 0 {
			lastSender, err = p.QuerySender().
				WithCountry().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			lastRecipient, err = p.QueryRecipient().
				WithCountry().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			lastParcelShop, err = p.QueryParcelShop().
				OnlyID(ctx)
			if err != nil && !ent.IsNotFound(err) {
				return nil, err
			}

			output = append(output, make([]*ent.Colli, 0))
			output[len(output)-1] = append(output[len(output)-1], p)
		} else {

			currentParcelShop, err := p.QueryParcelShop().
				OnlyID(ctx)
			if err != nil && !ent.IsNotFound(err) {
				return nil, err
			}

			currentSender, err := p.QuerySender().
				WithCountry().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			currentRecipient, err := p.QueryRecipient().
				WithCountry().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			senderAddressMatches := currentSender.Matches(lastSender)
			senderCountryMatches := currentSender.Edges.Country.ID == lastSender.Edges.Country.ID
			senderMatches := currentSender.Matches(lastSender) && senderCountryMatches && senderAddressMatches

			recipientAddressMatches := currentRecipient.Matches(lastRecipient)
			recipientCountryMatches := currentRecipient.Edges.Country.ID == lastRecipient.Edges.Country.ID
			recipientMatches := currentRecipient.Matches(lastRecipient) && recipientCountryMatches && recipientAddressMatches

			parcelShopMatches := currentParcelShop == lastParcelShop

			if !senderMatches || !recipientMatches || !parcelShopMatches {
				output = append(output, make([]*ent.Colli, 0))
			}

			output[len(output)-1] = append(output[len(output)-1], p)

			lastParcelShop = currentParcelShop
			lastSender = currentSender
			lastRecipient = currentRecipient
		}
	}

	return output, nil
}

func GroupReturnPackages(ctx context.Context, allPackages []*ent.ReturnColli) ([][]*ent.ReturnColli, error) {
	output := make([][]*ent.ReturnColli, 0)
	var lastSender *ent.Address
	var lastRecipient *ent.Address
	var err error

	for i, p := range allPackages {
		if i == 0 {
			lastSender, err = p.QuerySender().
				WithCountry().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			output = append(output, make([]*ent.ReturnColli, 0))
			output[len(output)-1] = append(output[len(output)-1], p)
		} else {

			currentSender, err := p.QuerySender().
				WithCountry().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			currentRecipient, err := p.QueryRecipient().
				WithCountry().
				Only(ctx)
			if err != nil {
				return nil, err
			}

			senderAddressMatches := currentSender.Matches(lastSender)
			senderCountryMatches := currentSender.Edges.Country.ID == lastSender.Edges.Country.ID
			senderMatches := currentSender.Matches(lastSender) && senderCountryMatches && senderAddressMatches

			recipientAddressMatches := currentRecipient.Matches(lastRecipient)
			recipientCountryMatches := currentRecipient.Edges.Country.ID == lastRecipient.Edges.Country.ID
			recipientMatches := currentRecipient.Matches(lastRecipient) && recipientCountryMatches && recipientAddressMatches

			if !senderMatches || !recipientMatches {
				output = append(output, make([]*ent.ReturnColli, 0))
			}

			output[len(output)-1] = append(output[len(output)-1], p)

			lastSender = currentSender
			lastRecipient = currentRecipient

		}
	}

	return output, nil
}

// ColliWeightGram ColliWeight calculates the weight in grams of the colli
// base on the order lines
func ColliWeightGram(ctx context.Context, lines []*ent.OrderLine) (int, error) {

	totalWeightGrams := 0
	for _, l := range lines {

		units := l.Units

		variant, err := l.ProductVariant(ctx)
		if err != nil {
			return 0, fmt.Errorf("colli weight: %w", err)
		}

		w := 0
		if variant.WeightG != nil {
			w = *variant.WeightG * units
		}
		totalWeightGrams += w
	}
	return totalWeightGrams, nil
}

func ColliWeightKG(ctx context.Context, lines []*ent.OrderLine) (float64, error) {
	weightGrams, err := ColliWeightGram(ctx, lines)
	if err != nil {
		return 0, err
	}
	return float64(weightGrams) / float64(1000), nil
}

func NetOrderLinePrice(ol *ent.OrderLine) float64 {
	return (ol.UnitPrice * float64(ol.Units)) - ol.DiscountAllocationAmount
}

// Applies the various layers of packaging defaults
// if colli does not have packaging assigned directly
func ColliPackaging(ctx context.Context, c *ent.Colli) (*ent.Packaging, error) {
	p, err := c.QueryPackaging().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if p != nil {
		return p, nil
	}

	// Add workstation default before DO default
	// WS needs to save packaging to the colli
	// so it is used on the return

	// Should always have a DO by this point
	dp, err := c.QueryDeliveryOption().
		QueryDefaultPackaging().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if ent.IsNotFound(err) {
		return nil, fmt.Errorf("packaging or default packaging required: %w", err)
	}

	return dp, nil
}

// Applies the various layers of packaging defaults
// if colli does not have packaging assigned directly
func ReturnColliPackaging(ctx context.Context, c *ent.ReturnColli) (*ent.Packaging, error) {
	p, err := c.QueryPackaging().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if p != nil {
		return p, nil
	}

	dp, err := c.QueryDeliveryOption().
		QueryDefaultPackaging().
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if err == nil {
		return dp, nil
	}

	// Default to outbound packaging if Return Colli does not have packaging
	outboundColli, err := c.QueryReturnOrderLine().
		QueryOrderLine().
		QueryColli().
		First(ctx)
	if err != nil {
		return nil, err
	}

	outboundPackaging, err := ColliPackaging(ctx, outboundColli)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	} else if ent.IsNotFound(err) {
		return nil, err
	}

	return outboundPackaging, nil
}
