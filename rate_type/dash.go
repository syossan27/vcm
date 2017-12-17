package rate_type

import (
	"sync"
	"vcm/api"
)

type DASH string

func (d *DASH) UpdateRate(wg *sync.WaitGroup) {
	defer wg.Done()

	rate, err := api.FetchRate("dash_jpy")
	if err != nil {
		panic(err)
	}

	*d = DASH(rate)
}
