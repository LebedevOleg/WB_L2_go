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

const fileName = "log.txt"

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	log := Logger{fileName: fileName}
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reqError := NewResponse("err", "Read bodyData error", 400)
		reqError.SendResponse(w, log)
		return
	}
	log.StartLog(r.RemoteAddr, r.RequestURI, string(reqJSON))
	var parsedEvent EventJSON
	err = json.Unmarshal(reqJSON, &parsedEvent)
	if err != nil {
		reqError := NewResponse("err", "Error json parsing", 400)
		reqError.SendResponse(w, log)
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
			res.SendResponse(w, log)
			return
		}
	}
	DB[len(DB)] = SimpleDBTable{Date: parsedEvent.Date,
		Events: []Event{Event{parsedEvent.Name, parsedEvent.Description}}}
	res := NewResponse("res", "Create event correct", 200)
	res.SendResponse(w, log)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	log := Logger{fileName: fileName}
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reqError := NewResponse("err", "Read bodyData error", 400)
		reqError.SendResponse(w, log)
		return
	}
	log.StartLog(r.RemoteAddr, r.RequestURI, string(reqJSON))
	var updateEvent UpdateEventJSON
	err = json.Unmarshal(reqJSON, &updateEvent)
	if err != nil {
		reqError := NewResponse("err", "Error json parsing", 400)
		reqError.SendResponse(w, log)
		return
	}
	for key, value := range DB {
		if value.Date == updateEvent.Previous.Date {
			for i, v := range value.Events {
				if v.Name == updateEvent.Previous.Name {
					value.Events = append(value.Events[:i], value.Events[i+1:]...)
					if len(value.Events) == 0 {
						delete(DB, key)
					}
					if value.Date == updateEvent.New.Date {
						value.Events = append(value.Events,
							Event{updateEvent.New.Name, updateEvent.New.Description})
						DB[key] = value
						res := NewResponse("res", "Update event correct", 200)
						res.SendResponse(w, log)
						return
					}
					for k, val := range DB {
						if val.Date == updateEvent.New.Date {
							val.Events = append(val.Events,
								Event{updateEvent.New.Name, updateEvent.New.Description})
							DB[k] = val
							res := NewResponse("res", "Update event correct", 200)
							res.SendResponse(w, log)
							return
						}
					}
					DB[len(DB)] = SimpleDBTable{Date: updateEvent.New.Date,
						Events: []Event{Event{updateEvent.New.Name, updateEvent.New.Description}}}
					res := NewResponse("res", "Update event correct", 200)
					res.SendResponse(w, log)
					return
				}
			}

			reqError := NewResponse("err", "Previous event not found", 503)
			reqError.SendResponse(w, log)
			return
		}
	}
	reqError := NewResponse("err", "Date with event not found", 400)
	reqError.SendResponse(w, log)
}
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	log := Logger{fileName: fileName}
	reqJSON, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reqError := NewResponse("err", "Read bodyData error", 400)
		reqError.SendResponse(w, log)
		return
	}
	log.StartLog(r.RemoteAddr, r.RequestURI, string(reqJSON))
	var parsedEvent EventJSON
	err = json.Unmarshal(reqJSON, &parsedEvent)
	if err != nil {
		reqError := NewResponse("err", "Error json parsing", 400)
		reqError.SendResponse(w, log)
		return
	}
	for key, value := range DB {
		if value.Date == parsedEvent.Date {
			for _, v := range value.Events {
				if v.Name == parsedEvent.Name {
					delete(DB, key)
					res := NewResponse("res", "Delete event correct", 200)
					res.SendResponse(w, log)
					return
				}
			}
			reqError := NewResponse("err", "Event dosn't exist", 503)
			reqError.SendResponse(w, log)
			return
		}
	}
	reqError := NewResponse("err", "Event dosn't exist", 503)
	reqError.SendResponse(w, log)

}
func EventForDay(w http.ResponseWriter, r *http.Request) {
	log := Logger{fileName: fileName}
	parseDate := r.URL.Query().Get("date")
	if parseDate == "" {
		reqError := NewResponse("err", "Read date error", 400)
		reqError.SendResponse(w, log)
		return
	}
	log.StartLog(r.RemoteAddr, r.RequestURI, "")
	serchedDate, err := time.Parse("2006-01-02", parseDate)
	if err != nil {
		reqError := NewResponse("err", "Parse date error", 400)
		reqError.SendResponse(w, log)
		return
	}
	responseJSON := make([]SimpleDBTable, 0, 1)
	for _, elem := range DB {
		if elem.Date == DateStruct(serchedDate) {
			responseJSON = append(responseJSON, elem)
			text, _ := json.Marshal(responseJSON)
			res := NewResponse("data", string(text), 200)
			res.SendResponse(w, log)
			return
		}
	}
	reqError := NewResponse("err", "Day without events", 503)
	reqError.SendResponse(w, log)
}
func EventForWeek(w http.ResponseWriter, r *http.Request) {
	log := Logger{fileName: fileName}
	parseDate := r.URL.Query().Get("date")
	if parseDate == "" {
		reqError := NewResponse("err", "Read date error", 400)
		reqError.SendResponse(w, log)
		return
	}
	log.StartLog(r.RemoteAddr, r.RequestURI, "")
	startDate, err := time.Parse("2006-01-02", parseDate)
	if err != nil {
		reqError := NewResponse("err", "Parse date error", 400)
		reqError.SendResponse(w, log)
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
	resData.SendResponse(w, log)
}
func EventForMonth(w http.ResponseWriter, r *http.Request) {
	log := Logger{fileName: fileName}
	parseDate := r.URL.Query().Get("date")
	if parseDate == "" {
		reqError := NewResponse("err", "Read date error", 400)
		reqError.SendResponse(w, log)
		return
	}
	log.StartLog(r.RemoteAddr, r.RequestURI, "")
	startDate, err := time.Parse("2006-01-02", parseDate)
	if err != nil {
		reqError := NewResponse("err", "Parse date error", 400)
		reqError.SendResponse(w, log)
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
	resData.SendResponse(w, log)
}

func main() {
	DB = make(map[int]SimpleDBTable)
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
