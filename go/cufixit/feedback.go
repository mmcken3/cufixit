package cufixit

import "time"

// Feedback struct represents a users feedback.
type Feedback struct {
	ID          int       `db:"feedback_id" json:"feedback_id"`
	UserName    string    `db:"user_name" json:"user_name"`
	Description string    `db:"description" json:"description"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	ImageURL    string    `db:"image_url" json:"image_url"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Building    `db:"building" json:"building"`
	Type        `db:"type" json:"type"`
}

// Building represents a building on Clemson's campus.
type Building struct {
	ID   int    `db:"building_id" json:"building_id"`
	Name string `db:"name" json:"name"`
}

// Type represents an issue someone using the app would have.
type Type struct {
	ID      int    `db:"type_id" json:"type_id"`
	Type    string `db:"type" json:"type"`
	Contact string `db:"contact" json:"contact"`
}

// Location represents a maps locations.
type Location struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}
