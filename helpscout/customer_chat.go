package helpscout

const (
	CustomerChatTypeAim   = "aim"
	CustomerChatTypeGtalk = "gtalk"
	CustomerChatTypeIcq   = "icq"
	CustomerChatTypeXmpp  = "xmpp"
	CustomerChatTypeMsn   = "msn"
	CustomerChatTypeSkype = "skype"
	CustomerChatTypeYahoo = "yahoo"
	CustomerChatTypeQq    = "qq"
	CustomerChatTypeOther = "other"
)

type CustomerChat struct {
	Id     int
	Value  string
	TypeOf string
}
