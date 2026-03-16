package world

type InteractionResult interface {
	Accept(consumer InteractionResultConsumer)
}

type InteractionResultConsumer interface {
	ConsumeJoinResult(result *JoinParticipantResult)
	//ConsumeActionResult(result *ActionResult)
	//ConsumeMovementResult(result *MovementResult)
	ConsumeTickResult(result *TickResult)
}

type Interactions struct {
	JoinChan chan JoinParticipantRequest
	//ActionChan   chan ActionRequest   // Channel to receive actions
	//MovementChan chan MovementRequest // Channel to receive movement requests
	//EndTurnChan  chan EndTurnRequest  // Channel to receive end turn messages
	//EndBattleChan chan bool            // Channel to signal the end of the battle
}

func NewWorldInteractions(bufferSize int) *Interactions {
	return &Interactions{
		JoinChan: make(chan JoinParticipantRequest, bufferSize),
		//ActionChan:   make(chan ActionRequest, bufferSize),
		//MovementChan: make(chan MovementRequest, bufferSize),
		//EndTurnChan:  make(chan EndTurnRequest, bufferSize),
		//EndBattleChan: make(chan bool),
	}
}

func (pi *Interactions) Close() {
	close(pi.JoinChan)
	//close(pi.ActionChan)
	//close(pi.MovementChan)
	//close(pi.EndTurnChan)
	//close(pi.EndBattleChan)
}
