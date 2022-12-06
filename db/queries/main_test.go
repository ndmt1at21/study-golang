package queries

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var queries IUserQueries

func TestMain(m *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	mysqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		panic(err)
	}

	dbEndpoint, err := mysqlC.Endpoint(ctx, "3306/tcp")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open(dbEndpoint))
	if err != nil {
		panic(err)
	}

	queries = NewUserQueries(db)

	m.Run()
}
