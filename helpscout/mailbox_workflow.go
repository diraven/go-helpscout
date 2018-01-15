package helpscout

import "time"

const (
	MailboxWorkflowTypeAutomatic = "automatic"
	MailboxWorkflowTypeManual    = "manual"

	MailboxWorkflowStatusActive   = "active"
	MailboxWorkflowStatusInactive = "inactive"
	MailboxWorkflowStatusInvalid  = "invalid"
)

type MailboxWorkflow struct {
	Id         int
	MailboxId  int
	TypeOf     string
	Status     string
	Order      int
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}
