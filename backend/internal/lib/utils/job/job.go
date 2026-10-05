package job

import (
	"github.com/Prajwal-Neupane/go-boilerplate/internal/config"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
)

type JobService struct {
	Client *asynq.Client
	server *asynq.Server
	logger *zerolog.Logger
}

func NewJobService(logger *zerolog.Logger, cfg *config.Config) *JobService {
	redisAddr := cfg.Redis.Address

	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: redisAddr,
	})

	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6, //Higher priority queue for important emails
				"default":  3, //Default priority for most emails
				"low":      1, //Lower priority for non-urgent emails
			},
		},
	)
	return &JobService{
		Client: client,
		server: server,
		logger: logger,
	}
}
