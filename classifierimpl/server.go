package classifierimpl

import (
	"github.com/leaq-ru/classifier/classifier"
	pbClassifier "github.com/leaq-ru/proto/codegen/go/classifier"
	"github.com/rs/zerolog"
)

type server struct {
	pbClassifier.UnimplementedClassifierServer
	logger     zerolog.Logger
	classifier classifier.Classifier
}
