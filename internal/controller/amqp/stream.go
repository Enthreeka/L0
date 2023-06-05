package amqp

type Publish interface {
	Publish(subject string) error
}

type Subscribe interface {
	Subscribe(subject string) error
}
