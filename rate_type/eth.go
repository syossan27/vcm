package rate_type

import (
	"sync"
	"vcm/api"
)

type ETH string

func (e *ETH) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("eth_jpy")
	if err != nil {
		panic(err)
	}

	*e = ETH(rate)
}
