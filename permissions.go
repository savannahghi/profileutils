package profileutils

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"time"
)

// PermissionGroups for the various resources
const (
	PermissionGroupRole PermissionGroup = "Roles"

	PermissionGroupEmployee PermissionGroup = "Employees"

	PermissionGroupAgent PermissionGroup = "Agents"

	PermissionGroupPartner PermissionGroup = "Partners"

	PermissionGroupKYC PermissionGroup = "KYC"

	PermissionGroupCovers PermissionGroup = "Covers"

	PermissionGroupConsumer PermissionGroup = "Consumers"

	PermissionGroupPatient PermissionGroup = "Patients"
)

// PermissionGroup used to group permissions that have related resources
type PermissionGroup string

// IsValid ..
func (p PermissionGroup) IsValid() bool {
	switch p {
	case PermissionGroupRole,
		PermissionGroupEmployee,
		PermissionGroupAgent,
		PermissionGroupPartner,
		PermissionGroupKYC,
		PermissionGroupConsumer,
		PermissionGroupCovers,
		PermissionGroupPatient:
		return true
	}
	return false
}

//String ...
func (p PermissionGroup) String() string {
	return string(p)
}

// UnmarshalGQL ..
func (p *PermissionGroup) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*p = PermissionGroup(str)
	if !p.IsValid() {
		return fmt.Errorf("%s is not a valid PermissionGroup", str)
	}
	return nil
}

// MarshalGQL ..
func (p PermissionGroup) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(p.String()))
}

// role management permissions
var (
	CanViewRole = Permission{
		Group:       PermissionGroupRole.String(),
		Scope:       "role.view",
		Description: "Can view role",
	}

	CanCreateRole = Permission{
		Group:       PermissionGroupRole.String(),
		Scope:       "role.create",
		Description: "Can create role",
	}

	CanEditRole = Permission{
		Group:       PermissionGroupRole.String(),
		Scope:       "role.edit",
		Description: "Can edit role",
	}

	CanAssignRole = Permission{
		Group:       PermissionGroupRole.String(),
		Scope:       "role.assign",
		Description: "Can assign a role",
	}
)

// employee management permissions
var (
	CanViewEmployee = Permission{
		Group:       PermissionGroupEmployee.String(),
		Scope:       "employee.view",
		Description: "Can create employee",
	}

	CanCreateEmployee = Permission{
		Group:       PermissionGroupEmployee.String(),
		Scope:       "employee.create",
		Description: "Can create employee",
	}

	CanRemoveEmployee = Permission{
		Group:       PermissionGroupEmployee.String(),
		Scope:       "employee.remove",
		Description: "Can remove employee",
	}
)

// agent management permissions
var (
	CanViewAgent = Permission{
		Group:       PermissionGroupAgent.String(),
		Scope:       "agent.view",
		Description: "Can view agents",
	}

	CanRegisterAgent = Permission{
		Group:       PermissionGroupAgent.String(),
		Scope:       "agent.register",
		Description: "Can register agent",
	}

	CanIdentifyAgent = Permission{
		Group:       PermissionGroupAgent.String(),
		Scope:       "agent.identify",
		Description: "Can identify agent",
	}

	CanSuspendAgent = Permission{
		Group:       PermissionGroupAgent.String(),
		Scope:       "agent.suspend",
		Description: "Can suspend agent",
	}

	CanUnsuspendAgent = Permission{
		Group:       PermissionGroupAgent.String(),
		Scope:       "agent.unsuspend",
		Description: "Can unsuspend agent",
	}
)

// partner management permissions
var (
	CanViewPartner = Permission{
		Group:       PermissionGroupPartner.String(),
		Scope:       "partner.view",
		Description: "Can view partners",
	}

	CanCreatePartner = Permission{
		Group:       PermissionGroupPartner.String(),
		Scope:       "partner.create",
		Description: "Can create partner",
	}
)

// kyc management permissions
var (
	CanProcessKYC = Permission{
		Group:       PermissionGroupKYC.String(),
		Scope:       "kyc.process",
		Description: "Can process KYC",
	}
	CanViewKYC = Permission{
		Group:       PermissionGroupKYC.String(),
		Scope:       "kyc.view",
		Description: "Can process KYC",
	}
)

// consumer management permissions
var (
	CanViewConsumers = Permission{
		Group:       PermissionGroupConsumer.String(),
		Scope:       "consumer.view",
		Description: "Can view consumers",
	}

	CanCreateConsumer = Permission{
		Group:       PermissionGroupConsumer.String(),
		Scope:       "consumer.create",
		Description: "Can add consumer",
	}
)

// cover management permissions
var (
	CanViewCovers = Permission{
		Group:       PermissionGroupCovers.String(),
		Scope:       "covers.view",
		Description: "Can view covers",
	}

	CanSellCovers = Permission{
		Group:       PermissionGroupCovers.String(),
		Scope:       "covers.sell",
		Description: "Can sell covers",
	}
)

// patient management permissions
var (
	CanViewPatient = Permission{
		Group:       PermissionGroupPatient.String(),
		Scope:       "patient.view",
		Description: "Can view patient",
	}

	CanCreatePatient = Permission{
		Group:       PermissionGroupPatient.String(),
		Scope:       "patient.create",
		Description: "Can add patient",
	}

	CanIdentifyPatient = Permission{
		Group:       PermissionGroupPatient.String(),
		Scope:       "patient.identify",
		Description: "Can identify patient",
	}
)

// Permission defines an approval of a mode of access to a resource
type Permission struct {
	// Scope is the resource and action
	//
	// It should be in the form "<resource>.<action>"
	// Example: "user.create"
	Scope string `json:"scope"`

	// Description is used to keep details for a particular permission
	Description string `json:"description"`

	Group string `json:"group"`

	Allowed bool `json:"allowed"`
}

// AllPermissions returns all the defined permissions
func AllPermissions(ctx context.Context) ([]Permission, error) {
	permissions := []Permission{
		//role management
		CanViewRole,
		CanCreateRole,
		CanEditRole,
		CanAssignRole,

		//employee management
		CanViewEmployee,
		CanCreateEmployee,
		CanRemoveEmployee,

		// agent management
		CanViewAgent,
		CanRegisterAgent,
		CanSuspendAgent,
		CanIdentifyAgent,
		CanUnsuspendAgent,

		// partner management
		CanViewPartner,
		CanCreatePartner,

		// KYC management
		CanProcessKYC,
		CanViewKYC,

		//covers management
		CanViewCovers,
		CanSellCovers,

		// consumer management
		CanViewConsumers,
		CanCreateConsumer,

		// patient management
		CanViewPatient,
		CanCreatePatient,
		CanIdentifyPatient,
	}

	return permissions, nil
}

// Role is the user group defining permissions
type Role struct {
	// Unique identifier for a role
	ID string `json:"id" firestore:"id"`

	// Organization ID for which the role is created and used
	OrganizationID string `json:"organizationID,omitempty" firestore:"organizationID"`

	// Name given to the role
	Name string `json:"name" firestore:"name"`

	// Description is used to keep details for a particular role
	Description string `json:"description" firestore:"description"`

	// Active indicates whether the role is operational
	Active bool `json:"active" firestore:"active"`

	// List of allowed permission scopes for a role
	Scopes []string `json:"scopes" firestore:"scopes"`

	// CreatedBy is the Profile ID of the user creating the role.
	CreatedBy string `json:"createdBy,omitempty" firestore:"createdBy"`

	// Created is the timestamp indicating when the role was created
	Created time.Time `json:"created" firestore:"created"`

	// UpdatedBy is the Profile ID of the user who last updated the role.
	UpdatedBy string `json:"updatedBy,omitempty" firestore:"updatedBy"`

	// Updated is the timestamp indicating when the role was last updated
	Updated time.Time `json:"updated" firestore:"updated"`
}

// Permissions returns all permissions with role scoped marked as allowed
func (r Role) Permissions(ctx context.Context) ([]Permission, error) {
	perms := []Permission{}

	allPermissions, err := AllPermissions(ctx)
	if err != nil {
		return nil, err
	}

	for _, perm := range allPermissions {
		// check if this permission is among the role's scopes
		// and mark is as allowed
		for _, scope := range r.Scopes {
			if perm.Scope == scope {
				perm.Allowed = true
			}
		}
		perms = append(perms, perm)
	}

	return perms, nil
}

// HasPermission checks if a role has the permission defined
func (r Role) HasPermission(ctx context.Context, scope string) bool {
	for _, op := range r.Scopes {
		if op == scope {
			return true
		}
	}
	return false
}

// UniquePermissions validates that each permission has a unique code
// The validation has a test against the declared perms to ensure no permission duplication
func UniquePermissions(ctx context.Context, permissions []Permission) error {
	occurred := map[string]bool{}
	for _, perm := range permissions {
		// check if already the mapped
		if !occurred[perm.Scope] {
			occurred[perm.Scope] = true
		} else {
			return fmt.Errorf("permission exist, permission scope: %v", perm.Scope)
		}
	}

	return nil
}

// GetUniquePermissions return s a list of unique permissions from list given
func GetUniquePermissions(ctx context.Context, permissions []Permission) []Permission {
	occurred := map[string]bool{}
	result := []Permission{}

	for _, perm := range permissions {
		if !occurred[perm.Scope] {
			result = append(result, perm)
		}
	}

	return result
}

// GetPermissionByScope retrieves a single permission using its scope
func GetPermissionByScope(ctx context.Context, scope string) (*Permission, error) {
	permissions, err := AllPermissions(ctx)
	if err != nil {
		return nil, err
	}

	var p *Permission
	for _, permission := range permissions {
		if permission.Scope == scope {
			p = &permission
			break
		}
	}

	if p == nil {
		return nil, fmt.Errorf("permission with the scope %v not found", scope)
	}

	return p, nil

}
