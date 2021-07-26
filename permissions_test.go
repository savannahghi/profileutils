package profileutils_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/savannahghi/profileutils"
)

func TestPermission_String(t *testing.T) {
	type fields struct {
		Code        string
		Resource    string
		Action      string
		Description string
		Allowed     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success: format permission",
			fields: fields{
				Code:        "001",
				Resource:    "role",
				Action:      "create",
				Description: "Can add a role",
			},
			want: "role.create",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := profileutils.Permission{
				Code:        tt.fields.Code,
				Resource:    tt.fields.Resource,
				Action:      tt.fields.Action,
				Description: tt.fields.Description,
				Allowed:     tt.fields.Allowed,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Permission.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniquePermissions(t *testing.T) {
	ctx := context.Background()
	p, err := profileutils.AllPermissions(ctx)
	if err != nil {
		t.Errorf("cannot retrieve all permissions, err:%v", err)
		return
	}

	duplicate := []profileutils.Permission{
		{
			Code:     "001",
			Resource: "dupe",
			Action:   "create",
		},
		{
			Code:     "001",
			Resource: "dupe",
			Action:   "create",
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

func TestGetPermission(t *testing.T) {
	ctx := context.Background()

	code := profileutils.CreateAgent.Code
	perm := profileutils.CreateAgent

	type args struct {
		ctx  context.Context
		code string
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
				ctx:  ctx,
				code: code,
			},
			want:    &perm,
			wantErr: false,
		},
		{
			name: "invalid: retrieve an invalid permission",
			args: args{
				ctx:  ctx,
				code: "ludicrous",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := profileutils.GetPermission(tt.args.ctx, tt.args.code)
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

func TestRole_GetPermissions(t *testing.T) {
	type fields struct {
		Permissions []profileutils.Permission
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "valid: string permissions",
			fields: fields{
				Permissions: []profileutils.Permission{
					{
						Resource: "test",
						Action:   "build",
					},
					{
						Resource: "test",
						Action:   "run",
					},
				},
			},
			want: []string{"test.build", "test.run"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := profileutils.Role{
				Permissions: tt.fields.Permissions,
			}
			if got := r.GetPermissions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Role.GetPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
