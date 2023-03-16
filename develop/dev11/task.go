package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var DB map[int]SimpleDBTable
var log Logger //filename

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reqError := NewResponse("err", "Read bodyData error", 400)
		reqError.SendResponse(w)
		return
	}
	var parsedEvent EventJSON
	err = json.Unmarshal(reqJSON, &parsedEvent)
	if err != nil {
		reqError := NewResponse("err", "Error json parsing", 400)
		reqError.SendResponse(w)
		return
	}
	for idx, elem := range DB {
		if elem.Date == parsedEvent.Date {
			for _, v := range elem.Events {
				if v.Name == parsedEvent.Name {
					reqError := NewResponse("err", "Event with this name alredy exist", 400)
					reqError.SendResponse(w, log)
					return
				}
			}
			elem.Events = append(DB[idx].Events, Event{parsedEvent.Name, parsedEvent.Description})
			DB[idx] = elem
			res := NewResponse("res", "Create event correct", 200)
			res.SendResponse(w)
			return
		}
	}
	DB[len(DB)] = SimpleDBTable{Date: parsedEvent.Date,
		Events: []Event{Event{parsedEvent.Name, parsedEvent.Description}}}
	res := NewResponse("res", "Create event correct", 200)
	res.SendResponse(w)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reqError := NewResponse("err", "Read bodyData error", 400)
		reqError.SendResponse(w)
		return
	}
	var updateEvent UpdateEventJSON
	err = json.Unmarshal(reqJSON, &updateEvent)
	if err != nil {
		reqError := NewResponse("err", "Error json parsing", 400)
		reqError.SendResponse(w)
		return
	}
	for key, value := range DB {
		if value.Date == updateEvent.Previous.Date {
			for i, v := range value.Events {
				if v.Name == updateEvent.Previous.Name {
					DB[key].Events[i] = Event{updateEvent.New.Name, updateEvent.New.Description}
					res := NewResponse("res", "Update event correct", 200)
					res.SendResponse(w)
					return
				}
			}
			reqError := NewResponse("err", "Previous event not found", 503)
			reqError.SendResponse(w)
			return
		}
	}
	reqError := NewResponse("err", "Date with event not found", 400)
	reqError.SendResponse(w)
}
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reqError := NewResponse("err", "Read bodyData error", 400)
		reqError.SendResponse(w)
		return
	}
	var parsedEvent EventJSON
	err = json.Unmarshal(reqJSON, &parsedEvent)
	if err != nil {
		reqError := NewResponse("err", "Error json parsing", 400)
		reqError.SendResponse(w)
		return
	}
	for key, value := range DB {
		if value.Date == parsedEvent.Date {
			for _, v := range value.Events {
				if v.Name == parsedEvent.Name {
					delete(DB, key)
					res := NewResponse("res", "Delete event correct", 200)
					res.SendResponse(w)
					return
				}
			}
			reqError := NewResponse("err", "Event dosn't exist", 503)
			reqError.SendResponse(w)
			return
		}
	}
	reqError := NewResponse("err", "Event dosn't exist", 503)
	reqError.SendResponse(w)

}
func EventForDay(w http.ResponseWriter, r *http.Request) {
	parseDate := r.URL.Query().Get("date")
	if parseDate == "" {
		reqError := NewResponse("err", "Read date error", 400)
		reqError.SendResponse(w)
		return
	}
	serchedDate, err := time.Parse("2006-01-02", parseDate)
	if err != nil {
		reqError := NewResponse("err", "Parse date error", 400)
		reqError.SendResponse(w)
		return
	}
	responseJSON := make([]SimpleDBTable, 0, 1)
	for _, elem := range DB {
		if elem.Date == DateStruct(serchedDate) {
			responseJSON = append(responseJSON, elem)
			text, _ := json.Marshal(responseJSON)
			res := NewResponse("data", string(text), 200)
			res.SendResponse(w)
			return
		}
	}
	reqError := NewResponse("err", "Day without events", 503)
	reqError.SendResponse(w)
}
func EventForWeek(w http.ResponseWriter, r *http.Request) {
	parseDate := r.URL.Query().Get("date")
	if parseDate == "" {
		reqError := NewResponse("err", "Read date error", 400)
		reqError.SendResponse(w)
		return
	}
	startDate, err := time.Parse("2006-01-02", parseDate)
	if err != nil {
		reqError := NewResponse("err", "Parse date error", 400)
		reqError.SendResponse(w)
		return
	}
	endDate := startDate.Add(24 * 7 * time.Hour)
	//var responseString bytes.Buffer
	responseArr := make([]SimpleDBTable, 0, len(DB))
	for _, value := range DB {
		if time.Time(value.Date).After(startDate) && time.Time(value.Date).Before(endDate) {
			responseArr = append(responseArr, value)
		}
	}
	text, _ := json.Marshal(responseArr)
	resData := NewResponse("data", string(text), 200)
	resData.SendResponse(w)
}
func EventForMonth(w http.ResponseWriter, r *http.Request) {
	parseDate := r.URL.Query().Get("date")
	if parseDate == "" {
		reqError := NewResponse("err", "Read date error", 400)
		reqError.SendResponse(w)
		return
	}
	startDate, err := time.Parse("2006-01-02", parseDate)
	if err != nil {
		reqError := NewResponse("err", "Parse date error", 400)
		reqError.SendResponse(w)
		return
	}
	endDate := startDate.Add(24 * 30 * time.Hour)
	//var responseString bytes.Buffer
	responseArr := make([]SimpleDBTable, 0, len(DB))
	for _, value := range DB {
		if time.Time(value.Date).After(startDate) && time.Time(value.Date).Before(endDate) {
			responseArr = append(responseArr, value)
		}
	}
	text, _ := json.Marshal(responseArr)
	resData := NewResponse("data", string(text), 200)
	resData.SendResponse(w)
}

func main() {
	DB = make(map[int]SimpleDBTable)
	log = Logger{fileName: "log.txt"}
	http.HandleFunc("/create_event", CreateEvent)
	http.HandleFunc("/update_event", UpdateEvent)
	http.HandleFunc("/delete_event", DeleteEvent)
	http.HandleFunc("/event_for_day", EventForDay)
	http.HandleFunc("/event_for_month", EventForMonth)
	http.HandleFunc("/event_for_week", EventForWeek)
	err := http.ListenAndServe(":6666", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
