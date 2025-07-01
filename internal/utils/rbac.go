package utils

import (
	"context"
)

type RBAC struct {
	ctx  context.Context
	role string
}

func NewRBAC() *RBAC { return new(RBAC) }

func DefaultRBAC(ctx context.Context, role string) *RBAC {
	return NewRBAC().WithContext(ctx).WithRole(role)
}

func (r *RBAC) WithContext(ctx context.Context) *RBAC {
	r.ctx = ctx
	return r
}

func (r *RBAC) WithRole(role string) *RBAC {
	r.role = role
	return r
}

func (r *RBAC) IsAuthorized() bool {
	if r.ctx == nil {
		return true
	}

	if ValueInSlice(r.role, r.ctx.Value(ContextKeyUserRoles).([]string)) {
		return true
	}

	return false
}
