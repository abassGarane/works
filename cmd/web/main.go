package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/abassGarane/work/domain"
	"github.com/abassGarane/work/ports/api"
	"github.com/abassGarane/work/ports/repositories/mongo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	// mongoURL := flag.String("MONGO_URL", "mongodb://root:root@localhost:27017/work", "Mongodb connection string")
	// mongoTimeout := flag.Int("MONGO_TIMEOUT", 10, "Mongo timeout")
	// mongoDB := flag.String("MONGO_DB", "work", "Mongo db string")

	//Repo
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Hour*180))
	defer cancel()
	mongoURL := os.Getenv("MONGO_URL")
	mongoDB := os.Getenv("MONGO_DB")
	mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	repo, err := mongo.NewMongoRepository(mongoURL, mongoDB, mongoTimeout, ctx)
	if err != nil {
		log.Fatal(err)
	}
	// service
	service := domain.NewJobService(repo)

	// application
	handler := api.NewJobHandler(service)
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
	// app.Use(limiter.New())
	app.Use(cache.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Access-Control-Allow-Origin",
	}))
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})
	// endpoints
	app.Get("/:id", handler.Get)
	app.Get("/", handler.GetAll)
	app.Post("/", handler.AddJob)
	app.Patch("/:id", handler.UpdateJob)

	//app init
	errChan := make(chan error, 2)
	go func() {
		fmt.Printf("Listening on %s\n", httpPort())
		errChan <- app.Listen(httpPort())
	}()

	// sig stops
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT)
		errChan <- fmt.Errorf("%s", <-sigChan)
	}()

	fmt.Printf("Terminated %s\n", <-errChan)
}
func httpPort() string {
	port := "8000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
