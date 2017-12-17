package rate_type

import (
	"sync"
	"vcm/api"
)

type XEM string

func (x *XEM) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("xem_jpy")
	if err != nil {
		panic(err)
	}

	*x = XEM(rate)
}
