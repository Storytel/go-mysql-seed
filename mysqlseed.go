package mysqlseed

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// ApplySeedWithCmd loads a seed sql file and executes it against the db.
// expects hostnameAndPort to be on the form `127.0.0.1:8080`
// Requires MySQL Command-Line Tool to be installed
func ApplySeedWithCmd(hostnameAndPort string, dbUser string, dbName string, seedFilePath string) error {
	instanceHostAndPort := strings.Split(hostnameAndPort, ":")
	hostName := instanceHostAndPort[0]
	if hostName == "localhost" {
		hostName = "127.0.0.1"
	}
	hostPort := instanceHostAndPort[1]

	cmd := exec.Command("mysql", fmt.Sprintf("-h%s", hostName), fmt.Sprintf("-u%s", dbUser), fmt.Sprintf("-P%s", hostPort), dbName, "-e", fmt.Sprintf("source %s", seedFilePath))

	var out, stderr bytes.Buffer

	cmd.Stdout = &bytes.Buffer{}
	cmd.Stderr = &bytes.Buffer{}

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error executing query. Command Output: %+v\n: %+v, %v", out.String(), stderr.String(), err)
	}

	return nil
}

// ApplySeedWithDB loads a seed sql file and executes it against the db.
// Requires MySQL connection to use `multiStatements=true`
func ApplySeedWithDB(db *sql.DB, seedFilePath string) error {
	fileBytes, err := ioutil.ReadFile(seedFilePath)
	if err != nil {
		return fmt.Errorf("Could not read seed-file, ", err)
	}

	_, err = db.Exec(string(fileBytes))
	if err != nil {
		return fmt.Errorf("Could not apply seed, ", err)
	}

	return nil
}
