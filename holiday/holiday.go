package holiday

import "time"

type Holiday struct {
	Destination  string
	Hotel        string
	Departure    string
	Duration     int
	StartDate    time.Time
	Board        string
	priceHistory map[int]float64
}

func (h *Holiday) AddCurrentPrice(p float64) {
	h.priceHistory[len(h.priceHistory)] = p
}

func (h *Holiday) GetCurrentPrice() float64 {
	numPrices := len(h.priceHistory)
	return h.priceHistory[numPrices-1]
}

type HolidayList struct {
	Name  string
	items []Holiday
}

func (hl *HolidayList) AddHoliday(h ...Holiday) {
	hl.items = append(hl.items, h...)
}

func (hl *HolidayList) GetHolidayItems() []Holiday {
	return hl.items
}
