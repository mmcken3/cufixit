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
		params = append(params, fb.UserName, fb.Type, bID, fb.Location,
			fb.Description, fb.Email, fb.PhoneNumber)
		query := `
		INSERT INTO feedback (user_name, type, building_id, location,
		description, fix_email, phone_number) VALUES ` + buildValues(7)
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

// GetAllFeedback gets all of the feedback from the table and returns it as a slice.
func (db *DB) GetAllFeedback() ([]cufixit.Feedback, error) {
	var feedback []cufixit.Feedback
	err := db.Transact(func(tx *sqlx.Tx) error {
		err := tx.Select(&feedback, `
			SELECT 
				feedback_id, 
				user_name, 
				type, 
				name "building.name", 
				location, 
				description, 
				fix_email, 
				phone_number,
				updated_at,
				b.building_id "building.building_id"
			FROM feedback f INNER JOIN 
			building b ON 
			f.building_id = b.building_id`)
		return errors.Wrapf(err, "Error getting ID from buildings table.")
	})
	return feedback, err
}
