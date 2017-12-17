package main

import (
	"sync"
	rateType "vcm/rate_type"
)

type Rate struct {
	rateType.BTC
	rateType.XRP
	rateType.XEM
	rateType.ETH
	rateType.ETC
	rateType.LTC
	rateType.FCT
	rateType.XMR
	rateType.REP
	rateType.ZEC
	rateType.DASH
	rateType.BCH
}

func (r *Rate) Update() {
	wg := sync.WaitGroup{}
	wg.Add(12)
	go r.BTC.UpdateRate(&wg)
	go r.XRP.UpdateRate(&wg)
	go r.XEM.UpdateRate(&wg)
	go r.ETH.UpdateRate(&wg)
	go r.ETC.UpdateRate(&wg)
	go r.LTC.UpdateRate(&wg)
	go r.FCT.UpdateRate(&wg)
	go r.XMR.UpdateRate(&wg)
	go r.REP.UpdateRate(&wg)
	go r.ZEC.UpdateRate(&wg)
	go r.DASH.UpdateRate(&wg)
	go r.BCH.UpdateRate(&wg)
	wg.Wait()
}
