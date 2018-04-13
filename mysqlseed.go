package mysqlseed

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ApplySeed loads a seed sql file and executes it against the db.
// expect hostnameAndPort to be on the form `127.0.0.1:8080`
func ApplySeed(hostnameAndPort string, dbUser string, dbName string, seedFilePath string) error {
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
