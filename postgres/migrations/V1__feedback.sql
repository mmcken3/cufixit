CREATE TABLE building (
    building_id serial PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE feedback (
    feedback_id serial PRIMARY KEY,
    user_name TEXT NOT NULL,
    type TEXT NOT NULL,
    building INT REFERENCES building(building_id) NOT NULL,
    location TEXT,
    description TEXT NOT NULL,
    fix_email TEXT NOT NULL,
    updated_at timestamptz
);

-- autoupdate is a generic trigger to update the updated_at column to current time.
CREATE FUNCTION autoupdate() RETURNS trigger AS $$
BEGIN
    new.updated_at = now();
    RETURN new;
END;
$$ LANGUAGE plpgsql;

-- feedback_trigger updates the feedback table's updated_at column on update.
CREATE TRIGGER feedback_trigger
    BEFORE INSERT OR UPDATE
    ON feedback FOR EACH ROW EXECUTE PROCEDURE autoupdate();