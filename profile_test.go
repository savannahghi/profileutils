package profileutils_test

import (
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

func TestUserProfile_HasFavorite(t *testing.T) {
	type args struct {
		title string
		user  profileutils.UserProfile
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid:user_has_favorite",
			args: args{
				title: "Home",
				user: profileutils.UserProfile{
					FavNavActions: []string{"Home"},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.user.HasFavorite(tt.args.title); got != tt.want {
				t.Errorf("UserProfile.HasFavorite() = %v, want %v", got, tt.want)
			}
		})
	}
}
