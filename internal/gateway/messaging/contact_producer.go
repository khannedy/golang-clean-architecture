package messaging

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"golang-clean-architecture/internal/model"
)

type ContactProducer struct {
	Producer[*model.ContactEvent]
}

func NewContactProducer(producer *kafka.Producer, log *logrus.Logger) *ContactProducer {
	return &ContactProducer{
		Producer: Producer[*model.ContactEvent]{
			Producer: producer,
			Topic:    "contacts",
			Log:      log,
		},
	}
}
