package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"demo/internal/db"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/go-pg/pg/v10"
)

const (
	port int = 5678
)

func main() {
	// Create in-memory database
	embeddedDB := createEmbeddedDatabase()
	defer embeddedDB.Stop()

	// Run migration script(s) - dbmate up
	setupDatabase()

	// Run dummy code
	database := connectToDatabase()
	repository := db.NewEntityRepository(database)

	// - insert
	e := &db.Entity{
		ID:          "id1",
		Description: "desc",
	}
	repository.CreateEntity(e)

	// - get
	es, _ := repository.GetEntities()
	fmt.Println("Fetched entities:")
	for _, e := range es {
		fmt.Println(e)
	}
}

func createEmbeddedDatabase() *embeddedpostgres.EmbeddedPostgres {
	edb := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().Port(uint32(port)))
	err := edb.Start()
	if err != nil {
		panic(err)
	}
	return edb
}

func setupDatabase() {
	u, _ := url.Parse("postgresql://postgres:postgres@localhost:" + strconv.Itoa(port) + "/postgres?sslmode=disable")
	db := dbmate.New(u)
	db.AutoDumpSchema = false
	err := db.CreateAndMigrate()
	if err != nil {
		panic(err)
	}
}

func connectToDatabase() *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:        "localhost:" + strconv.Itoa(port),
		User:        "postgres",
		Password:    "postgres",
		Database:    "postgres",
		DialTimeout: time.Second * 5,
	})
}
