package classifier

import "github.com/jbrukh/bayesian"

func (c Classifier) Predict(html string) bayesian.Class {
	keys := HTMLToKeywords(html)
	_, i, _ := c.model.LogScores(keys)
	return c.model.Classes[i]
}
