package helpscout

const (
	SourceTypeEmail = "email"
	SourceTypeWeb = "web"
	SourceTypeNotification = "notification"
	SourceTypeEmailfwd = "emailfwd"
	SourceTypeAPI = "api"
	SourceTypeChat = "chat"

	SourceViaCustomer = "customer"
	SourceViaUser = "user"
)

type Source struct {
	TypeOf string
	Via string
}
