# go-mysql-seed

Utility for seeding a MySQL database with Go.
Useful during for testing.

This package supports 2 types of MySQL seeding, both requires a path to a `.sql` file.

1. Seeding via `MySQL Command Line Tool`:
   Initialize a import using your locally installed MySQL CLT.

2. Seeding via `*sql.DB` struct: Start an import of the specified file using a `*sql.DB` struct.

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

This package is especially useful for testing. With `go-mysql-seed` you can seed new database instances.

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
	err = mysqlseed.ApplySeedWithDB(db, "my-seed-file.sql")
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
	err = mysqlseed.ApplySeedWithCmd("127.0.0.1:3306", "root", "password", "testdb", "my-seed-file.sql")
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
