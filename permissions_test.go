package profileutils_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/savannahghi/profileutils"
)

func TestUniquePermissions(t *testing.T) {
	ctx := context.Background()
	p, err := profileutils.AllPermissions(ctx)
	if err != nil {
		t.Errorf("cannot retrieve all permissions, err:%v", err)
		return
	}

	duplicate := []profileutils.Permission{
		{
			Scope: "test.run",
		},
		{
			Scope: "test.run",
		},
	}
	type args struct {
		ctx         context.Context
		permissions []profileutils.Permission
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid: all permissions are unique code",
			args: args{
				ctx:         ctx,
				permissions: p,
			},
			wantErr: false,
		},
		{
			name: "invalid: duplicate code in permissions",
			args: args{
				ctx:         ctx,
				permissions: duplicate,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := profileutils.UniquePermissions(tt.args.ctx, tt.args.permissions); (err != nil) != tt.wantErr {
				t.Errorf("UniquePermissions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_GetPermissionByScope(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx   context.Context
		scope string
	}
	tests := []struct {
		name    string
		args    args
		want    *profileutils.Permission
		wantErr bool
	}{
		{
			name: "valid: retrieve a single permission",
			args: args{
				ctx:   ctx,
				scope: "role.view",
			},
			want:    &profileutils.CanViewRole,
			wantErr: false,
		},
		{
			name: "invalid: retrieve an invalid permission",
			args: args{
				ctx:   ctx,
				scope: "ludicrous",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := profileutils.GetPermissionByScope(tt.args.ctx, tt.args.scope)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPermission() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRole_Permissions(t *testing.T) {
	ctx := context.Background()
	type args struct {
		ctx  context.Context
		role profileutils.Role
	}
	tests := []struct {
		name    string
		args    args
		want    []profileutils.Permission
		wantErr bool
	}{
		{
			name: "happy got permissions",
			args: args{
				ctx: ctx,
				role: profileutils.Role{
					Scopes: []string{"role.edit"},
				},
			},
			want:    []profileutils.Permission{profileutils.CanEditRole},
			wantErr: false,
		},
		{
			name: "sad did not get permissions",
			args: args{
				ctx: ctx,
				role: profileutils.Role{
					Scopes: []string{"role.submit.unknown"},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.role.Permissions(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Role.Permissions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Role.Permissions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermissionGroup_IsValid(t *testing.T) {

	tests := []struct {
		name string
		p    profileutils.PermissionGroup
		want bool
	}{
		{
			name: "happy case",
			p:    "Roles",
			want: true,
		},

		{
			name: "happy case",
			p:    "Vendors",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsValid(); got != tt.want {
				t.Errorf("PermissionGroup.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
