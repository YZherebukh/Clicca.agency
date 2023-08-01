package quote

import (
	"math/rand"
)

type Quote struct {
	quotes []string
}

func New() Quote {
	return Quote{
		quotes: []string{`You create your own opportunities. ...`,
			`Never break your promises. ...`,
			`You are never as stuck as you think you are. ...`,
			`Happiness is a choice. ...`,
			`Habits develop into character. ...`,
			`Be happy with who you are. ...`,
			`Don't seek happiness–create it. ...`,
			`If you want to be happy, stop complaining.`},
	}
}

// Random returnes random quote
func (q Quote) Random() string {
	return q.quotes[rand.Intn(len(q.quotes))]
}
