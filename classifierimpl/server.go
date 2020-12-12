package classifierimpl

import (
	"github.com/nnqq/scr-classifier/classifier"
	pbClassifier "github.com/nnqq/scr-proto/codegen/go/classifier"
	"github.com/rs/zerolog"
)

type server struct {
	pbClassifier.UnimplementedClassifierServer
	logger     zerolog.Logger
	classifier classifier.Classifier
}
