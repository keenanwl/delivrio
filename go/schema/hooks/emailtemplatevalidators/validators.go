package emailtemplatevalidators

import (
	"context"
	"delivrio.io/go/mergeutils"
	"fmt"

	ent2 "delivrio.io/go/ent"
	"delivrio.io/go/ent/emailtemplate"
	"delivrio.io/go/ent/hook"
	"entgo.io/ent"
)

func CreateUpdateEmailTemplate() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.EmailTemplateFunc(func(ctx context.Context, m *ent2.EmailTemplateMutation) (ent.Value, error) {

			temp, tempExists := m.HTMLTemplate()
			if tempExists {
				var err error
				dataType, exists := m.MergeType()
				if !exists {
					dataType, err = m.OldMergeType(ctx)
					if err != nil {
						return nil, err
					}
				}

				switch dataType {
				case emailtemplate.MergeTypeOrderConfirmation:
					_, err := mergeutils.MergeTemplate(temp, mergeutils.NewTestOrderConfirmation(ctx))
					if err != nil {
						return nil, fmt.Errorf("merge type %v: %w", dataType, err)
					}
				case emailtemplate.MergeTypeReturnColliLabel:
					_, err := mergeutils.MergeTemplate(temp, mergeutils.NewTestReturnConfirmationLabel(ctx))
					if err != nil {
						return nil, fmt.Errorf("merge type %v: %w", dataType, err)
					}
				case emailtemplate.MergeTypeReturnColliQr:
					_, err := mergeutils.MergeTemplate(temp, mergeutils.NewTestReturnConfirmationQRCode(ctx))
					if err != nil {
						return nil, fmt.Errorf("merge type %v: %w", dataType, err)
					}
				case emailtemplate.MergeTypeReturnColliReceived:
					_, err := mergeutils.MergeTemplate(temp, mergeutils.NewTestReturnReceived(ctx))
					if err != nil {
						return nil, fmt.Errorf("merge type %v: %w", dataType, err)
					}
				case emailtemplate.MergeTypeReturnColliAccepted:
					_, err := mergeutils.MergeTemplate(temp, mergeutils.NewTestReturnAccepted(ctx))
					if err != nil {
						return nil, fmt.Errorf("merge type %v: %w", dataType, err)
					}
				case emailtemplate.MergeTypeOrderPicked:
					_, err := mergeutils.MergeTemplate(temp, mergeutils.NewTestColliPacked(ctx))
					if err != nil {
						return nil, fmt.Errorf("merge type %v: %w", dataType, err)
					}
				default:
					return nil, fmt.Errorf("merge type %s not supported", dataType)
				}

			}

			return next.Mutate(ctx, m)
		})
	}
	return hook.On(hk, ent.OpUpdateOne|ent.OpCreate)
}
