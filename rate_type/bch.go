package rate_type

import (
	"sync"
	"vcm/api"
)

type BCH string

func (b *BCH) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("bch_jpy")
	if err != nil {
		panic(err)
	}

	*b = BCH(rate)
}
