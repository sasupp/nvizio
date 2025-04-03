package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"xtrinio.com/db"
)

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func setupCloseHandler(db *db.Db, app *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		db.Close()
		app.Shutdown()
	}()
}

func Server(hosts []string) {

	cfg := db.DbConfig{
		Hosts: hosts,
	}
	db := db.NewDb(cfg)
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	// Create a new engine
	// engine := html.New("./views", ".html")
	// engine.ShouldReload = true
	// engine.AddFunc(
	// 	// add unescape function
	// 	"unescape", func(s string) template.HTML {
	// 		return template.HTML(s)
	// 	},
	// )

	srv := fiber.New(fiber.Config{
		// Views: engine,
	})

	srv.Static("/", "./public")

	SetupNseFeedRoutes(srv, db)
	// SetupFilingRoutes(srv, db)

	setupCloseHandler(db, srv)
	log.Fatal(srv.Listen("localhost:3003"))
}
