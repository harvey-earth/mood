# #!/usr/bin/env bash

# Check if an argument is passed
if [ $# -eq 0 ]; then
    echo "Usage: $0 [mysql|sqlite|psql]"
    exit 1
fi

# Set variables
DB_TYPE=$1

MYSQL_SCRIPT="mood-mysql.sql"
SQLITE_SCRIPT="mood-sqlite.sql"
POSTGRES_SCRIPT="mood-postgres.sql"
DATABASE_USER=${DATABASE_USER}
DATABASE_PASSWORD=${DATABASE_PASSWORD}
DATABASE_HOST=${DATABASE_HOST}
DATABASE_NAME=${DATABASE_NAME}

# Run scripts
case "$DB_TYPE" in
    mysql)
        if [ -f "$MYSQL_SCRIPT" ]; then
            echo "Running MySQL script..."
            mysql -h "$DATABASE_HOST" -u "$DATABASE_USER" -p"$DATABASE_PASSWORD" $DATABASE_NAME < "$MYSQL_SCRIPT"
        else
            echo "MySQL script not found."
        fi
        ;;
    sqlite)
        if [ -f "$SQLITE_SCRIPT" ]; then
            echo "Running SQLite script..."
            sqlite3 mood.db < "$SQLITE_SCRIPT"
        else
            echo "SQLite script not found."
        fi
        ;;
    psql)
        if [ -f "$POSTGRES_SCRIPT" ]; then
            echo "Running PostgreSQL script..."
            PGPASSWORD="$DATABASE_PASSWORD" psql -h "$DATABASE_HOST" -U "$DATABASE_USER" -d "$DATABASE_NAME" -f "$POSTGRES_SCRIPT"
        else
            echo "PostgreSQL script not found."
        fi
        ;;
    *)
        echo "Invalid option. Use 'mysql', 'sqlite', or 'psql'."
        exit 1
        ;;
esac
