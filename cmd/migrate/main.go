package main

import (
	"blockchain/internal/config"
	"blockchain/internal/db"
	"blockchain/internal/db/migrator"
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}

	command := "up"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	cfg := config.LoadConfig()
	ctx := context.Background()
	pool := db.NewPostgresPool(ctx, cfg)
	defer pool.Close()

	switch command {
	case "up":
		if err := migrator.Up(ctx, pool, "internal/db/migrations"); err != nil {
			log.Fatal(err)
		}
		log.Println("migrations applied successfully")
	case "down":
		fs := flag.NewFlagSet("down", flag.ExitOnError)
		steps := fs.Int("steps", 1, "number of migrations to roll back")
		_ = fs.Parse(os.Args[2:])

		if err := migrator.Down(ctx, pool, "internal/db/migrations", *steps); err != nil {
			log.Fatal(err)
		}
		log.Printf("rolled back %d migration(s)\n", *steps)
	case "status":
		lines, err := migrator.Status(ctx, pool, "internal/db/migrations")
		if err != nil {
			log.Fatal(err)
		}

		if len(lines) == 0 {
			fmt.Println("no migrations found")
			return
		}

		for _, line := range lines {
			fmt.Println(line)
		}
	default:
		log.Fatalf("unknown command %q. use: up | down [-steps=N] | status", command)
	}
}
