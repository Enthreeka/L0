package amqp

type Publish interface {
	Publish(subject string)
}

type Subscribe interface {
	Subscribe(subject string)
}
