# go-mysql-seed

Utility for seeding a MySQL database with Go.
Useful during testing.

This package supports 2 types of MySQL seeding, both requires a path to a `.sql` file.

1. Seeding via `MySQL Command Line Tool`:
   Initialize a import using your locally installed MySQL CLT.

2. Seeding via `*sql.DB` struct: Start an import of the specified file using a `*sql.DB` struct.

The `.sql` file should contain one or more valid MySQL statements. To be able to execute multiple statements at a time, each statement must end with a `;`. If you are using option `2` above the SQL connection must use `multiStatements=true` to be able to run more than one statement at a time.

## Installation

Install the package with:

```
go get github.com/Storytel/go-mysql-seed
```

Once installed import it to your code:

```
import mysqlseed "github.com/Storytel/go-mysql-seed"
```

## Examples

This package is especially useful for testing. `go-mysql-seed` allows you to seed new database instances.

**Note:** The path to the `.sql` file is recommended to be absolute. The following snippet can be modified and used if you are uncertain.

```
fmt.Sprintf("%s/src/github.com/Storytel/go-example-service/%s.sql", filepath.ToSlash(os.Getenv("GOPATH")), seedFileName)
```

Below is a typical and simple example using `go-docker-initiator` in a test using the `*sql.DB` struct.

```
package example_test

import (
	"log"
	"testing"

	mysqlseed "github.com/Storytel/go-mysql-seed"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseIntegration(t *testing.T) {
	// Establish a database connection to the exposed environment variables
	db, err := InitAndCreateDatabase()
	assert.NoError(t, err)
	defer db.Close()

	// Run DB seeds
	err = mysqlseed.ApplySeedWithDB(db, "/path/to/my-seed-file.sql")
	assert.NoError(t, err)

	// Setup your service and inject the database
	exampleService := ExampleService{
		Db: db,
	}

	// Test your integration
	_, err = exmapleService.Create()
	assert.NoError(t, err)
}
```

You can also apply a seed with the `MySQL Command Line Tool`.

```
package example_test

import (
	"log"
	"testing"

	mysqlseed "github.com/Storytel/go-mysql-seed"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseIntegration(t *testing.T) {
	// Establish a database connection to the exposed environment variables
	db, err := InitAndCreateDatabase()
	assert.NoError(t, err)
	defer db.Close()

	// Run DB seeds
	err = mysqlseed.ApplySeedWithCmd("127.0.0.1:3306", "root", "password", "testdb", "/path/to/my-seed-file.sql")
	assert.NoError(t, err)

	// Setup your service and inject the database
	exampleService := ExampleService{
		Db: db,
	}

	// Test your integration
	_, err = exmapleService.Create()
	assert.NoError(t, err)
}
```

## Storytel Go

https://github.com/Storytel/go-docker-initiator - Utility for starting docker containers from Go.
