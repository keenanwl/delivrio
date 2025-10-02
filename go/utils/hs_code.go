package utils

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/country"
	"delivrio.io/go/ent/countryharmonizedcode"
	"delivrio.io/shared-utils/pulid"
	"fmt"
)

func HsCode(ctx context.Context, inventoryItem *ent.InventoryItem, destCountryID pulid.ID) (string, string, error) {

	if inventoryItem == nil {
		return "", "", fmt.Errorf("inventory item not found")
	}

	tariffCode := ""
	if inventoryItem.Code != nil {
		tariffCode = *inventoryItem.Code
	}

	originCountry, err := inventoryItem.QueryCountryOfOrigin().
		Only(ctx)
	if err != nil {
		return "", "", err
	}

	// Follows same structure as Shopify
	// https://shopify.dev/docs/api/admin-rest/2024-01/resources/inventoryitem
	countrySpecificCode, err := inventoryItem.QueryCountryHarmonizedCode().
		Where(countryharmonizedcode.HasCountryWith(country.ID(destCountryID))).
		Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return "", "", err
	} else if err == nil {
		tariffCode = countrySpecificCode.Code
	}
	return tariffCode, originCountry.Alpha2, nil
}
