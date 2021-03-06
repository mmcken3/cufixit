package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/mmcken3/cufixit/go/cufixit"
	"github.com/pkg/errors"
)

// CreateFeedback saves the feedback into the DB.
func (db *DB) CreateFeedback(fb cufixit.Feedback) error {
	err := db.Transact(func(tx *sqlx.Tx) error {
		var params []interface{}
		bID, err := db.GetBuildingID(fb.Building, tx)
		if err != nil {
			return err
		}
		tID, err := db.GetTypeID(fb.Type, tx)
		if err != nil {
			return err
		}
		params = append(params, fb.UserName, tID, bID,
			fb.Description, fb.PhoneNumber, fb.ImageURL)
		query := `
		INSERT INTO feedback (user_name, type_id, building_id,
		description, phone_number, image_url) VALUES ` + buildValues(6)
		_, err = tx.Exec(query, params...)
		return errors.Wrapf(err, "Error inserting the feedback into the database.")
	})
	return err
}

// GetBuildingID gets the building ID from the entered name.
func (db *DB) GetBuildingID(b cufixit.Building, tx *sqlx.Tx) (int, error) {
	var bID []int
	err := tx.Select(&bID, `
		SELECT DISTINCT ON (building_id) building_id FROM building WHERE name = '`+b.Name+`'`)
	return bID[0], errors.Wrapf(err, "Error getting ID from buildings table.")
}

// GetTypeID gets the building ID from the entered name.
func (db *DB) GetTypeID(t cufixit.Type, tx *sqlx.Tx) (int, error) {
	var tID []int
	err := tx.Select(&tID, `
		SELECT DISTINCT ON (type_id) type_id FROM type WHERE type = '`+t.Type+`'`)
	return tID[0], errors.Wrapf(err, "Error getting ID from type table.")
}

// GetTypeContact gets the building ID from the entered name.
func (db *DB) GetTypeContact(t cufixit.Type) (string, error) {
	var tC []string
	err := db.Transact(func(tx *sqlx.Tx) error {
		err := tx.Select(&tC, `
		SELECT DISTINCT ON (type_id) contact FROM type WHERE type = '`+t.Type+`'`)
		return err
	})
	return tC[0], errors.Wrapf(err, "Error getting Contact from type table.")
}

// GetAllFeedback gets all of the feedback from the table and returns it as a slice.
func (db *DB) GetAllFeedback() ([]cufixit.Feedback, error) {
	var feedback []cufixit.Feedback
	err := db.Transact(func(tx *sqlx.Tx) error {
		err := tx.Select(&feedback, `
			SELECT 
				feedback_id, 
				user_name, 
				t.type_id "type.type_id",
				type "type.type", 
				contact "type.contact",
				name "building.name",
				description, 
				phone_number,
				image_url,
				updated_at,
				b.building_id "building.building_id"
			FROM feedback f INNER JOIN 
			building b ON 
			f.building_id = b.building_id
			INNER JOIN type t ON
			f.type_id = t.type_id`)
		return errors.Wrapf(err, "Error getting ID from buildings table.")
	})
	return feedback, err
}

// GetFeedbackofType gets all of the feedback from the table of passed type and returns it as a slice.
func (db *DB) GetFeedbackofType(t string) ([]cufixit.Feedback, error) {
	var feedback []cufixit.Feedback
	err := db.Transact(func(tx *sqlx.Tx) error {
		err := tx.Select(&feedback, `
			SELECT 
				feedback_id, 
				user_name, 
				t.type_id "type.type_id",
				type "type.type", 
				contact "type.contact",
				name "building.name",
				description, 
				phone_number,
				image_url,
				updated_at,
				b.building_id "building.building_id"
			FROM feedback f INNER JOIN 
			building b ON 
			f.building_id = b.building_id
			INNER JOIN type t ON
			f.type_id = t.type_id WHERE t.type = '`+t+`'`)
		return errors.Wrapf(err, "Error getting ID from buildings table.")
	})
	return feedback, err
}

// GetFeedbackofBuilding gets all of the feedback from the table of passed building and returns it as a slice.
func (db *DB) GetFeedbackofBuilding(b string) ([]cufixit.Feedback, error) {
	var feedback []cufixit.Feedback
	err := db.Transact(func(tx *sqlx.Tx) error {
		err := tx.Select(&feedback, `
			SELECT 
				feedback_id, 
				user_name, 
				t.type_id "type.type_id",
				type "type.type", 
				contact "type.contact",
				name "building.name",
				description, 
				phone_number,
				image_url,
				updated_at,
				b.building_id "building.building_id"
			FROM feedback f INNER JOIN 
			building b ON 
			f.building_id = b.building_id
			INNER JOIN type t ON
			f.type_id = t.type_id WHERE b.name = '`+b+`'`)
		return errors.Wrapf(err, "Error getting ID from buildings table.")
	})
	return feedback, err
}

// GetFeedbackofUser gets all of the feedback from the table of passed building and returns it as a slice.
func (db *DB) GetFeedbackofUser(u string) ([]cufixit.Feedback, error) {
	var feedback []cufixit.Feedback
	err := db.Transact(func(tx *sqlx.Tx) error {
		err := tx.Select(&feedback, `
			SELECT 
				feedback_id, 
				user_name, 
				t.type_id "type.type_id",
				type "type.type", 
				contact "type.contact",
				name "building.name",
				description, 
				phone_number,
				image_url,
				updated_at,
				b.building_id "building.building_id"
			FROM feedback f INNER JOIN 
			building b ON 
			f.building_id = b.building_id
			INNER JOIN type t ON
			f.type_id = t.type_id WHERE f.user_name = '`+u+`'`)
		return errors.Wrapf(err, "Error getting ID from buildings table.")
	})
	return feedback, err
}
