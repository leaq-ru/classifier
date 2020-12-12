package classifierimpl

import (
	"context"
	pbClassifier "github.com/nnqq/scr-proto/codegen/go/classifier"
)

func (s *server) Predict(
	_ context.Context,
	req *pbClassifier.PredictRequest,
) (
	res *pbClassifier.PredictResponse,
	err error,
) {
	res = &pbClassifier.PredictResponse{
		CategoryClass: string(s.classifier.Predict(req.GetHtml())),
	}
	return
}
