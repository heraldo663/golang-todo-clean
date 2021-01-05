package main

import (
	"fmt"
	"log"
	"numtostr/gotodo/config"
	"numtostr/gotodo/config/database"
	"numtostr/gotodo/modules/auth"
	"numtostr/gotodo/modules/todo"
	"numtostr/gotodo/shared/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	database.Connect()
	database.Migrate(&auth.User{}, &todo.Todo{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	auth.AuthRoutes(app)
	todo.TodoRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(fmt.Sprintf(":%v", config.PORT)); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	db, _ := database.DB.DB()
	db.Close()
	// redisConn.Close()

}
