package EasyGoQ_Example

import (
	"github.com/ptncafe/EasyGoQ"
	"os"
)

func PublishQueue() error {
	rabbitMq, err := EasyGoQ.NewRabbitMqClient(os.Getenv("RABBITMQ_URL"), EasyGoQ.NewDefaultLogger(), true)
	if err != nil {
		return err
	}
	rabbitMq.Publish()
}
