package types

import (
	"time"
)

type UserStore interface {
	CreateUser(User) error
	GetUserByEmail(email string) (*User, error)
}

type RegisterUserPayload struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"-"`

	CreatedAt string `json:"created_at"`
}

type Event struct {
	ID        int       `json:"id"`
	Sport     string    `json:"sport"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	CreatedAt  string `json:"created_at"`
	UpadatedAt string `json:"updated_at"`
}

type Bet struct {
	ID        int         `json:"id"`
	UserId    int         `json:"user_id"`
	EventId   int         `json:"event_id`
	Status    EventStatus `json:"status"`
	Selection string      `json:"selection"`
	Result    string      `json:"result"`

	CreatedAt  string `json:"created_at"`
	UpadatedAt string `json:"updated_at"`
}

type EventStatus string

const (
	Upcoming  EventStatus = "upcoming"
	Ongoing   EventStatus = "ongoing"
	Completed EventStatus = "completed"
)
