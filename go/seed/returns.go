package seed

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/carrierservice"
)

func ReturnPortal(ctx context.Context) {
	c := ent.TxFromContext(ctx)
	rp := c.ReturnPortal.Create().
		SetName("Web-Denmark").
		SetTenantID(tenantID).
		SetConnection(conn).
		AddReturnLocationIDs(locationID).
		SaveX(ctx)

	c.ReturnPortalClaim.Create().
		SetReturnPortal(rp).
		SetName("Too big").
		SetDescription("").
		SetRestockable(false).
		SetArchived(false).
		SetTenantID(tenantID).
		ExecX(ctx)

	c.ReturnPortalClaim.Create().
		SetReturnPortal(rp).
		SetName("Too small").
		SetDescription("").
		SetRestockable(true).
		SetArchived(false).
		SetTenantID(tenantID).
		ExecX(ctx)

	c.ReturnPortalClaim.Create().
		SetReturnPortal(rp).
		SetName("Too archived").
		SetDescription("").
		SetRestockable(true).
		SetArchived(true).
		SetTenantID(tenantID).
		ExecX(ctx)
}

func ReturnDeliveryOption(ctx context.Context) *ent.DeliveryOption {
	c := ent.TxFromContext(ctx)

	serv := c.CarrierService.Query().
		Where(carrierservice.ReturnEQ(true)).
		FirstX(ctx)

	return c.DeliveryOption.Create().
		SetConnection(conn).
		SetCarrierID(carrierID).
		SetSortOrder(1).
		SetName("Return drop-off").
		SetTenantID(tenantID).
		SetCarrierService(serv).
		SaveX(ctx)
}
