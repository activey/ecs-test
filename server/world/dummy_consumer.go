package world

type DummyConsumer struct {
}

func NewDummyConsumer() InteractionResultConsumer {
	return &DummyConsumer{}
}

func (d DummyConsumer) ConsumeJoinResult(result *JoinParticipantResult) {
}

func (d DummyConsumer) ConsumeTickResult(result *TickResult) {
}
