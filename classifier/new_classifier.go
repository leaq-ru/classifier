package classifier

import (
	"github.com/jbrukh/bayesian"
	"os"
	"path"
)

func NewClassifier() (cl Classifier, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	bCL, err := bayesian.NewClassifierFromFile(path.Join(wd, "classifier", "model"))
	if err != nil {
		return
	}

	cl = Classifier{
		model: bCL,
	}
	return
}
