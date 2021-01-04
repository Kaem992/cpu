package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu   string
	cpu2  string
	ready []string
	io1   []string
	io2   []string
	io3   []string
	io4   []string
)

func initialized() {
	cpu = ""
	cpu2 = ""
	ready = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)
}

func showProcess() {
	fmt.Printf("\n-----------\n")
	fmt.Printf("CPU 1 -> %s\n", cpu)
	fmt.Printf("CPU 2 -> %s\n", cpu2)
	fmt.Printf("Ready -> ")
	for i := range ready {
		fmt.Printf("%s ", ready[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 1 -> ")
	for i := range io1 {
		fmt.Printf("%s ", io1[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 2 -> ")
	for i := range io2 {
		fmt.Printf("%s ", io2[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 3 -> ")
	for i := range io3 {
		fmt.Printf("%s ", io3[i])
	}
	fmt.Printf("\n")
	fmt.Printf("I/O 4 -> ")
	for i := range io4 {
		fmt.Printf("%s ", io4[i])
	}
	fmt.Printf("\n\nCommand > ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, " \n")
	return data
}

func commandNew(p string) {
	if cpu == "" {
		cpu = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}

}
func commandTerminate() {

	if cpu != "" {
		cpu = deleteQueue(ready)
	} else if cpu2 != "" {
		cpu2 = deleteQueue(ready)
	}

}

func commandExpire() {
	p := deleteQueue(ready)
	if p == "" {
		return
	}
	insertQueue(ready, cpu)
	cpu = p
}

func commandIo1() {

	if cpu != "" {
		insertQueue(io1, cpu)
		cpu = ""
		commandExpire()
	} else if cpu2 != "" {
		insertQueue(io1, cpu2)
		cpu2 = ""
	}
}

func commandIo2() {
	if cpu != "" {
		insertQueue(io2, cpu)
		cpu = ""
		commandExpire()
	} else if cpu2 != "" {
		insertQueue(io2, cpu2)
		cpu2 = ""
	}
}

func commandIo3() {

	if cpu != "" {
		insertQueue(io3, cpu)
		cpu = ""
		commandExpire()
	} else if cpu2 != "" {
		insertQueue(io3, cpu2)
		cpu2 = ""
	}
}

func commandIo4() {

	if cpu != "" {
		insertQueue(io4, cpu)
		cpu = ""
		commandExpire()
	} else if cpu2 != "" {
		insertQueue(io4, cpu2)
		cpu2 = ""
	}
}

func commandIo1x() {
	p := deleteQueue(io1)
	if p == "" {
		return
	}
	if cpu == "" {
		cpu = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func commandIo2x() {
	p := deleteQueue(io2)
	if p == "" {
		return
	}
	if cpu == "" {
		cpu = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func commandIo3x() {
	p := deleteQueue(io3)
	if p == "" {
		return
	}
	if cpu == "" {
		cpu = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func commandIo4x() {
	p := deleteQueue(io4)
	if p == "" {
		return
	}
	if cpu == "" {
		cpu = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}

func insertQueue(q []string, data string) {
	for i := range q {
		if q[i] == "" {
			q[i] = data
			break
		}
	}
}

func deleteQueue(q []string) string {
	result := q[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
	}
	q[9] = ""
	return result
}

func main() {
	initialized()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "new":
			for i := range commandx {
				if i == 0 {
					continue
				}
				commandNew(commandx[i])
			}
		case "terminate":
			commandTerminate()
		case "expire":
			commandExpire()
		case "io1":
			commandIo1()
		case "io2":
			commandIo2()
		case "io3":
			commandIo3()
		case "io4":
			commandIo4()
		case "io1x":
			commandIo1x()
		case "io2x":
			commandIo2x()
		case "io3x":
			commandIo3x()
		case "io4x":
			commandIo4x()
		default:
			fmt.Printf("\nSorry !!! Command Error !!!\n")
		}
	}
}
