package seed

import (
	"context"
	"embed"
	_ "embed"
	b64 "encoding/base64"
	"encoding/csv"
	"fmt"
	"strings"

	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrieradditionalservicepostnord"
	"delivrio.io/go/ent/carrierbrand"
	"delivrio.io/go/ent/carrierservice"
	"delivrio.io/go/ent/carrierservicepostnord"
	"delivrio.io/go/ent/colli"
	"delivrio.io/go/ent/connectionbrand"
	"delivrio.io/go/ent/country"
	"delivrio.io/go/ent/currency"
	"delivrio.io/go/ent/deliveryruleconstraint"
	"delivrio.io/go/ent/deliveryruleconstraintgroup"
	"delivrio.io/go/ent/language"
	"delivrio.io/go/ent/order"
	"delivrio.io/go/schema/fieldjson"
	"delivrio.io/go/utils"
	"delivrio.io/shared-utils/pulid"
)

var (
	tenantID          pulid.ID
	planID            pulid.ID
	languageID        pulid.ID
	carrierBrandPNID  pulid.ID
	carrierID         pulid.ID
	conn              *ent.Connection
	locationID        pulid.ID
	locationReturnsID pulid.ID
	deliveryOptionID  pulid.ID
	adminUserID       pulid.ID
	productVariant    = make([]*ent.ProductVariant, 0)
)

func readCsvFileFS(fs embed.FS, filePath string) ([][]string, error) {
	f, err := fs.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read input file %s: %w", filePath, err)
	}

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Unable to parse file as CSV for %s: %w", filePath, err)
	}

	return records, nil
}

func ToCountryRegion(val string) country.Region {

	switch val {
	case "Africa":
		return country.RegionAfrica
	case "Asia":
		return country.RegionAsia
	case "Americas":
		return country.RegionAmericas
	case "Europe":
		return country.RegionEurope
	case "Oceania":
		return country.RegionOceania
	}

	panic("no region provided")

}

func GetTenantID() pulid.ID {
	if len(tenantID) == 0 {
		panic("TenantID not set")
	}
	return tenantID
}

func GetAdminUser() pulid.ID {
	if len(adminUserID) == 0 {
		panic("AdminUserID not set")
	}
	return adminUserID
}

func GetProductVariants() []*ent.ProductVariant {
	if len(productVariant) == 0 {
		panic("ProductVariants not set")
	}
	return productVariant
}

func Products(ctx context.Context, count int) {
	c := ent.TxFromContext(ctx)

	productVariant = make([]*ent.ProductVariant, 0)

	i := 0
	for i < count {
		i++
		prod := c.Product.Create().
			SetInput(ent.CreateProductInput{
				Title: fmt.Sprintf("Peanuts %v", i),
			}).
			SetExternalID(fmt.Sprintf("%v", i)).
			SetTenantID(tenantID).
			SaveX(ctx)

		weightG := 50 + i
		description := "Brown & Salty"
		ean := fmt.Sprintf("0000110000%v", i)
		productVariant = append(productVariant, c.ProductVariant.Create().
			SetInput(ent.CreateProductVariantInput{
				Description:     &description,
				EanNumber:       &ean,
				WeightG:         &weightG,
				DimensionLength: &weightG,
				DimensionWidth:  &weightG,
				DimensionHeight: &weightG,
			}).
			SetExternalID(fmt.Sprintf("%v", i)).
			SetProduct(prod).
			SetTenantID(tenantID).
			SaveX(ctx))

	}

}

func Location(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	denmark := c.Country.Query().Where(country.Alpha2ContainsFold("DK")).OnlyX(ctx)

	locAddress := c.Address.Create().
		SetAddressOne("Industry road").
		SetAddressTwo("Suite 52").
		SetCountry(denmark).
		SetCity("Aarhus").
		SetZip("8000").
		SetCompany("").
		SetState("").
		SetFirstName("Some first name").
		SetLastName("Some last name").
		SetEmail("Some email").
		SetVatNumber("DK999999").
		SetPhoneNumber("Some phone number").
		SetUniquenessID("default-adress").
		SetTenantID(tenantID).
		SaveX(ctx)

	locationTags := c.LocationTag.Query().AllX(ctx)
	lid := c.Location.Create().
		SetName("Sender name X").
		SetAddress(locAddress).
		SetTenantID(tenantID).
		AddLocationTags(locationTags...).
		SaveX(ctx)
	locationID = lid.ID

	locAddressReturns := c.Address.Create().
		SetAddressOne("Return2me road").
		SetAddressTwo("Behind the restaurant on the corner").
		SetCountry(denmark).
		SetCity("Copenhagen").
		SetZip("2000").
		SetCompany("*").
		SetState("++++").
		SetUniquenessID("return2me").
		SetFirstName("Returns department").
		SetLastName("Lænard").
		SetEmail("returns@example.com").
		SetVatNumber("DK999999").
		SetPhoneNumber("Some phone number").
		SetTenantID(tenantID).
		SaveX(ctx)

	lr := c.Location.Create().
		SetName("Primary return address").
		SetAddress(locAddressReturns).
		SetTenantID(tenantID).
		AddLocationTags(locationTags...).
		SaveX(ctx)
	locationReturnsID = lr.ID
}

func DemoLocation(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	denmark := c.Country.Query().Where(country.Alpha2ContainsFold("DK")).OnlyX(ctx)
	locationTags := c.LocationTag.Query().AllX(ctx)

	locAddress := c.Address.Create().
		SetAddressOne("Krekærvangen 30").
		SetAddressTwo("").
		SetCountry(denmark).
		SetCity("Malling").
		SetZip("8340").
		SetCompany("DELIVRIO ApS").
		SetState("").
		SetFirstName("Anders").
		SetLastName("Goosmann").
		SetEmail("info@delivrio.com").
		SetVatNumber("DK43912690").
		SetPhoneNumber("+45 52 51 12 13").
		SetUniquenessID("anders-address").
		SetTenantID(tenantID).
		SaveX(ctx)

	lid := c.Location.Create().
		SetName("DELIVRIO HQ").
		SetAddress(locAddress).
		SetTenantID(tenantID).
		AddLocationTags(locationTags...).
		SaveX(ctx)
	locationID = lid.ID

	locAddressReturns := c.Address.Create().
		SetAddressOne("Krekærvangen 30").
		SetAddressTwo("").
		SetCountry(denmark).
		SetCity("Malling").
		SetZip("8340").
		SetCompany("DELIVRIO ApS").
		SetFirstName("Anders").
		SetLastName("Goosmann").
		SetEmail("info@delivrio.com").
		SetVatNumber("DK43912690").
		SetPhoneNumber("+45 52 51 12 13").
		SetState("").
		SetUniquenessID("anders-return").
		SetTenantID(tenantID).
		SaveX(ctx)

	l := c.Location.Create().
		SetName("DELIVRIO Warehouse").
		SetAddress(locAddressReturns).
		SetTenantID(tenantID).
		AddLocationTags(locationTags...).
		SaveX(ctx)
	locationReturnsID = l.ID
}

func SeedCarrierConnection(ctx context.Context) {

	c := ent.TxFromContext(ctx)
	car := c.Carrier.Create().
		SetName("PostNord agreement 1").
		SetTenantID(tenantID).
		SetCarrierBrandID(carrierBrandPNID).
		SaveX(ctx)

	carrierID = car.ID

	c.CarrierPostNord.Create().
		SetCarrier(car).
		SetCustomerNumber("111111111").
		SetTenantID(tenantID).
		ExecX(ctx)

	glsBrand := c.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDGLS)).
		OnlyX(ctx)
	car2 := c.Carrier.Create().
		SetName("GLS agreement 1").
		SetTenantID(tenantID).
		SetCarrierBrand(glsBrand).
		SaveX(ctx)

	c.CarrierGLS.Create().
		SetCarrier(car2).
		SetGLSUsername("2080060960").
		SetGLSPassword("API1234").
		SetCustomerID("2080060960").
		SetContactID("208a144Uoo").
		SetTenantID(tenantID).
		ExecX(ctx)

	uspsBrand := c.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDUSPS)).
		OnlyX(ctx)
	car3 := c.Carrier.Create().
		SetName("USPS agreement 1").
		SetTenantID(tenantID).
		SetCarrierBrand(uspsBrand).
		SaveX(ctx)

	c.CarrierUSPS.Create().
		SetCarrier(car3).
		SetConsumerKey("Isq478AhZdPGTwInecPPxCuLUlWMiy9C").
		SetConsumerSecret("lHkVwm97J1IR3VPA").
		SetMid("901097701").
		SetCrid("94879959").
		SetManifestMid("901097699").
		SetIsTestAPI(true).
		SetEpsAccountNumber("1000005549").
		SetTenantID(tenantID).
		ExecX(ctx)

	cur := c.Currency.Query().Where(currency.CurrencyCodeEQ(currency.CurrencyCodeDKK)).OnlyX(ctx)
	connBrand := c.ConnectionBrand.Query().Where(connectionbrand.LabelEQ("Shopify")).OnlyX(ctx)
	conn = c.Connection.Create().
		SetPickupLocationID(locationID).
		SetReturnLocationID(locationReturnsID).
		SetSellerLocationID(locationID).
		SetSenderLocationID(locationID).
		SetConnectionBrand(connBrand).
		SetCurrency(cur).
		SetSyncOrders(true).
		SetSyncProducts(true).
		SetConvertCurrency(true).
		SetFulfillAutomatically(true).
		SetDispatchAutomatically(true).
		SetName("Shopify DK `'*øæ~~~^@£€¤").
		SetTenantID(tenantID).
		SaveX(ctx)

	c.ConnectionShopify.Create().
		SetTenantID(tenantID).
		SetConnection(conn).
		SetStoreURL("https://delivrio.myshopify.com").
		SetAPIKey("somecode111111").
		SaveX(ctx)
}

func DeliveryOption(ctx context.Context) {
	c := ent.TxFromContext(ctx)

	cs := c.CarrierService.Query().
		Where(carrierservice.HasCarrierBrandWith(carrierbrand.InternalIDEQ(carrierbrand.InternalIDPostNord))).
		FirstX(ctx)

	do := c.DeliveryOption.Create().
		SetCarrierID(carrierID).
		SetCarrierServiceID(c.CarrierService.Query().FirstIDX(ctx)).
		SetTenantID(tenantID).
		SetConnection(conn).
		SetCarrierService(cs).
		SetName("PostNord Home").
		SetSortOrder(1).
		SaveX(ctx)

	deliveryOptionID = do.ID

	c.DeliveryOptionPostNord.Create().
		SetTenantID(tenantID).
		SetFormatZpl(true).
		SetDeliveryOption(do).
		SaveX(ctx)

	cur := c.Currency.Query().
		Where(currency.CurrencyCodeEQ(currency.CurrencyCodeDKK)).
		OnlyX(ctx)

	dr := c.DeliveryRule.Create().
		SetDeliveryOption(do).
		SetTenantID(tenantID).
		SetName("Cart minimum").
		SetPrice(20.00).
		SetCurrency(cur).
		SaveX(ctx)

	cg := c.DeliveryRuleConstraintGroup.Create().
		SetTenantID(tenantID).
		SetDeliveryRule(dr).
		SetTenantID(tenantID).
		SetConstraintLogic(deliveryruleconstraintgroup.ConstraintLogicAnd).
		SaveX(ctx)

	c.DeliveryRuleConstraint.Create().
		SetTenantID(tenantID).
		SetComparison(deliveryruleconstraint.ComparisonGreaterThan).
		SetSelectedValue(&fieldjson.DeliveryRuleConstraintSelectedValue{
			Numeric: 0,
		}).
		SetDeliveryRuleConstraintGroup(cg).
		SetPropertyType(deliveryruleconstraint.PropertyTypeCartTotal).
		ExecX(ctx)

}

func SeedOrders(ctx context.Context, count int) {

	c := ent.TxFromContext(ctx)
	denmark := c.Country.Query().Where(country.Alpha2ContainsFold("DK")).OnlyX(ctx)

	i := 0

	for i < count {
		i++
		r1 := c.Address.Create().
			SetAddressOne("999 main st.").
			SetAddressTwo("Apt 12").
			SetCountry(denmark).
			SetCity("Aarhus").
			SetZip("8000").
			SetCompany("Pam's Company Ghmb").
			SetState("Midtjylland").
			SetFirstName("Pam").
			SetLastName("Armstrong").
			SetEmail("pam@example.com").
			SetVatNumber("DK0000000").
			SetPhoneNumber("+45 22 22 22 22").
			SetUniquenessID(fmt.Sprintf("r%v", i)).
			SetTenantID(tenantID).
			SaveX(ctx)

		s1 := c.Address.Create().
			SetAddressOne("Shipper avenue").
			SetAddressTwo("Department 1").
			SetCountry(denmark).
			SetCity("Aarhus").
			SetZip("8000").
			SetCompany("DELIVRIO of course").
			SetState("Midtjylland").
			SetFirstName("Returns").
			SetLastName("Department").
			SetVatNumber("DK909090").
			SetEmail("support@example.com").
			SetPhoneNumber("+45 11 11 11 11").
			SetUniquenessID(fmt.Sprintf("s%v", i)).
			SetTenantID(tenantID).
			SaveX(ctx)
		fmt.Printf("TENNANT: %v", tenantID)
		ord := c.Order.Create().
			SetStatus(order.StatusPending).
			SetOrderPublicID(fmt.Sprintf("100%v", i)).
			SetConnection(conn).
			SetTenantID(tenantID).
			SaveX(ctx)

		col := c.Colli.Create().
			SetSender(s1).
			SetRecipient(r1).
			SetStatus(colli.StatusPending).
			SetOrder(ord).
			SetTenantID(tenantID).
			SetDeliveryOptionID(deliveryOptionID).
			SaveX(ctx)

		cur := c.Currency.Query().
			Where(currency.CurrencyCodeEQ(currency.CurrencyCodeDKK)).
			OnlyX(ctx)

		_ = c.OrderLine.Create().
			SetColli(col).
			SetProductVariant(productVariant[0]).
			SetTenantID(tenantID).
			SetUnits(1).
			SetUnitPrice(900).
			SetDiscountAllocationAmount(99).
			SetCurrency(cur).
			SaveX(ctx)
	}

}

func ExtraOrderLines(ctx context.Context) {
	c := ent.TxFromContext(ctx)

	col := c.Colli.Query().FirstX(ctx)

	cur := c.Currency.Query().
		Where(currency.CurrencyCodeEQ(currency.CurrencyCodeDKK)).
		OnlyX(ctx)

	// Duplicate product, different price
	_ = c.OrderLine.Create().
		SetColli(col).
		SetProductVariant(productVariant[0]).
		SetTenantID(tenantID).
		SetUnits(5).
		SetUnitPrice(200).
		SetDiscountAllocationAmount(9).
		SetCurrency(cur).
		SaveX(ctx)

	_ = c.OrderLine.Create().
		SetColli(col).
		SetProductVariant(productVariant[1]).
		SetTenantID(tenantID).
		SetUnits(50).
		SetUnitPrice(1).
		SetDiscountAllocationAmount(0).
		SetCurrency(cur).
		SaveX(ctx)
}

func SeedCustomerUser(ctx context.Context) *ent.User {
	tx := ent.TxFromContext(ctx)
	return tx.User.Create().
		SetTenantID(tenantID).
		SetName("John").
		SetIsAccountOwner(true).
		SetHash("llllll").
		SetEmail("john@example.com").
		SaveX(ctx)
}

func Languages(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	lang := c.Language.Create().
		SetLabel("English").
		SetInternalID(language.InternalIDEN).
		SaveX(ctx)
	languageID = lang.ID

	c.Language.Create().
		SetLabel("Dansk").
		SetInternalID(language.InternalIDDA).
		SaveX(ctx)
}

func SeedSignup(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	_ = c.ConnectOptionCarrier.Create().SetName("Bring").SaveX(ctx)
	_ = c.ConnectOptionCarrier.Create().SetName("DHL").SaveX(ctx)
	_ = c.ConnectOptionCarrier.Create().SetName("PostNord").SaveX(ctx)
	_ = c.ConnectOptionCarrier.Create().SetName("FedEx").SaveX(ctx)
	_ = c.ConnectOptionCarrier.Create().SetName("DAO").SaveX(ctx)
	_ = c.ConnectOptionCarrier.Create().SetName("USPS").SaveX(ctx)
	_ = c.ConnectOptionCarrier.Create().SetName("UPS").SaveX(ctx)
	_ = c.ConnectOptionCarrier.Create().SetName("Burd").SaveX(ctx)
	_ = c.ConnectOptionCarrier.Create().SetName("Danske Fragtmænd").SaveX(ctx)
	_ = c.ConnectOptionCarrier.Create().SetName("DSV").SaveX(ctx)

	_ = c.ConnectOptionPlatform.Create().SetName("Shopify").SaveX(ctx)
	_ = c.ConnectOptionPlatform.Create().SetName("WooCommerce").SaveX(ctx)
	_ = c.ConnectOptionPlatform.Create().SetName("Magento").SaveX(ctx)
}

func SeedPlans(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	p := c.Plan.Create().
		SetLabel("Free").
		SetPriceDkk(0).
		SetRank(0).
		SaveX(ctx)
	planID = p.ID
	_ = c.Plan.Create().SetLabel("Basic").SetPriceDkk(649).SetRank(1).
		SaveX(ctx)
	_ = c.Plan.Create().SetLabel("Growth").SetPriceDkk(1649).SetRank(2).
		SaveX(ctx)
	_ = c.Plan.Create().SetLabel("Pro").SetPriceDkk(3749).SetRank(3).
		SaveX(ctx)
}

func SeedTenant(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	t := c.Tenant.Create().
		SetPlanID(planID).
		SetDefaultLanguageID(languageID).
		SetName("DELIVRIO ApS").
		SetVatNumber("DK1234567").
		SaveX(ctx)
	tenantID = t.ID
}

func LocationTags(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	c.LocationTag.Create().
		SetLabel("Sender").
		SetInternalID("sender").
		SaveX(ctx)
	c.LocationTag.Create().
		SetLabel("Seller").
		SetInternalID("seller").
		SaveX(ctx)
	c.LocationTag.Create().
		SetLabel("Return").
		SetInternalID("return").
		SaveX(ctx)
	c.LocationTag.Create().
		SetLabel("Pickup").
		SetInternalID("pickup").
		SaveX(ctx)
	c.LocationTag.Create().
		SetLabel("Agent").
		SetInternalID("agent").
		SaveX(ctx)
	c.LocationTag.Create().
		SetLabel("Click & collect").
		SetInternalID("click_and_collect").
		SaveX(ctx)
}
func CarrierBrands(ctx context.Context) []pulid.ID {
	c := ent.TxFromContext(ctx)
	cb := c.CarrierBrand.Create().
		SetLabel("GLS").
		SetLabelShort("GLS").
		SetInternalID(carrierbrand.InternalIDGLS).
		SetBackgroundColor("#061AB1").
		SetTextColor("#FFFFFF").
		SaveX(ctx)

	cbUSPS := c.CarrierBrand.Create().
		SetLabel("USPS").
		SetLabelShort("USPS").
		SetInternalID(carrierbrand.InternalIDUSPS).
		SetBackgroundColor("#333366").
		SetTextColor("#FFFFFF").
		SaveX(ctx)

	cbpn := c.CarrierBrand.Create().
		SetLabel("PostNord").
		SetLabelShort("PN").
		SetInternalID(carrierbrand.InternalIDPostNord).
		SetBackgroundColor("#0097B9").
		SetTextColor("#FFFFFF").
		SaveX(ctx)
	carrierBrandPNID = cbpn.ID

	c.CarrierBrand.Create().
		SetLabel("Bring").
		SetLabelShort("BR").
		SetInternalID(carrierbrand.InternalIDBring).
		SetBackgroundColor("#002e18").
		SetTextColor("#FFFFFF").
		SaveX(ctx)

	c.CarrierBrand.Create().
		SetLabel("DAO").
		SetLabelShort("DA").
		SetInternalID(carrierbrand.InternalIDDAO).
		SetBackgroundColor("#e30613").
		SetTextColor("#FFFFFF").
		SaveX(ctx)

	c.CarrierBrand.Create().
		SetLabel("DSV").
		SetLabelShort("DS").
		SetInternalID(carrierbrand.InternalIDDSV).
		SetBackgroundColor("#182a61FF").
		SetTextColor("#FFFFFF").
		SaveX(ctx)

	c.CarrierBrand.Create().
		SetLabel("Danske Fragtmænd").
		SetLabelShort("DF").
		SetInternalID(carrierbrand.InternalIDDF).
		SetBackgroundColor("#009ADA").
		SetTextColor("#FFFFFF").
		SaveX(ctx)

	c.CarrierBrand.Create().
		SetLabel("EasyPost").
		SetLabelShort("EP").
		SetInternalID(carrierbrand.InternalIDEasyPost).
		SetBackgroundColor("#164dff").
		SetTextColor("#FFFFFF").
		SaveX(ctx)

	return []pulid.ID{cb.ID, cbUSPS.ID, carrierBrandPNID}
}

//go:embed post_nord_service_codes.csv
var embeddedFSPN embed.FS

const PostNordDeliveryPointServiceInternalID = "optional_service_point"

func PNServices(ctx context.Context) {
	pnServiceCodes, err := readCsvFileFS(embeddedFSPN, "post_nord_service_codes.csv")
	if err != nil {
		panic(err)
	}

	c := ent.TxFromContext(ctx)

	pnCarrierBrand := c.CarrierBrand.Query().
		Where(carrierbrand.InternalIDEQ(carrierbrand.InternalIDPostNord)).
		OnlyX(ctx)

	for _, sc := range pnServiceCodes {
		if sc[2] == "DK" {

			internalID := strings.Replace(strings.ToLower(sc[4]), " ", "_", -1)

			isReturn := false
			if internalID == "return_pickup" || internalID == "return_dropoff" {
				isReturn = true
			}
			cs, err := c.CarrierService.Create().
				SetLabel(sc[4]).
				SetInternalID(fmt.Sprintf("PN_%s", internalID)).
				SetCarrierBrand(pnCarrierBrand).
				SetReturn(isReturn).
				OnConflictColumns(carrierservice.FieldInternalID).
				UpdateInternalID().
				ID(ctx)
			if err != nil {
				panic(err)
			}
			serviceID, err := c.CarrierServicePostNord.Create().
				SetLabel(sc[4]).
				SetAPICode(sc[3]).
				SetInternalID(internalID).
				SetCarrierServiceID(cs).
				OnConflictColumns(carrierservicepostnord.FieldInternalID).
				UpdateInternalID().
				ID(ctx)
			if err != nil {
				panic(err)
			}

			internalID = strings.Replace(strings.ToLower(sc[6]), " ", "_", -1)
			allCountriesConsignee := sc[7] == "ALL"
			allCountriesConsignor := sc[8] == "ALL"
			mandatory := sc[9] == "TRUE"

			create := c.CarrierAdditionalServicePostNord.Create().
				SetLabel(sc[6]).
				SetAPICode(sc[5]).
				SetMandatory(mandatory).
				SetAllCountriesConsignee(allCountriesConsignee).
				SetAllCountriesConsignor(allCountriesConsignor)

			if !allCountriesConsignee {
				countryService := c.Country.Query().Where(country.Alpha2ContainsFold(sc[7])).OnlyX(ctx)
				create = create.AddCountriesConsignee(countryService)
			}

			if !allCountriesConsignor {
				countryService := c.Country.Query().Where(country.Alpha2ContainsFold(sc[8])).OnlyX(ctx)
				create = create.AddCountriesConsignor(countryService)
			}

			count, err := c.CarrierAdditionalServicePostNord.Query().
				Where(
					carrieradditionalservicepostnord.And(
						carrieradditionalservicepostnord.InternalIDEQ(internalID),
						carrieradditionalservicepostnord.HasCarrierServicePostNordWith(
							carrierservicepostnord.ID(serviceID),
						),
					)).
				Count(ctx)
			if err != nil {
				panic(err)
			}

			if count == 0 {
				err = create.SetInternalID(internalID).
					SetCarrierServicePostNordID(serviceID).
					Exec(ctx)
				if err != nil {
					panic(err)
				}
			}

		}
	}
}

//go:embed countries_iso_3166.csv
var embeddedFSCountries embed.FS

func Countries(ctx context.Context) {
	countries, err := readCsvFileFS(embeddedFSCountries, "countries_iso_3166.csv")
	if err != nil {
		panic(err)
	}

	c := ent.TxFromContext(ctx)

	allCountryCreates := make([]*ent.CountryCreate, 0)

	for i, c1 := range countries {
		if i > 0 {
			allCountryCreates = append(allCountryCreates, c.Country.Create().
				SetLabel(c1[0]).
				SetAlpha2(c1[1]).
				SetAlpha3(c1[2]).
				SetCode(c1[3]).
				SetRegion(ToCountryRegion(c1[5])),
			)
		}
	}

	_, err = c.Country.CreateBulk(allCountryCreates...).
		Save(ctx)

}

func Currency(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	c.Currency.Create().
		SetCurrencyCode(currency.CurrencyCodeDKK).
		SetDisplay("DKK").
		Exec(ctx)
	c.Currency.Create().
		SetCurrencyCode(currency.CurrencyCodeEUR).
		SetDisplay("EUR").
		Exec(ctx)
	c.Currency.Create().
		SetCurrencyCode(currency.CurrencyCodeUSD).
		SetDisplay("USD").
		Exec(ctx)
}

func AccessRights(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	c.AccessRight.Create().
		SetLabel("Orders").
		SetInternalID("orders").
		SaveX(ctx)

	c.AccessRight.Create().
		SetLabel("Shipments").
		SetInternalID("shipments").
		SaveX(ctx)
}

func ConnectionBrands(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	_ = c.ConnectionBrand.Create().
		SetLabel("Shopify").
		SaveX(ctx)
	_ = c.ConnectionBrand.Create().
		SetLabel("WooCommerce").
		SaveX(ctx)
}

func APICredentials(ctx context.Context) string {
	c := ent.TxFromContext(ctx)
	key := utils.RandomX(30)
	tokenHashed := utils.HashPasswordX(key)

	t := c.APIToken.Create().
		SetName("External system").
		SetTenantID(GetTenantID()).
		SetUserID(adminUserID).
		SetHashedToken(tokenHashed).
		SaveX(ctx)
	return b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", key, t.ID)))
}
