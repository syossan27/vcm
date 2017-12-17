package rate_type

import (
	"sync"
	"vcm/api"
)

type LISK string

func (l *LISK) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("lisk_jpy")
	if err != nil {
		panic(err)
	}

	*l = LISK(rate)
}
