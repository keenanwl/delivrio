package background

import (
	"context"
	"delivrio.io/go/ent"
	"delivrio.io/go/ent/changehistory"
	"delivrio.io/go/ent/emailtemplate"
	"delivrio.io/go/ent/returncolli"
	"delivrio.io/go/ent/returnportal"
	"delivrio.io/go/mergeutils"
	"delivrio.io/go/schema/hooks/history"
	"delivrio.io/go/viewer"
	"fmt"
	"time"
)

func handleReturnStatusConfirmationLabelEmails(ctx context.Context) (int, []error) {
	db := ent.FromContext(ctx)
	count := 0
	allErrors := make([]error, 0)

	returnColliesToNotify, err := db.ReturnColli.Query().
		WithReturnPortal().
		Where(returncolli.And(
			// Check Delivery Option for more precision?
			returncolli.Or(returncolli.QrCodePng(""), returncolli.QrCodePngIsNil()),
			returncolli.EmailConfirmationLabelIsNil(),
			returncolli.StatusEQ(returncolli.StatusPending),
			returncolli.HasReturnPortalWith(
				returnportal.HasEmailConfirmationLabel(),
			))).
		Order(returncolli.ByID()).
		Limit(maxNotificationProcessed).
		All(ctx)
	if err != nil {
		return count, append(allErrors, fmt.Errorf("error handle return confirm label email: %w", err))
	}

	for _, rc := range returnColliesToNotify {
		// Required for the change history
		tenantCtx := viewer.MergeViewerTenantID(ctx, rc.TenantID)
		err = SendReturnColliEmail(tenantCtx, rc, emailtemplate.MergeTypeReturnColliLabel)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("error handle return confirm label email: %w", err))
			continue
		} else {
			err := rc.Update().
				SetEmailConfirmationLabel(time.Now()).
				Exec(history.NewConfig(tenantCtx).
					SetOrigin(changehistory.OriginBackground).
					SetDescription("Return confirm label email sent to customer").
					Ctx())
			if err != nil {
				allErrors = append(allErrors, fmt.Errorf("error handle return confirm label email: %w", err))
				continue
			}
		}
	}

	return count, allErrors
}

func handleReturnStatusConfirmationQRCodeEmails(ctx context.Context) (int, []error) {
	db := ent.FromContext(ctx)
	count := 0
	allErrors := make([]error, 0)

	returnColliesToNotify, err := db.ReturnColli.Query().
		WithReturnPortal().
		Where(returncolli.And(
			// Check Delivery Option for more precision?
			returncolli.Or(returncolli.QrCodePngNEQ(""), returncolli.QrCodePngNotNil()),
			returncolli.EmailConfirmationQrCodeIsNil(),
			returncolli.StatusEQ(returncolli.StatusPending),
			returncolli.HasReturnPortalWith(
				returnportal.HasEmailConfirmationQrCode(),
			))).
		Order(returncolli.ByID()).
		Limit(maxNotificationProcessed).
		All(ctx)
	if err != nil {
		return count, append(allErrors, fmt.Errorf("error handle return confirm QR email: %w", err))
	}

	for _, rc := range returnColliesToNotify {
		// Required for the change history
		tenantCtx := viewer.MergeViewerTenantID(ctx, rc.TenantID)
		err = SendReturnColliEmail(tenantCtx, rc, emailtemplate.MergeTypeReturnColliQr)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("error handle return confirm QR email: %w", err))
			continue
		} else {
			err := rc.Update().
				SetEmailConfirmationQrCode(time.Now()).
				Exec(history.NewConfig(tenantCtx).
					SetOrigin(changehistory.OriginBackground).
					SetDescription("Return confirm QR email sent to customer").
					Ctx())
			if err != nil {
				allErrors = append(allErrors, fmt.Errorf("error handle return confirm QR email: %w", err))
				continue
			}
		}
	}

	return count, allErrors
}

func handleReturnStatusReceivedEmails(ctx context.Context) (int, []error) {
	db := ent.FromContext(ctx)
	count := 0
	allErrors := make([]error, 0)

	returnColliesToNotify, err := db.ReturnColli.Query().
		WithReturnPortal().
		Where(returncolli.And(
			returncolli.EmailReceivedIsNil(),
			returncolli.StatusEQ(returncolli.StatusReceived),
			returncolli.HasReturnPortalWith(
				returnportal.HasEmailReceived(),
			))).
		Order(returncolli.ByID()).
		Limit(maxNotificationProcessed).
		All(ctx)
	if err != nil {
		return count, append(allErrors, fmt.Errorf("error handle return received email: %w", err))
	}

	for _, rc := range returnColliesToNotify {
		// Required for the change history
		tenantCtx := viewer.MergeViewerTenantID(ctx, rc.TenantID)
		err = SendReturnColliEmail(tenantCtx, rc, emailtemplate.MergeTypeReturnColliReceived)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("error handle return received email: %w", err))
			continue
		} else {
			err := rc.Update().
				SetEmailReceived(time.Now()).
				Exec(history.NewConfig(tenantCtx).
					SetOrigin(changehistory.OriginBackground).
					SetDescription("Return received email sent to customer").
					Ctx())
			if err != nil {
				allErrors = append(allErrors, fmt.Errorf("error handle return received email: %w", err))
				continue
			}
		}
	}

	return count, allErrors
}

func handleReturnStatusAcceptedEmails(ctx context.Context) (int, []error) {
	db := ent.FromContext(ctx)
	count := 0
	allErrors := make([]error, 0)

	returnColliesToNotify, err := db.ReturnColli.Query().
		WithReturnPortal().
		Where(returncolli.And(
			returncolli.EmailAcceptedIsNil(),
			returncolli.StatusEQ(returncolli.StatusAccepted),
			returncolli.HasReturnPortalWith(
				returnportal.HasEmailAccepted(),
			))).
		Order(returncolli.ByID()).
		Limit(maxNotificationProcessed).
		All(ctx)
	if err != nil {
		return count, append(allErrors, fmt.Errorf("error handle return accepted email: %w", err))
	}

	for _, rc := range returnColliesToNotify {
		// Required for the change history
		tenantCtx := viewer.MergeViewerTenantID(ctx, rc.TenantID)
		err = SendReturnColliEmail(tenantCtx, rc, emailtemplate.MergeTypeReturnColliAccepted)
		if err != nil {
			allErrors = append(allErrors, fmt.Errorf("error handle return accepted email: %w", err))
			continue
		} else {
			err := rc.Update().
				SetEmailAccepted(time.Now()).
				Exec(history.NewConfig(tenantCtx).
					SetOrigin(changehistory.OriginBackground).
					SetDescription("Return accepted email sent to customer").
					Ctx())
			if err != nil {
				allErrors = append(allErrors, fmt.Errorf("error handle return accepted email: %w", err))
				continue
			}
		}
	}

	return count, allErrors
}

func SendReturnColliEmail(ctx context.Context, returnColli *ent.ReturnColli, mergeType emailtemplate.MergeType) error {
	recip, err := returnColli.Recipient(ctx)
	if err != nil {
		return err
	}

	mergeValues, err := returnEmailValues(ctx, returnColli, mergeType)
	if err != nil {
		return err
	}

	rp, err := returnColli.ReturnPortal(ctx)
	if err != nil {
		return err
	}

	tmpl, err := returnPortalEmailTemplate(ctx, rp, mergeType)
	if err != nil {
		return err
	}

	_, err = mergeutils.SendTransactionalEmail(
		tmpl.HTMLTemplate,
		tmpl.Subject,
		recip.Email,
		mergeValues,
	)
	if err != nil {
		return err
	}

	return nil
}

func returnEmailValues(ctx context.Context, returnColli *ent.ReturnColli, mergeType emailtemplate.MergeType) (interface{}, error) {
	switch mergeType {
	case emailtemplate.MergeTypeReturnColliLabel:
		return mergeutils.ReturnColliConfirmationLabelMerge(ctx, returnColli)
	case emailtemplate.MergeTypeReturnColliQr:
		return mergeutils.ReturnColliConfirmationQRMerge(ctx, returnColli)
	case emailtemplate.MergeTypeReturnColliReceived:
		return mergeutils.ReturnColliReceivedMerge(ctx, returnColli)
	case emailtemplate.MergeTypeReturnColliAccepted:
		return mergeutils.ReturnColliAcceptedMerge(ctx, returnColli)
	default:
		return nil, fmt.Errorf("return status email: unknown merge type: %s\n", mergeType)
	}
}

func returnPortalEmailTemplate(ctx context.Context, rp *ent.ReturnPortal, mergeType emailtemplate.MergeType) (*ent.EmailTemplate, error) {
	switch mergeType {
	case emailtemplate.MergeTypeReturnColliLabel:
		return rp.EmailConfirmationLabel(ctx)
	case emailtemplate.MergeTypeReturnColliQr:
		return rp.EmailConfirmationQrCode(ctx)
	case emailtemplate.MergeTypeReturnColliReceived:
		return rp.EmailReceived(ctx)
	case emailtemplate.MergeTypeReturnColliAccepted:
		return rp.EmailAccepted(ctx)
	}
	return nil, fmt.Errorf("return status email: unknown merge type: %s\n", mergeType)
}
