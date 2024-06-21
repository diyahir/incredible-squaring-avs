package types

import (
	"errors"

	cstaskmanager "anzen-avs/contracts/bindings/IncredibleSquaringTaskManager"
)

type TaskResponseData struct {
	TaskResponse              cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
	TaskResponseMetadata      cstaskmanager.IIncredibleSquaringTaskManagerTaskResponseMetadata
	NonSigningOperatorPubKeys []cstaskmanager.BN254G1Point
}

var (
	NoErrorInTaskResponse = errors.New("100. Task response is valid")
)
