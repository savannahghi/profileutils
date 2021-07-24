package profileutils

import (
	"reflect"
	"testing"
)

func TestGetPermissionByOperation(t *testing.T) {
	type args struct {
		operation string
	}
	tests := []struct {
		name    string
		args    args
		want    *Permission
		wantNil bool
	}{
		{
			name:    "happy case",
			args:    args{"role.create"},
			want:    &CanCreateRole,
			wantNil: false,
		},
		{
			name:    "sad case",
			args:    args{"some.random.unexisting.operation"},
			want:    nil,
			wantNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPermissionByOperation(tt.args.operation)
			if tt.wantNil {
				if got != nil {
					t.Errorf("GetPermissionByOperation() = %v, wantNil %v", got, tt.wantNil)
				}
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetPermissionByOperation() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
