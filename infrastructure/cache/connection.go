package cache

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func Connect() (*redis.Client, error) {
	url := os.Getenv("CACHE_URL")
	if url == "" {
		log.Fatalln("missing env variable CACHE URL")
	}

	options, err := redis.ParseURL(url)

	if err != nil {
		return nil, err
	}

	client := redis.NewClient(options)

	_, err = client.Ping(context.Background()).Result()

	if err != nil {
		return nil, err
	}

	return client, nil
}
