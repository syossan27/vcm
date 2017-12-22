package rate_type

import (
	"sync"
	"vcm/api"
)

type XRP string

func (x *XRP) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("xrp_jpy")
	if err != nil {
		panic(err)
	}

	*x = XRP(rate)
}
