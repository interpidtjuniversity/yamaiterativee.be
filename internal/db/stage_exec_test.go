package db

import (
	"testing"
)

func Test_RunStage(t *testing.T) {
	NewEngine()
	var devStageExecs []StageExec
	var itgStageExecs []StageExec
	var preStageExecs []StageExec

	devStageExecs = append(devStageExecs,StageExec{ActId: 1, StageId: 1, StageStartUnix: 1,StageEndUnix: 2})
	devStageExecs = append(devStageExecs,StageExec{ActId: 1, StageId: 2, StageStartUnix: 1,StageEndUnix: 2})
	devStageExecs = append(devStageExecs,StageExec{ActId: 1, StageId: 3, StageStartUnix: 1,StageEndUnix: 2})
	devStageExecs = append(devStageExecs,StageExec{ActId: 1, StageId: 4, StageStartUnix: 1,StageEndUnix: 2})
	devStageExecs = append(devStageExecs,StageExec{ActId: 1, StageId: 5, StageStartUnix: 1,StageEndUnix: 2})
	devStageExecs = append(devStageExecs,StageExec{ActId: 1, StageId: 6, StageStartUnix: 1,StageEndUnix: 2})
	devStageExecs = append(devStageExecs,StageExec{ActId: 1, StageId: 7, StageStartUnix: 1,StageEndUnix: 2})

	itgStageExecs = append(itgStageExecs,StageExec{ActId: 2, StageId: 1,StageStartUnix: 1,StageEndUnix: 2})
	itgStageExecs = append(itgStageExecs,StageExec{ActId: 2, StageId: 2,StageStartUnix: 1,StageEndUnix: 2})
	itgStageExecs = append(itgStageExecs,StageExec{ActId: 2, StageId: 3,StageStartUnix: 1,StageEndUnix: 2})
	itgStageExecs = append(itgStageExecs,StageExec{ActId: 2, StageId: 4,StageStartUnix: 1,StageEndUnix: 2})
	itgStageExecs = append(itgStageExecs,StageExec{ActId: 2, StageId: 5,StageStartUnix: 1,StageEndUnix: 2})
	itgStageExecs = append(itgStageExecs,StageExec{ActId: 2, StageId: 6,StageStartUnix: 1,StageEndUnix: 2})
	itgStageExecs = append(itgStageExecs,StageExec{ActId: 2, StageId: 7,StageStartUnix: 1,StageEndUnix: 2})

	preStageExecs = append(preStageExecs,StageExec{ActId: 3, StageId: 1,StageStartUnix: 1,StageEndUnix: 2})
	preStageExecs = append(preStageExecs,StageExec{ActId: 3, StageId: 2,StageStartUnix: 1,StageEndUnix: 2})
	preStageExecs = append(preStageExecs,StageExec{ActId: 3, StageId: 3,StageStartUnix: 1,StageEndUnix: 2})
	preStageExecs = append(preStageExecs,StageExec{ActId: 3, StageId: 4,StageStartUnix: 1,StageEndUnix: 2})
	preStageExecs = append(preStageExecs,StageExec{ActId: 3, StageId: 5,StageStartUnix: 1,StageEndUnix: 2})
	preStageExecs = append(preStageExecs,StageExec{ActId: 3, StageId: 6,StageStartUnix: 1,StageEndUnix: 2})
	preStageExecs = append(preStageExecs,StageExec{ActId: 3, StageId: 7,StageStartUnix: 1,StageEndUnix: 2})

	for _, v := range devStageExecs {
		x.Insert(v)
	}
	for _, v := range itgStageExecs {
		x.Insert(v)
	}
	for _, v := range preStageExecs {
		x.Insert(v)
	}
}
