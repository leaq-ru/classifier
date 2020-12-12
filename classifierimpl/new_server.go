package classifierimpl

import (
	"github.com/nnqq/scr-classifier/classifier"
	"github.com/rs/zerolog"
)

func NewServer(logger zerolog.Logger, cls classifier.Classifier) *server {
	return &server{
		logger:     logger,
		classifier: cls,
	}
}
