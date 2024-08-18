package holiday

import (
	"fmt"
	"testing"
	"time"
)

var startDate = time.Date(
	2024, 11, 17, 0, 0, 0, 0, time.UTC)

var holiday = Holiday{Destination: "Kos", Hotel: "Hotel California", Departure: "BMX", Duration: 7, Board: "AA", StartDate: startDate, priceHistory: map[int]float64{0: 1000.00}}

func TestGetCurrentPrice(t *testing.T) {
	got := holiday.GetCurrentPrice()
	want := 1000.00

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAddCurrentPrice(t *testing.T) {
	holiday.AddCurrentPrice(2000.00)
	got := holiday.GetCurrentPrice()
	want := 2000.00

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	fmt.Println(holiday)
}
