package rate_type

import (
	"sync"
	"vcm/api"
)

type LTC string

func (l *LTC) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("ltc_jpy")
	if err != nil {
		panic(err)
	}

	*l = LTC(rate)
}
