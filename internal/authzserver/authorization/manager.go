package authorization

import (
	"context"

	"github.com/marmotedu/errors"
	"github.com/ory/ladon"
)

// PolicyManager is a mysql implementation for Manager to store
// policies persistently.
type PolicyManager struct {
	client AuthorizationInterface
}

// NewPolicyManager initializes a new PolicyManager for given apimachinery api
// client.
func NewPolicyManager(client AuthorizationInterface) ladon.Manager {
	return &PolicyManager{
		client: client,
	}
}

// Create persists the policy.
func (*PolicyManager) Create(ctx context.Context, policy ladon.Policy) error {
	return nil
}

// Update updates an existing policy.
func (*PolicyManager) Update(ctx context.Context, policy ladon.Policy) error {
	return nil
}

// Get retrieves a policy.
func (*PolicyManager) Get(ctx context.Context, id string) (ladon.Policy, error) {
	return nil, nil
}

// Delete removes a policy.
func (*PolicyManager) Delete(ctx context.Context, id string) error {
	return nil
}

// GetAll retrieves all policies.
func (*PolicyManager) GetAll(ctx context.Context, limit, offset int64) (ladon.Policies, error) {
	return nil, nil
}

// FindRequestCandidates returns candidates that could match the request object. It either returns
// a set that exactly matches the request, or a superset of it. If an error occurs, it returns nil and
// the error.
func (m *PolicyManager) FindRequestCandidates(ctx context.Context, r *ladon.Request) (ladon.Policies, error) {
	username := ""

	if user, ok := r.Context["username"].(string); ok {
		username = user
	}

	policies, err := m.client.List(username)
	if err != nil {
		return nil, errors.Wrap(err, "list policies failed")
	}

	ret := make([]ladon.Policy, 0, len(policies))
	for _, policy := range policies {
		ret = append(ret, policy)
	}

	return ret, nil
}

// FindPoliciesForSubject returns policies that could match the subject. It either returns
// a set of policies that applies to the subject, or a superset of it.
// If an error occurs, it returns nil and the error.
func (m *PolicyManager) FindPoliciesForSubject(ctx context.Context, subject string) (ladon.Policies, error) {
	return nil, nil
}

// FindPoliciesForResource returns policies that could match the resource. It either returns
// a set of policies that apply to the resource, or a superset of it.
// If an error occurs, it returns nil and the error.
func (m *PolicyManager) FindPoliciesForResource(ctx context.Context, resource string) (ladon.Policies, error) {
	return nil, nil
}
