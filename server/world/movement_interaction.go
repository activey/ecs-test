package world

//import (
//	"fmt"
//)
//
//type MovementRequest struct {
//	Source        *Participant
//	MovementDelta SceneMovement
//
//	ResultListener chan *MovementResult
//	ErrorListener  chan error
//}
//
//func (r *MovementRequest) Process(world *World) (*MovementResult, error) {
//	if r.Source.MovementDistance == 0 {
//		return nil, fmt.Errorf("not enough movement distance")
//	}
//	movementPoints, err := world.Scene.DeltaMoveParticipant(r.Source, r.MovementDelta)
//	if err != nil {
//		return nil, err
//	}
//
//	r.Source.CommitMovementDistance(movementPoints)
//	r.Source.WaitForNextTurn()
//
//	// TODO check difficult terrain and trigger Dexterity check on ice
//
//	res := &MovementResult{
//		Source:      r.Source,
//		NewPosition: r.Source.Position,
//	}
//	return res, nil
//}
//
//type MovementResult struct {
//	Source      *Participant
//	NewPosition SceneCoordinates
//}
//
//func (m *MovementResult) Accept(consumer InteractionResultConsumer) {
//	consumer.ConsumeMovementResult(m)
//}
