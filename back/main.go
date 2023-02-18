package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AddTimePayload struct {
	Years   int `json:"years,omitempty"`
	Months  int `json:"months,omitempty"`
	Days    int `json:"days,omitempty"`
	Hours   int `json:"hours,omitempty"`
	Minutes int `json:"minutes,omitempty"`
}

// accept arg period and possible values hour and second
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	res := r.URL.Query()
	nw := time.Now()
	var toWrite string
	if val := res.Get("period"); val != "" {
		switch val {
		case "hour":
			toWrite = "HOUR: " + strconv.Itoa(nw.Hour())
		case "minute":
			toWrite = "MINUTE: " + strconv.Itoa(nw.Minute())
		default:
			toWrite = nw.Format("01/02/2006 15:04")
		}
	} else {
		toWrite = nw.Format("01/02/2006 15:04")
	}

	fmt.Fprint(w, toWrite)
}

// Accepts JSON via POST method
// format: {years: int, months: int, days: int, hours: int, minutes: int}
func AddTimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Server Error: %s", err)
			return
		}
		var timeToAdd AddTimePayload
		log.Println(string(data))
		err = json.Unmarshal(data, &timeToAdd)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Internal Server Error %s", err.Error())
			return
		}
		nw := time.Now()
		nw = nw.Add(time.Minute * time.Duration(timeToAdd.Years) * 525600)
		nw = nw.Add(time.Minute * time.Duration(timeToAdd.Months) * 43800)
		nw = nw.Add(time.Minute * time.Duration(timeToAdd.Days) * 24 * 60)
		nw = nw.Add(time.Minute * time.Duration(timeToAdd.Hours) * 60)
		nw = nw.Add(time.Minute * time.Duration(timeToAdd.Minutes))

		fmt.Fprintf(w, "In %d years and %d months and %d weeks and %d hours and %d minutes the time is: %s",
			timeToAdd.Years, timeToAdd.Months, timeToAdd.Days, timeToAdd.Hours, timeToAdd.Minutes, nw.Format("01/02/2006 15:04"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Use json ")
		return
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/add_time", AddTimeHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
