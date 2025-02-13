package policy

import (
	srvv1 "j-iam/internal/apiserver/service/v1"
	"j-iam/internal/apiserver/store"
)

// PolicyController create a policy handler used to handle request for policy resource.
type PolicyController struct {
	srv srvv1.Service
}

// NewPolicyController creates a policy handler.
func NewPolicyController(store store.Factory) *PolicyController {
	return &PolicyController{
		srv: srvv1.NewService(store),
	}
}
