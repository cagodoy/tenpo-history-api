package service

import (
	"fmt"

	history "github.com/cagodoy/tenpo-history-api"
	"github.com/cagodoy/tenpo-history-api/database"
	nats "github.com/nats-io/nats.go"
)

// NewHistory ...
func NewHistory(store database.Store, conn *nats.EncodedConn) *History {
	// listen to events
	go func() {
		conn.Subscribe("history.create", func(he *history.CreateHistoryEvent) {
			fmt.Printf("Received a History event: %+v\n", he)

			h := &history.History{
				UserID:    he.UserID,
				Latitude:  he.Latitude,
				Longitude: he.Longitude,
			}
			err := store.HistoryCreate(h)
			if err != nil {
				fmt.Printf("Error in conn.Subscribe(history.create), %v", err)
				return
			}
			fmt.Println("[NATS][HistoryService][History.Create][Subscribe] Created ok")
		})
	}()

	return &History{
		Store: store,
		Nats:  conn,
	}
}

// History ...
type History struct {
	Store database.Store
	Nats  *nats.EncodedConn
}

// ListByUserID ...
func (hs *History) ListByUserID(userID string) ([]*history.History, error) {
	q := &history.Query{
		UserID: userID,
	}

	return hs.Store.HistoryListByUserID(q)
}

// Create ...
func (hs *History) Create(h *history.History) error {
	return hs.Store.HistoryCreate(h)
}

// List ...
func (hs *History) List() ([]*history.History, error) {
	return hs.Store.HistoryList()
}
