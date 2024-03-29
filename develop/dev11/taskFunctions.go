package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type SimpleDBTable struct {
	Date   DateStruct `json:"date"`
	Events []Event    `json:"events"`
}

type UpdateEventJSON struct {
	Previous EventJSON
	New      EventJSON
}

type Event struct {
	Name        string
	Description string
}

type EventJSON struct {
	Date        DateStruct `json:"date"`
	Name        string     `json:"name"`
	Description string     `json:"desc"`
}

type IResponse interface {
	SendResponse(http.ResponseWriter, Logger)
}

type Error struct {
	Err     string `json:"error"`
	errCode int
}

func (e *Error) SendResponse(w http.ResponseWriter, log Logger) {
	w.WriteHeader(e.errCode)
	errText, _ := json.Marshal(e)
	fmt.Fprint(w, string(errText))
	log.EndLog(string(errText))
}

func NewResponse(value, text string, code int) IResponse {
	switch value {
	case "res":
		return &TextResult{text, code}
	case "err":
		return &Error{text, code}
	case "data":
		var SD []SimpleDBTable
		json.Unmarshal([]byte(text), &SD)
		return &DataResult{SD, code}
	default:
		return nil
	}
}

type TextResult struct {
	Res     string `json:"result"`
	resCode int
}

func (r *TextResult) SendResponse(w http.ResponseWriter, log Logger) {
	w.WriteHeader(r.resCode)
	resText, _ := json.Marshal(r)
	fmt.Fprint(w, string(resText))
	log.EndLog(string(resText))
}

type DataResult struct {
	Res     []SimpleDBTable `json:"result"`
	resCode int
}

func (r *DataResult) SendResponse(w http.ResponseWriter, log Logger) {
	w.WriteHeader(r.resCode)
	resText, _ := json.Marshal(r)
	fmt.Fprint(w, string(resText))
	log.EndLog(string(resText))
}

type DateStruct time.Time

func (d *DateStruct) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}
	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*d = DateStruct(t) //set result using the pointer
	return nil
}
func (d DateStruct) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format("2006-01-02") + `"`), nil
}
func (d1 DateStruct) Equals(d2 DateStruct) {}

type Logger struct {
	file     *os.File
	fileName string
	textLog  string
}

func (l *Logger) StartLog(addres, URI, body string) {
	l.textLog = time.Now().String() + "::" + addres + "::" + URI + ";Body:" + body + "\n"
}

func (l *Logger) EndLog(res string) {
	l.textLog += "\tresponce:" + res + "/"
	l.file, _ = os.OpenFile(l.fileName, os.O_APPEND|os.O_WRONLY, 0600)
	defer l.file.Close()
	l.file.WriteString(l.textLog + "\n")
}
