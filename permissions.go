package profileutils

import (
	"context"
	"fmt"
	"time"
)

// Permission defines an approval of a mode of access to a resource
type Permission struct {
	// Code acts as a unique attribute for a permission
	Code string `json:"code" firestore:"code"`

	// Resource represents the object on which access acts upon
	Resource string `json:"resource" firestore:"resource"`

	// Action is the right one has on a resource
	Action string `json:"action" firestore:"action"`

	// Description is used to keep details for a particular permission
	Description string `json:"description" firestore:"description"`

	Allowed bool `json:"allowed" firestore:"allowed"`
}

// String returns the operation formated as a string
//
// The operation is formatted as "{resource}.{action}"
// It can be used by clients which are required to b permissions
func (p Permission) String() string {
	return fmt.Sprintf("%v.%v", p.Resource, p.Action)
}

// Role is the user group defining permissions
type Role struct {
	// Unique identifier for a role
	ID string `json:"id" firestore:"id"`

	// Name given to the role
	Name string `json:"name" firestore:"name"`

	// Description is used to keep details for a particular role
	Description string `json:"description" firestore:"description"`

	// Profile ID of the user creating the role.
	ProfileID string `json:"profileID,omitempty" firestore:"profileID"`

	// Organization ID for which the role is created and used
	OrganizationID string `json:"organizationID,omitempty" firestore:"organizationID"`

	// Active indicates whether the role is operational
	Active bool `json:"active" firestore:"active"`

	// Created is the timestamp indicating when the role was created
	Created time.Time `json:"created" firestore:"created"`

	// Permissions
	Permissions []Permission `json:"permissions" firestore:"permissions"`
}

// GetPermissions returns the permissions for a role formatted as a string
func (r Role) GetPermissions() []string {
	perms := []string{}

	for _, perm := range r.Permissions {
		perms = append(perms, perm.String())
	}

	return perms
}

// Role resource related permissions
var (
	CreateRole Permission = Permission{
		Code:        "1",
		Resource:    "role",
		Action:      "create",
		Description: "Can add a role",
	}

	DeleteRole Permission = Permission{
		Code:        "2",
		Resource:    "role",
		Action:      "delete",
		Description: "Can remove a role",
	}

	UpdateRole Permission = Permission{
		Code:        "3",
		Resource:    "role",
		Action:      "update",
		Description: "Can edit a role",
	}

	ViewRole Permission = Permission{
		Code:        "4",
		Resource:    "role",
		Action:      "view",
		Description: "Can view a role",
	}
)

// Agent resource related permissions
var (
	CreateAgent Permission = Permission{
		Code:        "5",
		Resource:    "agent",
		Action:      "create",
		Description: "Can add an agent",
	}

	DeleteAgent Permission = Permission{
		Code:        "6",
		Resource:    "agent",
		Action:      "delete",
		Description: "Can remove an agent",
	}

	UpdateAgent Permission = Permission{
		Code:        "7",
		Resource:    "agent",
		Action:      "update",
		Description: "Can edit an agent",
	}

	ViewAgent Permission = Permission{
		Code:        "8",
		Resource:    "agent",
		Action:      "view",
		Description: "Can view an agent",
	}
)

// Employee resource related permissions
var (
	CreateEmployee Permission = Permission{
		Code:        "9",
		Resource:    "employee",
		Action:      "create",
		Description: "Can add an employee",
	}

	DeleteEmployee Permission = Permission{
		Code:        "10",
		Resource:    "employee",
		Action:      "delete",
		Description: "Can remove an employee",
	}

	UpdateEmployee Permission = Permission{
		Code:        "11",
		Resource:    "employee",
		Action:      "update",
		Description: "Can edit an employee",
	}

	ViewEmployee Permission = Permission{
		Code:        "12",
		Resource:    "employee",
		Action:      "view",
		Description: "Can view an employee",
	}
)

// Consumer resource related permissions
var (
	CreateConsumer Permission = Permission{
		Code:        "13",
		Resource:    "consumer",
		Action:      "create",
		Description: "Can add an consumer",
	}

	DeleteConsumer Permission = Permission{
		Code:        "14",
		Resource:    "consumer",
		Action:      "delete",
		Description: "Can remove an consumer",
	}

	UpdateConsumer Permission = Permission{
		Code:        "15",
		Resource:    "consumer",
		Action:      "update",
		Description: "Can edit an consumer",
	}

	ViewConsumer Permission = Permission{
		Code:        "16",
		Resource:    "consumer",
		Action:      "view",
		Description: "Can view an consumer",
	}
)

// Patient resource related permissions
var (
	CreatePatient Permission = Permission{
		Code:        "17",
		Resource:    "patient",
		Action:      "create",
		Description: "Can add a patient",
	}

	DeletePatient Permission = Permission{
		Code:        "18",
		Resource:    "patient",
		Action:      "delete",
		Description: "Can remove a patient",
	}

	UpdatePatient Permission = Permission{
		Code:        "19",
		Resource:    "patient",
		Action:      "update",
		Description: "Can edit a patient",
	}

	ViewPatient Permission = Permission{
		Code:        "20",
		Resource:    "patient",
		Action:      "view",
		Description: "Can view a patient",
	}
)

// KYC resource related permissions
var (
	CreateKYC Permission = Permission{
		Code:        "21",
		Resource:    "kyc",
		Action:      "create",
		Description: "Can add a kyc",
	}

	UpdateKYC Permission = Permission{
		Code:        "22",
		Resource:    "kyc",
		Action:      "update",
		Description: "Can edit a kyc",
	}

	ViewKYC Permission = Permission{
		Code:        "23",
		Resource:    "kyc",
		Action:      "view",
		Description: "Can view a kyc",
	}
)

// AllPermissions returns all the defined permissions
func AllPermissions(ctx context.Context) ([]Permission, error) {
	permissions := []Permission{
		CreateRole, UpdateRole, ViewRole, DeleteRole,
		CreateAgent, UpdateAgent, ViewAgent, DeleteAgent,
		CreateEmployee, UpdateEmployee, ViewEmployee, DeleteEmployee,
		CreatePatient, UpdatePatient, ViewPatient, DeletePatient,
		CreateKYC, UpdateKYC, ViewKYC,
	}
	// validate unique code for each permission
	// look into improvements on "code" creation // human error risk

	return permissions, nil
}

// UniquePermissions validates that each permission has a unique code
//
// The validation has a test against the declared perms to ensure no variable duplication
func UniquePermissions(ctx context.Context, permissions []Permission) error {
	occurred := map[string]bool{}
	for _, perm := range permissions {
		// check if already the mapped
		if !occurred[perm.Code] {
			occurred[perm.Code] = true
		} else {
			return fmt.Errorf("permission exist,permission code:%v", perm.Code)
		}
	}

	return nil
}

// GetPermission retrieves a single permission using its code
func GetPermission(ctx context.Context, code string) (*Permission, error) {
	permissions, err := AllPermissions(ctx)
	if err != nil {
		return nil, err
	}

	var p *Permission
	for _, permission := range permissions {
		if permission.Code == code {
			p = &permission
			break
		}
	}

	if p == nil {
		return nil, fmt.Errorf("permission not found")
	}

	return p, nil

}
