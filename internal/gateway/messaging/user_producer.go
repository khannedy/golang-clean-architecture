package messaging

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/internal/model"
)

type UserProducer struct {
	Producer[*model.UserEvent]
}

func NewUserProducer(producer *kafka.Producer, log *logrus.Logger) *UserProducer {
	return &UserProducer{
		Producer: Producer[*model.UserEvent]{
			Producer: producer,
			Topic:    "users",
			Log:      log,
		},
	}
}
