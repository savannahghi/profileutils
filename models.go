package profileutils

import (
	"time"

	"github.com/savannahghi/firebasetools"
)

// PermissionGroup used to group permissions that have related resources
type PermissionGroup string

// ActionGroup is the codes for grouping related navigationactions
type ActionGoup string

//NavigationAction is a Navigation Action that a user can perform on the app
type NavigationAction struct {
	// How Navaction and nested children are related
	Group ActionGoup `json:"group,omitempty"`

	//  The name of the action
	Title string `json:"title,omitempty"`

	// How the action is handled when tapped
	OnTapRoute string `json:"onTapRoute,omitempty"`

	// A link to a PNG image that would serve as an avatar
	Icon string `json:"icon,omitempty"`

	// Whether the user has marked the action as a favourite
	Favourite bool `json:"favourite,omitempty"`

	IsParent bool `json:"isParent,omitempty"`

	IsHighPriority bool

	// Sub menus in a navigation action
	Nested []interface{} `json:"nested,omitempty"`
}

//NavigationActions are Role based Navigation Actions for a User
type NavigationActions struct {
	// The primary actions the user can perform
	Primary []NavigationAction `json:"primary"`

	// The secondary action the user can perform
	Secondary []NavigationAction `json:"secondary"`
}

// Upload represents a file uploaded to cloud storage
type Upload struct {
	ID          string    `json:"id"          firestore:"id"`
	URL         string    `json:"url"         firestore:"url"`
	Size        int       `json:"size"        firestore:"size"`
	Hash        string    `json:"hash"        firestore:"hash"`
	Creation    time.Time `json:"creation"    firestore:"creation"`
	Title       string    `json:"title"       firestore:"title"`
	ContentType string    `json:"contentType" firestore:"contentType"`
	Language    string    `json:"language"    firestore:"language"`
	Base64data  string    `json:"base64data"  firestore:"base64data"`
}

// IsEntity marks upload as an apollo federation entity
func (u Upload) IsEntity() {}

// IsNode marks upload as a relay node
func (u Upload) IsNode() {}

// SetID marks upload as a relay node
func (u Upload) SetID(id string) {
	u.ID = id
}

// GetID marks upload as a relay node
func (u Upload) GetID() firebasetools.ID {
	return firebasetools.IDValue(u.ID)
}

// UploadInput is used to send data for new uploads
type UploadInput struct {
	Title       string `json:"title"`
	ContentType string `json:"contentType"`
	Language    string `json:"language"`
	Base64data  string `json:"base64data"`
	Filename    string `json:"filename"`
}
