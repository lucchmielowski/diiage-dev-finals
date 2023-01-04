package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"gopkg.in/yaml.v2"
)

type CatalogService struct {
	rdb *redis.Client
}

var ctx = context.Background()

func (c *CatalogService) Handler(w http.ResponseWriter, r *http.Request) {
	sku := r.URL.Query().Get("sku")
	fmt.Println(sku)

	if sku == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "sku parameter is required")
		return
	}

	val, err := c.rdb.Get(ctx, sku).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	io.WriteString(w, val)
}

func NewService() *CatalogService {
	// Read yaml file
	yamlFile, err := os.ReadFile(os.Getenv("SETUP_FILE_PATH"))

	if err != nil {
		panic(err)
	}

	startupConfig := make(map[string]string)

	err = yaml.Unmarshal(yamlFile, &startupConfig)

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	rdb.FlushDB(ctx)
	for key, val := range startupConfig {
		err := rdb.Set(ctx, key, val, 0).Err()
		if err != nil {
			panic(err)
		}
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
