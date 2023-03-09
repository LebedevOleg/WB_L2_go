package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type Message struct {
	text string
	err  error
}

func consoleInput(input chan Message, breakCH chan os.Signal) {
	reader := bufio.NewReader(os.Stdin)

	for {
		request, _ := reader.ReadString('\n')
		if request == "\x04\r\n" {
			input <- Message{"", errors.New("program closed")}
			continue
		}
		input <- Message{request, nil}
	}

}
func connectMessage(output chan Message, con *net.TCPConn) {
	reader := bufio.NewReader(con)
	for {
		reply, err := reader.ReadString('\n')
		if err != nil {
			output <- Message{reply, errors.New("Connect Problem: " + err.Error())}
			continue
		}
		output <- Message{reply, nil}
	}

}

func Telnet(input string) error {
	comandArr := strings.Split(input, " ")
	timerCheck := false
	var timeout time.Duration
	for _, v := range comandArr {
		temp := strings.Split(v, "=")
		if temp[0] == "--timeout" {
			timerCheck = true
			sec, _ := strconv.Atoi(string([]rune(temp[1])[0]))
			timeout = time.Duration(sec * int(time.Second))
		}
	}
	serverAddress := comandArr[len(comandArr)-2] + ":" + comandArr[len(comandArr)-1]
	d := net.Dialer{}
	fmt.Println("Create TCP Address")
	if timerCheck {
		d = net.Dialer{Timeout: timeout}
	}
	fmt.Println("Try connect to " + serverAddress)
	connect, err := d.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Connect FAIL;" + err.Error())
		return err
	}
	defer func() { fmt.Println("Connection close"); connect.Close() }()

	fmt.Println("Connect success")
	d.KeepAlive = 5 * time.Second
	//reader := bufio.NewReader(os.Stdin)
	inputCH := make(chan Message)
	outputCH := make(chan Message)
	breakCH := make(chan os.Signal, 1)
	signal.Notify(breakCH, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM,
		syscall.SIGINT)
	workCheck := true
	go consoleInput(inputCH, breakCH)
	go connectMessage(outputCH, connect.(*net.TCPConn))
	for workCheck {
		select {
		case <-breakCH:
			fmt.Println("program closed")
			workCheck = false
			return nil
		case request := <-inputCH:
			if request.err != nil {
				fmt.Println(request.err)
				workCheck = false
				return request.err
			}
			fmt.Fprintf(connect, request.text)
			continue
		case reply := <-outputCH:
			if reply.err != nil {
				fmt.Println(reply.err)
				workCheck = false
				return reply.err
			}
			fmt.Println(string(reply.text))
			continue
		}
	}
	return nil

}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	Telnet(input.Text())
}
