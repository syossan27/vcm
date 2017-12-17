package rate_type

import (
	"sync"
	"vcm/api"
)

type ETC string

func (e *ETC) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("etc_jpy")
	if err != nil {
		panic(err)
	}

	*e = ETC(rate)
}
