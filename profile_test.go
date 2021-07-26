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
		profileutils.PermissionTypeRegisterAgent,
		profileutils.PermissionTypeSuspendAgent,
		profileutils.PermissionTypeUnsuspendAgent,
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

func TestUserProfile_HasPermission(t *testing.T) {
	user := profileutils.UserProfile{
		Permissions: profileutils.DefaultEmployeePermissions,
	}
	user2 := profileutils.UserProfile{
		Permissions: profileutils.DefaultAgentPermissions,
	}
	tests := []struct {
		name string
		user profileutils.UserProfile
		perm profileutils.PermissionType
		want bool
	}{
		{
			name: "valid: user has permission",
			user: user,
			perm: profileutils.PermissionTypeRegisterAgent,
			want: true,
		},
		{
			name: "valid: user do no have permission",
			user: user2,
			perm: profileutils.PermissionTypeRegisterAgent,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.HasPermission(tt.perm); got != tt.want {
				t.Errorf("UserProfile.HasPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserProfile_HasRole(t *testing.T) {
	type args struct {
		roleId string
		user   profileutils.UserProfile
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy",
			args: args{
				user: profileutils.UserProfile{
					Roles: []string{"sdkdasjhghjsd", "1212bhjbv"},
				},
				roleId: "sdkdasjhghjsd",
			},
			want: true,
		},
		{
			name: "sad",
			args: args{
				user: profileutils.UserProfile{
					Roles: []string{"sdkdasjhghjsd", "1212bhjbv"},
				},
				roleId: "123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.user.HasRole(tt.args.roleId); got != tt.want {
				t.Errorf("UserProfile.HasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}
