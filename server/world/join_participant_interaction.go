package world

import "ecs-test/server/character/domain"

type JoinParticipantRequest struct {
	Character *domain.Character
	//InitialPosition SceneCoordinates

	ResultListener chan *JoinParticipantResult
}

func (r *JoinParticipantRequest) Process(world *World) (*JoinParticipantResult, error) {
	//id, err := world.NewId()
	//if err != nil {
	//	return nil, err
	//}
	//
	//participant, err := NewParticipant(r.Character, ParticipantId(id))
	//if err != nil {
	//	return nil, err
	//}
	////
	//err = world.Scene.AddParticipant(participant, r.InitialPosition)
	//if err != nil {
	//	return nil, err
	//}
	//
	//result := &JoinParticipantResult{
	//	Participant:     participant,
	//	InitialPosition: r.InitialPosition,
	//}
	//return result, nil
	return nil, nil
}

type JoinParticipantResult struct {
	//Participant     *Participant
	//InitialPosition SceneCoordinates
}

func (r *JoinParticipantResult) Accept(consumer InteractionResultConsumer) {
	consumer.ConsumeJoinResult(r)
}
