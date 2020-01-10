package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jamf/go-mysqldump"
)

// setupDatabase try to open a connection to given dbInfo then returns the instance of connection or nil
func setupDatabase(dbInfo Config) *sql.DB {
	// nota! sql.Open apenas valida os argumentos, não abre uma conexão.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbInfo.Database.User, dbInfo.Database.Password, dbInfo.Database.Hostname, dbInfo.Database.Port, dbInfo.Database.Db))
	if err != nil {
		fmt.Printf("Error on setupDatabase() %s", err)
		return nil
	} else if db.Ping() != nil {
		fmt.Print("Error on setupDatabase() = ping failed")
		return nil
	}

	return db
}

// Validate if dir Exists and If not, create the dir. Returns the string with dir created
func createDirIfNotExists(dirPath string) string {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, os.ModePerm)
		return dirPath
	}
	return dirPath
}

// dumpDatabase receives a databaseName to dump
func dumpDatabase(config Config) (string, error) {
	dumpDir := createDirIfNotExists("dumps")
	dumpFileFormat := fmt.Sprintf("%s-2006-01-02 15_04", config.Database.Db)
	dumpFileName := time.Now().Format(dumpFileFormat) + ".sql"
	if db := setupDatabase(config); db != nil {
		dumper, err := mysqldump.Register(db, dumpDir, dumpFileFormat)
		if err != nil {
			fmt.Printf("Error %s", err)
			return "", err
		}

		dumper.Dump()
		dumper.Close()
		return dumpFileName, nil
	}
	return "", fmt.Errorf("Error while trying to connect to db")
}
