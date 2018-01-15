package helpscout

import "time"

const (
	CustomerPhotoTypeUnknown       = "unknown"
	CustomerPhotoTypeGravatar      = "gravatar"
	CustomerPhotoTypeTwitter       = "twitter"
	CustomerPhotoTypeFacebook      = "facebook"
	CustomerPhotoTypeGoogleprofile = "googleprofile"
	CustomerPhotoTypeGoogleplus    = "googleplus"
	CustomerPhotoTypeLinkedin      = "linkedin"

	CustomerGenderMale    = "male"
	CustomerGenderFemale  = "female"
	CustomerGenderUnknown = "unknown"
)

type Customer struct {
	Id           int
	FirstName    string
	LastName     string
	PhotoUrl     string
	PhotoType    string
	Gender       string
	Age          string
	Organization string
	JobTitle     string
	Location     string
	CreatedAt    time.Time
	ModifiedAt   time.Time

	Background     string
	Address        CustomerAddress
	SocialProfiles []CustomerSocialProfile
	Emails         []CustomerEmail
	Phones         []CustomerPhone
	Chats          []CustomerChat
	Websites       []CustomerWebsite
}
