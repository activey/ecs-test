package world

type TickResult struct {
	Timestamp int64
}

func (r *TickResult) Accept(consumer InteractionResultConsumer) {
	consumer.ConsumeTickResult(r)
}
