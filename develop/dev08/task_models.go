package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/sys/windows"
)

type Pipe struct {
	Next *IOperation
}

type IOperation interface {
	Operation()
}

type Echo struct {
	text string
	Next *IOperation
}

func (e *Echo) Operation() {
	fmt.Println(e.text)
}

type PWD struct {
	Next *IOperation
}

func (pwd *PWD) Operation() {
	path, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(path)
}

type CD struct {
	path string
	Next *IOperation
}

func (cd *CD) Operation() {
	os.Chdir(cd.path)
}

type Kill struct {
	processName string
	pID         uint32
	Next        *IOperation
}

func (k *Kill) Operation() {
	errID := k.processID()
	if errID != nil {
		return
	}
	proc, err := os.FindProcess(int(k.pID))
	if err != nil {
		return
	}
	errKill := proc.Kill()
	if errKill != nil {
		return
	}
}

func (k *Kill) processID() error {
	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if e != nil {
		return e
	}
	p := windows.ProcessEntry32{Size: 568}
	for {
		e := windows.Process32Next(h, &p)
		if e != nil {
			return e
		}
		if windows.UTF16ToString(p.ExeFile[:]) == k.processName {
			k.pID = p.ProcessID
			return nil
		}
	}
}

type PS struct {
	Next *IOperation
}

func (ps *PS) Operation() {
	var processes bytes.Buffer
	h, e := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if e != nil {
		return
	}
	p := windows.ProcessEntry32{Size: 568}
	for {
		e := windows.Process32Next(h, &p)
		if e != nil {
			break
		}
		processes.WriteString(
			strconv.FormatUint(uint64(p.ProcessID), 10) +
				"  " +
				windows.UTF16ToString(p.ExeFile[:]) + "\n")
	}
	fmt.Print(processes.String())
}
