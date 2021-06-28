package profileutils_test

import (
	"reflect"
	"testing"

	"github.com/savannahghi/profileutils"
)

func TestUserProfile_IsEntity(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "default case - just checking that the profile is marked as an entity",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := profileutils.UserProfile{}
			u.IsEntity()
		})
	}
}

func TestCover_IsEntity(t *testing.T) {
	type fields struct {
		PayerName      string
		PayerSladeCode int
		MemberNumber   string
		MemberName     string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "default case - just checking that the cover is marked as an entity",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := profileutils.Cover{}
			c.IsEntity()
		})
	}
}

func TestRoleType_Permissions(t *testing.T) {
	employeePermissions := []profileutils.PermissionType{
		profileutils.PermissionTypeCreateConsumer,
		profileutils.PermissionTypeUpdateConsumer,
		profileutils.PermissionTypeDeleteConsumer,
		profileutils.PermissionTypeCreatePatient,
		profileutils.PermissionTypeUpdatePatient,
		profileutils.PermissionTypeDeletePatient,
		profileutils.PermissionTypeIdentifyPatient,
	}
	agentPermissions := []profileutils.PermissionType{
		profileutils.PermissionTypeCreatePartner,
		profileutils.PermissionTypeUpdatePartner,
		profileutils.PermissionTypeCreateConsumer,
		profileutils.PermissionTypeUpdateConsumer,
	}

	tests := []struct {
		name    string
		r       profileutils.RoleType
		want    []profileutils.PermissionType
		wantErr bool
	}{
		{
			name: "valid role type permissions",
			r:    profileutils.RoleTypeEmployee,
			want: employeePermissions,
		},
		{
			name: "valid role type permissions",
			r:    profileutils.RoleTypeAgent,
			want: agentPermissions,
		},
		{
			name: "invalid role type permissions",
			r:    "IMPOSTER",
			want: []profileutils.PermissionType{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.Permissions()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoleType.Permissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoleType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		r    profileutils.RoleType
		want bool
	}{
		{
			name: "valid employee role type",
			r:    profileutils.RoleTypeEmployee,
			want: true,
		},
		{
			name: "valid agent role type",
			r:    profileutils.RoleTypeAgent,
			want: true,
		},
		{
			name: "invalid role type",
			r:    "IMPOSTER",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.IsValid(); got != tt.want {
				t.Errorf("RoleType.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
