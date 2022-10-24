//go:build ignore

package main

import (
	"context"
	"log"
	"os"

	_ "ariga.io/atlas/sql/sqlite"
	_ "github.com/mattn/go-sqlite3"

	"atlas-sqlite-modereplay/ent/migrate"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}

	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDialect(dialect.SQLite),
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	ctx := context.Background()
    if err = migrate.NamedDiff(ctx, "sqlite3://main.db?_fk=1", os.Args[1], opts...); err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
