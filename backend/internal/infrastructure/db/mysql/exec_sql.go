package mysql

import (
	"database/sql"
	"fmt"
	"io"
	"os"
)

func ExecSQLFile(db *sql.DB, filePath string) error {
	sqlFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening SQL file: %w", err)
	}
	defer sqlFile.Close()

	sqlBytes, err := io.ReadAll(sqlFile)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %w", err)
	}

	sqlQueries := string(sqlBytes)

	_, err = db.Exec(sqlQueries)
	if err != nil {
		return fmt.Errorf("error executing SQL queries: %w", err)
	}

	return nil
}
