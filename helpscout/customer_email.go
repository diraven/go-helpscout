package helpscout

const (
	CustomerEmailLocationHome  = "home"
	CustomerEmailLocationWork  = "work"
	CustomerEmailLocationOther = "other"
)

type CustomerEmail struct {
	Id       int
	Value    string
	Location string
}
