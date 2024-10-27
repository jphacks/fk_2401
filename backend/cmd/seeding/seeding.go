package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/db/mysql"
)

var sqlFilePath = "db/seeds"

func main() {
	db, err := mysql.ConnectDB()
	if err != nil {
		log.Printf("error connecting to the database: %v", err)
		return
	}
	defer db.Close()

	files, err := os.ReadDir(sqlFilePath)
	if err != nil {
		log.Printf("error reading sql file path: %v", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			filePath := filepath.Join(sqlFilePath, file.Name())
			err = mysql.ExecSQLFile(db, filePath)
			if err != nil {
				log.Printf("error executing sql file %s: %v", file.Name(), err)
			} else {
				fmt.Printf("successfully executed sql file %s\n", file.Name())
			}
		}
	}

}
