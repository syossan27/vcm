package rate_type

import (
	"sync"
	"vcm/api"
)

type BTC string

func (b *BTC) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("btc_jpy")
	if err != nil {
		panic(err)
	}

	*b = BTC(rate)
}

func (b *BTC) UpdateBalance(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("btc_jpy")
	if err != nil {
		panic(err)
	}

	*b = BTC(rate)
}
