package helpscout

const (
	PersonTypeUser     = "user"
	PersonTypeCustomer = "customer"
	PersonTypeTeam     = "team"
)

type Person struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	TypeOf    string
	PhotoUrl  string
}
