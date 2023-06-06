package sqlite

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/quii/go-fakes-and-contracts/adapters/persistence/sqlite/ent"
	"log"
)

func NewSQLiteClient() *ent.Client {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	//client, err := ent.Open(dialect.SQLite, "file.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
