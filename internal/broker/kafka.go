package broker

import (
	"fmt"
	"net/http"

	"github.com/IBM/sarama"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func ProducerMessage(c echo.Context, value string) error {
	brokerList := viper.GetStringSlice("kafka.brokerList")
	topic := viper.GetString("kafka.topic")
	producer, err := sarama.NewSyncProducer(brokerList, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(value),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, fmt.Sprintf("Message is stored in topic(%s)/partition(%d)/offset(%d) for message %d", topic, partition, offset, value))
}
