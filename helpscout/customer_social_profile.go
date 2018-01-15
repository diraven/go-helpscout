package helpscout

const (
	CustomerSocialProfileTypeTwitter    = "twitter"
	CustomerSocialProfileTypeFacebook   = "facebook"
	CustomerSocialProfileTypeLinkedin   = "linkedin"
	CustomerSocialProfileTypeAboutme    = "aboutme"
	CustomerSocialProfileTypeGoogle     = "google"
	CustomerSocialProfileTypeGoogleplus = "googleplus"
	CustomerSocialProfileTypeTungleme   = "tungleme"
	CustomerSocialProfileTypeQuora      = "quora"
	CustomerSocialProfileTypeFoursquare = "foursquare"
	CustomerSocialProfileTypeYoutube    = "youtube"
	CustomerSocialProfileTypeFlickr     = "flickr"
	CustomerSocialProfileTypeOther      = "other"
)

type CustomerSocialProfile struct {
	Id     int
	Value  string
	TypeOf string
}
