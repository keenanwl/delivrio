package viewer

import (
	"context"
	"delivrio.io/shared-utils/pulid"
)

// Role for viewer actions.
type Role int

// List of roles.
const (
	_ Role = 1 << iota
	Background
	BackgroundForTenant // Used when the tenant should be included in queries for enforcing permissions
	Admin
	CustomerTenant
	Anonymous
)

// Viewer describes the query/mutation viewer-context.
type Viewer interface {
	Admin() bool // If viewer is admin.
	Background() bool
	CurrentRole() Role
	TenantID() pulid.ID
	MyId() pulid.ID
	ContextID() pulid.ID
	AnonymousOverride() bool
}

// UserViewer describes a user-viewer.
type UserViewer struct {
	Role    Role // Attached roles.
	Tenant  pulid.ID
	MyID    pulid.ID
	Context pulid.ID
	// Allows endpoints to break through privacy restrictions
	// for limited use cases like creating returns.
	// TLDR: more limited scope for overriding privacy
	LimitedAnonymousOverride bool
}

func (v UserViewer) Admin() bool {
	return v.Role&Admin != 0
}

func (v UserViewer) Background() bool {
	return v.Role&Background != 0
}
func (v UserViewer) CurrentRole() Role {
	return v.Role
}

func (v UserViewer) TenantID() pulid.ID {
	return v.Tenant
}

func (v UserViewer) MyId() pulid.ID {
	return v.MyID
}
func (v UserViewer) AnonymousOverride() bool {
	return v.LimitedAnonymousOverride
}

func (v UserViewer) ContextID() pulid.ID {
	if v.Context == "" {
		panic("context ID must be set")
	}
	return v.Context
}

type ctxKey struct{}

// FromContext returns the Viewer stored in a context.
func FromContext(ctx context.Context) Viewer {
	v, _ := ctx.Value(ctxKey{}).(Viewer)
	return v
}

// NewContext returns a copy of parent context with the given Viewer attached with it.
func NewContext(parent context.Context, v Viewer) context.Context {
	return context.WithValue(parent, ctxKey{}, v)
}

func MergeViewerContextID(parent context.Context, v Viewer) context.Context {
	current := FromContext(parent)
	nextUserViewer := v.(UserViewer)
	if current == nil {
		nextUserViewer.Context = pulid.MustNew("CH")
	} else {
		nextUserViewer.Context = current.ContextID()
	}

	return context.WithValue(parent, ctxKey{}, nextUserViewer)
}

func MergeViewerTenantID(parent context.Context, tenantID pulid.ID) context.Context {
	current := FromContext(parent)
	nextUserViewer := current.(UserViewer)
	nextUserViewer.Tenant = tenantID

	return context.WithValue(parent, ctxKey{}, nextUserViewer)
}

func EnableAnonymousOverride(parent context.Context) context.Context {
	current := FromContext(parent)
	nextUserViewer, ok := current.(UserViewer)
	if !ok {
		panic("override may only be used on current viewer ctx")
	} else {
		nextUserViewer.LimitedAnonymousOverride = true
	}

	return context.WithValue(parent, ctxKey{}, nextUserViewer)
}

func NewBackgroundContext(parent context.Context) context.Context {
	return NewContext(
		parent,
		UserViewer{
			Role:    Background,
			Context: pulid.MustNew("CH"),
		},
	)
}
