package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/mmcken3/cufeedback/go/cufeedback"
	"github.com/pkg/errors"
)

// CreateFeedback saves the feedback into the DB.
func (db *DB) CreateFeedback(fb cufeedback.Feedback) error {
	err := db.Transact(func(tx *sqlx.Tx) error {
		var params []interface{}
		bID, err := db.GetBuildingID(fb.Building, tx)
		if err != nil {
			return err
		}
		params = append(params, fb.UserName, fb.Type, bID, fb.Location,
			fb.Description, fb.Email)
		query := `
		INSERT INTO feedback (user_name, type, building, location,
		description, fix_email) VALUES ` + buildValues(6)
		_, err = tx.Exec(query, params...)
		return errors.Wrapf(err, "Error inserting the feedback into the database.")
	})
	return err
}

// GetBuildingID gets the building ID from the entered name.
func (db *DB) GetBuildingID(b cufeedback.Building, tx *sqlx.Tx) (int, error) {
	var bID []int
	err := tx.Select(&bID, `
		SELECT DISTINCT ON (building_id) building_id FROM building WHERE name = '`+b.Name+`'`)
	return bID[0], errors.Wrapf(err, "Error getting ID from buildings table.")
}
