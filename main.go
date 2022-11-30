package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/abassGarane/work/ports/repositories/mongo"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	defer cancel()
	mongoURL := os.Getenv("MONGO_URL")
	mongoDB := os.Getenv("MONGO_DB")
	mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	repo, err := mongo.NewMongoRepository(mongoURL, mongoDB, mongoTimeout, ctx)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		j, _ := repo.GetAll()
		fmt.Println(j)
	}

}
