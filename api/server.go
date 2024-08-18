package main

import (
	"encoding/json"
	"fmt"
	holiday "go2/holiday"
	"log"
	"net/http"
	"time"
)

type Store struct {
	holidays holiday.HolidayList
}

var store = Store{}

func holidaysHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(store.holidays.GetHolidayItems())
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	var startDate = time.Date(
		2024, 11, 17, 0, 0, 0, 0, time.UTC)
	hol := holiday.Holiday{Destination: "Kos", Hotel: "Hotel California", Departure: "BMX", Duration: 7, Board: "AA", StartDate: startDate}
	hol2 := holiday.Holiday{Destination: "Crete", Hotel: "Palace Hotel", Departure: "MAN", Duration: 7, Board: "AA", StartDate: startDate}

	store.holidays.AddHoliday(hol, hol2)
	http.HandleFunc("/holidays", holidaysHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
