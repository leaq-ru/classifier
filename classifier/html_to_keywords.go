package classifier

import (
	"github.com/kljensen/snowball/russian"
	"regexp"
	"strings"
)

var rxRu = regexp.MustCompile(`[аАбБвВгГдДеЕёЁжЖзЗиИйЙкКлЛмМнНоОпПрРсСтТуУфФхХцЦчЧшШщЩъЪыЫьЬэЭюЮяЯ]+`)

func HTMLToKeywords(html string) (keywords []string) {
	words := rxRu.FindAllString(html, -1)

	for _, word := range words {
		keywords = append(keywords, strings.ToLower(russian.Stem(word, false)))
	}
	return
}
