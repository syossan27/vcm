package rate_type

import (
	"sync"
	"vcm/api"
)

type XMR string

func (x *XMR) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("xmr_jpy")
	if err != nil {
		panic(err)
	}

	*x = XMR(rate)
}
