package main

import (
	"bufio"
	"fmt"
	"os"
)

func manFlex() {
	fmt.Print("Welcome to the To-Do List App!\n1. Add Task\n2. View Tasks\n3. Remove Task\n4. Exit\nChoose an option (1-4):")
}

func delTask(slice []string) ([]string, bool) {
	var index int
	fmt.Println("Please enter which task you want to delete")
	fmt.Scan(&index)
	if index <= 0 || index > len(slice) {
		fmt.Println("Enter valid index please!")
		return slice, false
	}
	index--
	return append(slice[:index], slice[index+1:]...), true
}

func viewAll(slice []string) {
	if len(slice) == 0 {
		fmt.Println("No tasks available.")
	} else {
		fmt.Println("Here are all the tasks:\n========================")
		for i := range slice {
			fmt.Printf("%v-%s\n", i+1, slice[i])
		}
		fmt.Println("========================")
	}
}

func addTask(slice []string) []string {
	fmt.Println("Enter the task to add")
	fmt.Scanln()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()          // Waits for user input (including spaces)
	toAdd := scanner.Text() // Reads the whole line as a string
	return append(slice, toAdd)
}

func main() {
	var myList []string
	var key int
	var scs bool
	manFlex()
	fmt.Scan(&key)
	for key != 4 {
		if key < 1 || key > 4 {
			fmt.Println("Please enter a valide option")
		} else if key == 1 {
			myList = addTask(myList)
		} else if key == 2 {
			viewAll(myList)
		} else if key == 3 {
			myList, scs = delTask(myList)
			if scs {
				fmt.Println("Succesfully deleted!")
			} else {
				fmt.Println("ERR, can't delete the task.")
			}
		}
		manFlex()
		fmt.Scan(&key)
	}
	fmt.Print("Bye Bye!")
}
