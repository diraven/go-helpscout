package helpscout

import (
	"net/http"
	"time"
	"strconv"
	"fmt"
)

type MailboxRef struct {
	Id   int
	Name string
}

const (
	FolderTypeOpen           = "open"
	FolderTypeMine           = "mine"
	FolderTypeNeedsattention = "needsattention"
	FolderTypeDrafts         = "drafts"
	FolderTypeAssigned       = "assigned"
	FolderTypeClosed         = "closed"
	FolderTypeSpam           = "spam"
)

type Folder struct {
	Id          int
	Name        string
	TypeOf      string
	UserId      int
	TotalCount  int
	ActiveCount int
	ModifiedAt  time.Time
}

type Mailbox struct {
	MailboxRef
	Slug         string
	Email        string
	CreatedAt    time.Time
	ModifiedAt   time.Time
	CustomFields []CustomField
	Folders      []Folder
}

type MailboxResponse struct {
	Item Mailbox
}

type MailboxesResponse struct {
	Page  int
	Pages int
	Count int
	Items []Mailbox
}

type FoldersResponse struct {
	Page  int
	Pages int
	Count int
	Items []Folder
}

type MailboxesService service

// ListMailboxes https://developer.helpscout.com/help-desk-api/mailboxes/list/
func (s *MailboxesService) ListMailboxes(page int) (mailboxes []Mailbox, currPage, Pages, Count int, err error) {
	// Prepare variable to contain result.
	var result MailboxesResponse
	var path = "mailboxes.json"
	var params = map[string]string{
		"page": strconv.Itoa(page),
	}

	// Generate and perform request.
	err = s.client.DoNewRequest(http.MethodGet, path, &result, params)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	// Return the response.
	return result.Items, result.Page, result.Pages, result.Count, nil
}

// GetMailboxes https://developer.helpscout.com/help-desk-api/mailboxes/get/
func (s *MailboxesService) GetMailboxes(id int) (*Mailbox, error) {
	// Prepare variable to contain result.
	var result MailboxResponse
	var path = fmt.Sprintf("mailboxes/%d.json", id)
	var params = map[string]string{}

	// Generate and perform request.
	err := s.client.DoNewRequest(http.MethodGet, path, &result, params)
	if err != nil {
		return nil, err
	}

	// Return the response.
	return &result.Item, nil
}

// GetFolders https://developer.helpscout.com/help-desk-api/mailboxes/folders/
func (s *MailboxesService) GetFolders(id int) ([]Folder, error) {
	// Prepare variable to contain result.
	var result FoldersResponse
	var path = fmt.Sprintf("mailboxes/%d/folders.json", id)
	var params = map[string]string{}

	// Generate and perform request.
	err := s.client.DoNewRequest(http.MethodGet, path, &result, params)
	if err != nil {
		return nil, err
	}

	// Return the response.
	return result.Items, nil
}
