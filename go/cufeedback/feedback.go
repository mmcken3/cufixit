package cufeedback

import "time"

// Feedback struct represents a users feedback.
type Feedback struct {
	ID          int       `db:"feedback_id"`
	UserName    string    `db:"user_name"`
	Type        string    `db:"type"`
	Building    string    `db:"building"`
	Location    string    `db:"location"`
	Description string    `db:"description"`
	Email       string    `db:"fix_email"`
	UpdatedAt   time.Time `db:"updated_at"`
}
