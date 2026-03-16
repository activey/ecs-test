package world

import (
	"ecs-test/server/character/domain"
	"fmt"
	"time"
)

type World struct {
	//Scene          *Scene
	Interactions   *Interactions
	ResultConsumer InteractionResultConsumer

	//IdGenerator *IdGenerator
}

func NewWorld(resultConsumer InteractionResultConsumer) *World {
	return &World{
		//Scene: scene,

		Interactions:   NewWorldInteractions(10),
		ResultConsumer: resultConsumer,
		//IdGenerator:    NewIdGenerator(),
	}
}

//func (w *World) NewId() (string, error) {
//	return w.IdGenerator.NewId()
//}

func (w *World) RunWorldInteractionLoop() {
	ticker := time.NewTicker(1 * time.Second)

	defer w.Interactions.Close()
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			w.update()
		case joinRequest := <-w.Interactions.JoinChan:
			joinResult, err := w.processParticipantJoinRequest(joinRequest)

			if err != nil {
				continue
			}
			joinResult.Accept(w.ResultConsumer)
			joinRequest.ResultListener <- joinResult

			//case moveRequest := <-w.Interactions.MovementChan:
			//moveResult,w err := w.processMoveRequest(moveRequest)
			//
			//if err != nil {
			//	moveRequest.ErrorListener <- err
			//	continue
			//}
			//moveResult.Accept(w.ResultConsumer)
			//moveRequest.ResultListener <- moveResult

			//case actionRequest := <-w.Interactions.ActionChan:
			//	actionResult, err := w.processActionRequest(actionRequest)
			//	if err != nil {
			//		continue
			//	}
			//	actionResult.Accept(w.ResultConsumer)
		}
	}
}

func (w *World) JoinCharacter(character *domain.Character, position Coordinates) chan *JoinParticipantResult {
	resultChan := make(chan *JoinParticipantResult)
	w.Interactions.JoinChan <- JoinParticipantRequest{
		Character: character,
		//InitialPosition: position,
		ResultListener: resultChan,
	}
	return resultChan
}

func (w *World) processParticipantJoinRequest(request JoinParticipantRequest) (*JoinParticipantResult, error) {
	return request.Process(w)
}

//func (w *World) MoveParticipant(predicate ParticipantPredicate, delta SceneMovement) (chan *MovementResult, chan error) {
//	resultChan := make(chan *MovementResult)
//	errorChan := make(chan error)
//
//	participant := w.Scene.FindParticipant(predicate)
//	if participant == nil {
//		errorChan <- errors.New("unable to find participant")
//		return resultChan, errorChan
//	}
//
//	w.Interactions.MovementChan <- MovementRequest{
//		Source:        participant,
//		MovementDelta: delta,
//
//		ResultListener: resultChan,
//		ErrorListener:  errorChan,
//	}
//
//	return resultChan, errorChan
//}
//
//func (w *World) ExecuteAction(
//	participant *Participant,
//	action *Action,
//	arguments ActionArguments,
//) {
//	w.Interactions.ActionChan <- ActionRequest{
//		Source:    participant,
//		Action:    action,
//		Arguments: arguments,
//	}
//}

//func (w *World) processActionRequest(request ActionRequest) (*ActionResult, error) {
//	//if !battle.TurnCoordinator.IsParticipantTurn(request.Source) {
//	//	return nil, fmt.Errorf("it's not participant %s's turn", request.Source.Character.Name)
//	//}
//	return request.Process(w)
//}
//
//func (w *World) processMoveRequest(request MovementRequest) (*MovementResult, error) {
//	//if !battle.TurnCoordinator.IsParticipantTurn(request.Source) {
//	//	return nil, fmt.Errorf("it's not participant %s's turn", request.Source.Character.Name)
//	//}
//
//	return request.Process(w)
//}

func (w *World) update() {
	fmt.Println("Tik... Tak...")

	timestamp := time.Now().Round(time.Millisecond).UnixNano() / 1e6

	//var wg sync.WaitGroup
	//participants := w.Scene.AllParticipants()
	//wg.Add(len(participants))
	//
	//for _, participant := range participants {
	//	go func(p *Participant) {
	//		defer wg.Done()
	//		p.Update(timestamp)
	//	}(participant)
	//}
	//
	//wg.Wait()

	w.ResultConsumer.ConsumeTickResult(
		&TickResult{
			Timestamp: timestamp,
		},
	)
}
