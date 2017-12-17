package rate_type

import (
	"sync"
	"vcm/api"
)

type ZEC string

func (z *ZEC) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("zec_jpy")
	if err != nil {
		panic(err)
	}

	*z = ZEC(rate)
}
