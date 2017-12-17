package rate_type

import (
	"sync"
	"vcm/api"
)

type FCT string

func (f *FCT) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("fct_jpy")
	if err != nil {
		panic(err)
	}

	*f = FCT(rate)
}
