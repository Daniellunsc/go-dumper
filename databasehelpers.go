package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jamf/go-mysqldump"
)

// DatabaseConfig represent Database Configuration variables
type DatabaseConfig struct {
	username     string
	password     string
	hostname     string
	port         string
	databaseName string
}

// setupDatabaseInfo creates a new databaseConfig
func setupDatabaseInfo(_username string, _password string, _hostname string, _port string, _database string) DatabaseConfig {

	dbConfig := DatabaseConfig{
		username:     _username,
		password:     _password,
		hostname:     _hostname,
		port:         _port,
		databaseName: _database,
	}
	return dbConfig
}

// setupDatabase try to open a connection to given dbInfo then returns the instance of connection or nil
func setupDatabase(dbInfo DatabaseConfig) *sql.DB {
	// nota! sql.Open apenas valida os argumentos, não abre uma conexão.
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbInfo.username, dbInfo.password, dbInfo.hostname, dbInfo.port, dbInfo.databaseName))
	if err != nil {
		fmt.Printf("Error on setupDatabase() %s", err)
		return nil
	} else if db.Ping() != nil {
		fmt.Print("Error on setupDatabase() = ping failed")
		return nil
	}

	return db
}

// dumpDatabase receives a databaseName to dump
func dumpDatabase(databaseName string) (string, error) {
	// Configure aqui suas variáveis de ambiente
	databaseInfo := setupDatabaseInfo("root", "12345678", "localhost", "3306", databaseName)
	dumpDir := "dumps"
	dumpFileFormat := fmt.Sprintf("%s-2006-01-02 15_04", databaseInfo.databaseName)
	dumpFileName := time.Now().Format(dumpFileFormat) + ".sql"
	if db := setupDatabase(databaseInfo); db != nil {
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
