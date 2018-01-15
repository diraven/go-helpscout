package helpscout

import "time"

const (
	ThreadTypeLineitem      = "lineitem"
	ThreadTypeNote          = "note"
	ThreadTypeMessage       = "message"
	ThreadTypeChat          = "chat"
	ThreadTypeCustomer      = "customer"
	ThreadTypeForwardparent = "forwardparent"
	ThreadTypeForwardchild  = "forwardchild"
	ThreadTypePhone         = "phone"

	ThreadStatePublished   = "published"
	ThreadStateDraft       = "draft"
	ThreadStateUnderreview = "underreview"
	ThreadStateHidden      = "hidden"

	ThreadStatusNochange = "nochange"
	ThreadStatusActive   = "active"
	ThreadStatusPending  = "pending"
	ThreadStatusClosed   = "closed"
	ThreadStatusSpam     = "spam"

	ThreadActionTypeMovedFromMailbox      = "MovedFromMailbox"
	ThreadActionTypeMerged                = "merged"
	ThreadActionTypeImported              = "imported"
	ThreadActionTypeWorkflow              = "workflow"
	ThreadActionTypeImportedExternal      = "importedExternal"
	ThreadActionTypeChangedTickedCustomer = "changedTicketCustomer"
	ThreadActionTypeDeletedTicket         = "deletedTicket"
	ThreadActionTypeRestoreTicket         = "restoreTicket"
)

type Thread struct {
	Id                int
	AssignedTo        Person
	Status            string
	CreatedAt         time.Time
	OpenedAt          time.Time
	CreatedBy         Person
	Source            Source
	ActionType        string
	ActionSourceId    int
	FromMailbox       MailboxRef
	TypeOf            string
	State             string
	Customer          Person
	Body              string
	To                []string
	Cc                []string
	Bcc               []string
	Attachments       []Attachment
	SavedReplyId      int
	CreatedByCustomer bool
}
