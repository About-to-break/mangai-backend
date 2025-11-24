package services

type QueueService interface {
	Publish(queue string, body []byte) error
}
