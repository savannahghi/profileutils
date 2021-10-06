package mock

import (
	"context"

	"github.com/savannahghi/profileutils"
)

// FakeUserProfileRepository contains the profile repository mocks
type FakeUserProfileRepository struct {
	GetLoggedInUserFn func(ctx context.Context) (*profileutils.UserInfo, error)
}

// GetLoggedInUser is a mock of GetLoggedInUser method
func (f *FakeUserProfileRepository) GetLoggedInUser(ctx context.Context) (*profileutils.UserInfo, error) {
	return f.GetLoggedInUserFn(ctx)
}
