package serviceImpl

import (
	"context"
	"yama.io/yamaIterativeE/internal/iteration/pipeline"
)

type PileLineService struct{}

func (p *PileLineService) StartYaMaPipeLine(ctx context.Context, request *StartYaMaPipeLineRequest) (*StartYaMaPipeLineResponse, error) {
	return &StartYaMaPipeLineResponse{
		Success: true,
	}, nil
}

func (p *PileLineService) PassMergeRequestCodeReview(ctx context.Context, request *PassMergeRequestCodeReviewRequest) (*PassMergeRequestCodeReviewResponse, error) {
	pipeline.PassStep(request.ActionId, request.StageId, request.StepId)

	return &PassMergeRequestCodeReviewResponse{
		Success: true,
	}, nil
}

func (p *PileLineService) RestartYaMaPipeLine(ctx context.Context, request *RestartYaMaPipeLineRequest) (*RestartYaMaPipeLineResponse, error) {
	baseResponse := &RestartYaMaPipeLineResponse{}
	err := pipeline.ReStartBasicMRPipelineWithArgs(request.PipelineId, request.IterationId, request.ActionId, request.ActorName, request.SourceBranch, request.TargetBranch,
		request.Env, request.MrInfo, request.AppOwner, request.AppName, request.MrCodeReviews)
	if err==nil {
		baseResponse.Success = true
	}
	return baseResponse, err
}


