package cufeedback

import "time"

// Feedback struct represents a users feedback.
type Feedback struct {
	ID          int       `db:"feedback_id" json:"feedback_id"`
	UserName    string    `db:"user_name" json:"user_name"`
	Type        string    `db:"type" json:"type"`
	Location    string    `db:"location" json:"location"`
	Description string    `db:"description" json:"description"`
	Email       string    `db:"fix_email" json:"fix_email"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Building    `db:"building" json:"building"`
}

// Building represents a building on Clemson's campus.
type Building struct {
	ID   int    `db:"building_id" json:"building_id"`
	Name string `db:"name" json:"name"`
}
