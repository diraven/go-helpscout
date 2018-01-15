package helpscout

const (
	CustomerPhoneLocationHome   = "home"
	CustomerPhoneLocationWork   = "work"
	CustomerPhoneLocationMobile = "mobile"
	CustomerPhoneLocationFax    = "fax"
	CustomerPhoneLocationPager  = "pager"
	CustomerPhoneLocationOther  = "other"
)

type CustomerPhone struct {
	Id       int
	Value    string
	Location string
}
