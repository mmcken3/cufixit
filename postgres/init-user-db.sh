#!/bin/bash
set -e

echo "Creating database $DB_NAME"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE USER $DB_USER;
    CREATE DATABASE $DB_NAME;
    GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;
EOSQL

echo "Loading extensions into $DB_NAME"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="$DB_NAME" <<-EOSQL
    CREATE EXTENSION fuzzystrmatch;
EOSQL

psql -d $DB_NAME -U $DB_USER < ./migrations/V1__feedback.sql
psql -d $DB_NAME -U $DB_USER < ./migrations/V2__buildings.sql