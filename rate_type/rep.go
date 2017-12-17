package rate_type

import (
	"sync"
	"vcm/api"
)

type REP string

func (r *REP) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("rep_jpy")
	if err != nil {
		panic(err)
	}

	*r = REP(rate)
}
