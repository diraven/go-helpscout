package main

import (
	"log"
	"github.com/diraven/go-helpscout/helpscout"
)

func main() {
	hs := helpscout.New("yourApiKey")

	convos, currPage, pages, count, err := hs.Conversations.ListMailboxFolderConversations(123, 321, 0, "", nil, "")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(convos)
	log.Println(currPage)
	log.Println(pages)
	log.Println(count)
	log.Println(err)
}
