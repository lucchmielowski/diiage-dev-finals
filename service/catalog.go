package main

import (
	"context"
	"net/http"
	"io"
	"os"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type CatalogService struct {
	rdb *redis.Client
}

var ctx = context.Background()

func (c *CatalogService) Handler(w http.ResponseWriter, r *http.Request) {
    val, err := c.rdb.Get(ctx, "diiage").Result()
    if err != nil {
        panic(err)
    }
    io.WriteString(w, val)
}

func NewService() *CatalogService {
    rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	err := rdb.Set(ctx, "diiage", "it works !", 0).Err()
	if err != nil {
		panic(err)
	}

    return &CatalogService{
        rdb,
    }
}

func main() {
    service := NewService()
	http.HandleFunc("/get-key", service.Handler)
	fmt.Println("handle")
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
        panic(err)
	}
}
