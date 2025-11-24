package services

type QueueService interface {
	Publish(exchange string, key string, body []byte) error
}
