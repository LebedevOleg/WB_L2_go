package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Message struct {
	text string
	err  error
}

func consoleInput(input chan Message) {
	reader := bufio.NewReader(os.Stdin)
	for {
		request, err := reader.ReadString('\n')
		if err != nil {
			input <- Message{request, errors.New("Connect Problem: " + err.Error())}
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

func Telnet(input string) {
	comandArr := strings.Split(input, " ")
	serverAddress := comandArr[len(comandArr)-2] + ":" + comandArr[len(comandArr)-1]
	fmt.Println("Create TCP Address")
	tcpAdress, err := net.ResolveTCPAddr("tcp", serverAddress)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Try connect to " + serverAddress)
	connect, err := net.DialTCP("tcp", nil, tcpAdress)
	if err != nil {
		fmt.Println("Connect FAIL;" + err.Error())
		return
	}
	defer connect.Close()

	fmt.Println("Connect success")
	connect.SetKeepAlive(true)
	connect.SetKeepAlivePeriod(5 * time.Second)
	//reader := bufio.NewReader(os.Stdin)
	inputCH := make(chan Message)
	outputCH := make(chan Message)
	breakCH := make(chan os.Signal, 1)
	signal.Notify(breakCH, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM,
		syscall.SIGINT)
	workCheck := true
	go consoleInput(inputCH)
	go connectMessage(outputCH, connect)
	for workCheck {
		select {
		case <-breakCH:
			fmt.Println("Closed Connection")
			workCheck = false
			continue
		case request := <-inputCH:
			if request.err != nil {
				fmt.Println(request.err)
				workCheck = false
			}
			fmt.Fprintf(connect, request.text)
			continue
		case reply := <-outputCH:
			if reply.err != nil {
				fmt.Println(reply.err)
				workCheck = false

			}
			fmt.Println(string(reply.text))
			continue
		}
	}

	/* request, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Connect Problem: " + err.Error())
		break
	}
	fmt.Fprintf(connect, request+"\n")
	reply, err := bufio.NewReader(connect).ReadString('\n')
	if err != nil {
		fmt.Println("Connect Problem: " + err.Error())
		break
	}
	fmt.Println(string(reply)) */

	/* message, _ := bufio.NewReader(connect).ReadString('\n') */
	/* reply := make([]byte, 1024)
	fmt.Println("Wait message from server...")
	connect.Read(reply)
	fmt.Println("Message catched:")
	fmt.Println(string(reply)) */
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	Telnet(input.Text())
}
