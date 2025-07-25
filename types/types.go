package types

import (
	"time"
)

type UserStore interface {
	CreateUser(User) error
	GetUserByEmail(string) (*User, error)
	GetUserById(int) (*User, error)
}

type RegisterUserPayload struct {
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`

	CreatedAt time.Time `json:"created_at"`
}

type Event struct {
	ID        int       `json:"id"`
	Sport     string    `json:"sport"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	CreatedAt  time.Time `json:"created_at"`
	UpadatedAt time.Time `json:"updated_at"`
}

type Bet struct {
	ID        int         `json:"id"`
	UserId    int         `json:"user_id"`
	EventId   int         `json:"event_id"`
	Status    EventStatus `json:"status"`
	Selection string      `json:"selection"`
	Result    string      `json:"result"`

	CreatedAt  time.Time `json:"created_at"`
	UpadatedAt time.Time `json:"updated_at"`
}

type EventStatus string

const (
	Upcoming  EventStatus = "upcoming"
	Ongoing   EventStatus = "ongoing"
	Completed EventStatus = "completed"
)
