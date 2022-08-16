package sidekiq

import (
	"context"
	"encoding/json"

	"github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/pkg/redis"
)

var ctx = context.Background()

type sidekiqJob struct {
	Class string   `json:"class"`
	Args  []string `json:"args"`
	Retry bool     `json:"retry"`
	Queue string   `json:"queue"`
}

func Push(queue string, class string, args ...string) error {
	job := sidekiqJob{
		Class: class,
		Args:  args,
		Queue: queue,
		Retry: true,
	}

	redisClient, err := redis.GetRedis()
	if err != nil {
		return err
	}

	jobBytes, err := json.Marshal(job)
	if err != nil {
		return err
	}

	_, err = redisClient.RPush(ctx, "queue:"+queue, jobBytes).Result()
	return err
}
