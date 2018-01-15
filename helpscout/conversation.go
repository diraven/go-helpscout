package helpscout

import (
	"time"
	"net/http"
	"strconv"
	"fmt"
)

const (
	ConversationStatusActive  = "active"
	ConversationStatusPending = "pending"
	ConversationStatusClosed  = "closed"
	ConversationStatusSpam    = "spam"
)

type Conversation struct {
	Id             int
	TypeOf         string
	FolderId       int
	IsDraft        bool
	Number         int
	Owner          Person
	Mailbox        MailboxRef
	Customer       Person
	ThreadCount    int
	Status         string
	Subject        string
	Preview        string
	CreatedBy      Person
	CreatedAt      time.Time
	ModifiedAt     time.Time
	UserModifiedAt time.Time
	ClosedAt       time.Time
	ClosedBy       Person
	Source         Source
	Cc             []string
	Bcc            []string
	Tags           []string
	Threads        []Thread
	CustomFields   []CustomField
}

type ConversationsResponse struct {
	Page  int
	Pages int
	Count int
	Items []Conversation
}
type ConversationsService service

func (s *ConversationsService) listConversations(path string, page int, status string, modifiedSince *time.Time, tag string) (*ConversationsResponse, error) {
	// Prepare variable to contain result.
	var result ConversationsResponse
	var params = map[string]string{
		"page":   strconv.Itoa(page),
		"status": status,
		"tag":    tag,
	}

	if modifiedSince != nil {
		params["modifiedSince"] = modifiedSince.Format(time.RFC3339)
	}

	// Generate and perform request.
	err := s.client.DoNewRequest(http.MethodGet, path, &result, params)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ListMailboxConversations https://developer.helpscout.com/help-desk-api/conversations/list/ (1)
func (s *ConversationsService) ListMailboxConversations(mailboxId, page int, status string, modifiedSince *time.Time, tag string) (conversations []Conversation, currPage, Pages, Count int, err error) {
	// Prepare path.
	var path = fmt.Sprintf("mailboxes/%d/conversations.json", mailboxId)

	// Perform request.
	result, err := s.listConversations(path, page, status, modifiedSince, tag)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	// Return the response.
	return result.Items, result.Page, result.Pages, result.Count, nil
}

// ListMailboxConversations https://developer.helpscout.com/help-desk-api/conversations/list/ (2)
func (s *ConversationsService) ListMailboxFolderConversations(mailboxId, folderId, page int, status string, modifiedSince *time.Time, tag string) (conversations []Conversation, currPage, Pages, Count int, err error) {
	// Prepare path.
	var path = fmt.Sprintf("mailboxes/%d/folders/%d/conversations.json", mailboxId, folderId)

	// Perform request.
	result, err := s.listConversations(path, page, status, modifiedSince, tag)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	// Return the response.
	return result.Items, result.Page, result.Pages, result.Count, nil
}

// ListMailboxConversations https://developer.helpscout.com/help-desk-api/conversations/list/ (3)
func (s *ConversationsService) ListMailboxCustomerConversations(mailboxId, customerId, page int, status string, modifiedSince *time.Time, tag string) (conversations []Conversation, currPage, Pages, Count int, err error) {
	// Prepare path.
	var path = fmt.Sprintf("mailboxes/%d/customers/%d/conversations.json", mailboxId, customerId)

	// Perform request.
	result, err := s.listConversations(path, page, status, modifiedSince, tag)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	// Return the response.
	return result.Items, result.Page, result.Pages, result.Count, nil
}

// ListMailboxConversations https://developer.helpscout.com/help-desk-api/conversations/list/ (4)
func (s *ConversationsService) ListMailboxUserConversations(mailboxId, userId, page int, status string, modifiedSince *time.Time, tag string) (conversations []Conversation, currPage, Pages, Count int, err error) {
	// Prepare path.
	var path = fmt.Sprintf("mailboxes/%d/users/%d/conversations.json", mailboxId, userId)

	// Perform request.
	result, err := s.listConversations(path, page, status, modifiedSince, tag)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	// Return the response.
	return result.Items, result.Page, result.Pages, result.Count, nil
}
