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


