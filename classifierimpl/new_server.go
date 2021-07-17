package classifierimpl

import (
	"github.com/leaq-ru/classifier/classifier"
	"github.com/rs/zerolog"
)

func NewServer(logger zerolog.Logger, cls classifier.Classifier) *server {
	return &server{
		logger:     logger,
		classifier: cls,
	}
}
