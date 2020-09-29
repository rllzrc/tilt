package server

import "github.com/tilt-dev/tilt/pkg/model"

type AppendToTriggerQueueAction struct {
	Name   model.ManifestName
	Reason model.BuildReason
}

func (AppendToTriggerQueueAction) Action() {}

type SetTiltfileArgsAction struct {
	Args []string
}

func (SetTiltfileArgsAction) Action() {}
