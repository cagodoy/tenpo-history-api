package history

import (
	"time"

	pb "github.com/cagodoy/tenpo-challenge/lib/proto"
)

// History ...
type History struct {
	ID string `json:"id" db:"id"`

	UserID    string `json:"user_id" db:"user_id"`
	Latitude  string `json:"latitude" db:"latitude"`
	Longitude string `json:"longitude" db:"longitude"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

// Service ...
type Service interface {
	ListByUserID(userID string) ([]*History, error)
	Create(*History) error
	List() ([]*History, error)
}

// Query ...
type Query struct {
	UserID string
}

// ToProto ...
func (h *History) ToProto() *pb.History {
	return &pb.History{
		Id:        h.ID,
		UserId:    h.UserID,
		Latitude:  h.Latitude,
		Longitude: h.Longitude,
		CreatedAt: h.CreatedAt.UnixNano(),
		UpdatedAt: h.UpdatedAt.UnixNano(),
	}
}

// FromProto ...
func (h *History) FromProto(hh *pb.History) *History {
	h.ID = hh.Id
	h.UserID = hh.UserId
	h.Latitude = hh.Latitude
	h.Longitude = hh.Longitude
	h.CreatedAt = time.Unix(hh.CreatedAt, 0)
	h.UpdatedAt = time.Unix(hh.UpdatedAt, 0)

	return h
}

// CreateHistoryEvent ...
type CreateHistoryEvent struct {
	UserID    string
	Latitude  string
	Longitude string
}
