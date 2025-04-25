import "github.com/pressly/goose/v3"

func runMigrations() {
	err := goose.Up(db, "./db/migrations")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}
