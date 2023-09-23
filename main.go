package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/vincemis610/cli-golang/tasks"
)

/* Main function to execute program */
func main(){
	file, err := os.OpenFile("tasks.json", os.O_RDWR| os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		initProgram()
		return
	}

	switch os.Args[1] {
	case "list":
		task.ListTask(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("What's your task?")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		task.SaveOnFile(file, tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Id task must  be provided!")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("The id must be a number!")
			return
		}
		tasks = task.DeleteTask(tasks, id)
		task.SaveOnFile(file, tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Id task must  be provided!")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("The id must be a number!")
			return
		}
		tasks = task.CompleteTask(tasks, id)
		task.SaveOnFile(file, tasks)
	default:
		fmt.Println("The command not exists!")
		initProgram()
		return
	}
}

func initProgram() {
	fmt.Printf("%s\n", "What do you want to do? type the next commands: [list, add, delete, complete]")
	fmt.Printf("%s\n", "Usage: \n  go run main.go [command]\n")
}