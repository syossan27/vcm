package main

import (
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type (
	Termbox struct {
		Width  int
		Height int
		Rate
		Balance
		Selection
	}
)

const (
	minSelection = 15
	maxSelection = 16
)

var (
	color = map[string]termbox.Attribute{
		"default": termbox.ColorDefault,
		"white":   termbox.ColorWhite,
		"black":   termbox.ColorBlack,
		"magenta": termbox.ColorMagenta,
		"cyan":    termbox.ColorCyan,
		"red":     termbox.ColorRed,
		"green":   termbox.ColorGreen,
		"blue":    termbox.ColorBlue,
	}

	attr = map[string]termbox.Attribute{
		"underline": termbox.AttrUnderline,
	}

	style = map[string]termbox.Attribute{
		"fg": color["white"],
		"bg": color["default"],
	}
)

func NewTermbox() *Termbox {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	w, h := termbox.Size()
	return &Termbox{
		Width:     w,
		Height:    h,
		Rate:      Rate{},
		Selection: NewSelection(),
	}
}

func (t *Termbox) Display() {
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputAlt)

	t.UpdateData()
	t.Draw()

	ticker := time.NewTicker(3 * time.Second)
	stop := make(chan bool)

	go t.RefreshRateDraw(ticker, stop)

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				if t.Selection.Y-1 >= minSelection {
					t.Selection.Y--
					t.Draw()
				}
			case termbox.KeyArrowDown:
				if t.Selection.Y+1 <= maxSelection {
					t.Selection.Y++
					t.Draw()
				}
			case termbox.KeyEnter:
				t.Selection.Exec()
			case termbox.KeyCtrlC:
				ticker.Stop()
				close(stop)
				break loop
			case termbox.KeyCtrlD:
				ticker.Stop()
				close(stop)
				break loop
			}
		}
	}
}

func (t *Termbox) Draw() {
	termbox.Clear(color["default"], color["default"])
	t.RateDraw()
	t.BalanceDraw()
	t.CommandDraw()
	termbox.Flush()
}

func (t *Termbox) UpdateData() {
	t.Rate.Update()
	t.Balance.Update()
}

func (t *Termbox) RefreshRateDraw(ticker *time.Ticker, stop chan bool) {
	for {
		select {
		case <-ticker.C:
			t.UpdateData()
			t.Draw()
		case <-stop:
			break
		}
	}
}

func (t *Termbox) RateDraw() {
	t.Print(0, 0, color["default"], color["default"], "■レート")
	t.Print(0, 1, color["default"], color["default"], "BTC:  "+string(t.Rate.BTC))
	t.Print(0, 2, color["default"], color["default"], "XRP:  "+string(t.Rate.XRP))
	t.Print(0, 3, color["default"], color["default"], "XEM:  "+string(t.Rate.XEM))
	t.Print(0, 4, color["default"], color["default"], "ETH:  "+string(t.Rate.ETH))
	t.Print(0, 5, color["default"], color["default"], "ETC:  "+string(t.Rate.ETC))
	t.Print(0, 6, color["default"], color["default"], "LTC:  "+string(t.Rate.LTC))
	t.Print(0, 7, color["default"], color["default"], "FCT:  "+string(t.Rate.FCT))
	t.Print(0, 8, color["default"], color["default"], "XMR:  "+string(t.Rate.XMR))
	t.Print(0, 9, color["default"], color["default"], "REP:  "+string(t.Rate.REP))
	t.Print(0, 10, color["default"], color["default"], "ZEC:  "+string(t.Rate.ZEC))
	t.Print(0, 11, color["default"], color["default"], "DASH: "+string(t.Rate.DASH))
	t.Print(0, 12, color["default"], color["default"], "BCH:  "+string(t.Rate.BCH))
}

func (t *Termbox) BalanceDraw() {
	t.Print(30, 0, color["default"], color["default"], "■残高")
	t.Print(30, 1, color["default"], color["default"], "BTC:  "+string(t.Balance.BTC))
	t.Print(30, 2, color["default"], color["default"], "XRP:  "+string(t.Balance.XRP))
	t.Print(30, 3, color["default"], color["default"], "XEM:  "+string(t.Balance.XEM))
	t.Print(30, 4, color["default"], color["default"], "ETH:  "+string(t.Balance.ETH))
	t.Print(30, 5, color["default"], color["default"], "ETC:  "+string(t.Balance.ETC))
	t.Print(30, 6, color["default"], color["default"], "LTC:  "+string(t.Balance.LTC))
	t.Print(30, 7, color["default"], color["default"], "FCT:  "+string(t.Balance.FCT))
	t.Print(30, 8, color["default"], color["default"], "XMR:  "+string(t.Balance.XMR))
	t.Print(30, 9, color["default"], color["default"], "REP:  "+string(t.Balance.REP))
	t.Print(30, 10, color["default"], color["default"], "ZEC:  "+string(t.Balance.ZEC))
	t.Print(30, 11, color["default"], color["default"], "DASH: "+string(t.Balance.DASH))
	t.Print(30, 12, color["default"], color["default"], "BCH:  "+string(t.Balance.BCH))
}

func (t *Termbox) CommandDraw() {
	t.Print(0, 14, color["default"], color["default"], "■注文")
	switch t.Selection.Y {
	case 15:
		t.Print(0, 15, color["default"], color["default"], "⇒")
	case 16:
		t.Print(0, 16, color["default"], color["default"], "⇒")
	}
	t.Print(2, 15, color["default"], color["default"], "買い")
	t.Print(2, 16, color["default"], color["default"], "売り")
	/* TODO: 実装
	t.Print(0, 17, color["default"], color["default"], "指値買")
	t.Print(0, 18, color["default"], color["default"], "指値売")
	t.Print(0, 19, color["default"], color["default"], "逆指値買")
	t.Print(0, 20, color["default"], color["default"], "逆指値売")
	*/
}

func (t *Termbox) Print(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		width := runewidth.RuneWidth(c)
		x = x + width
	}
}
