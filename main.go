package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to my Application")
	fmt.Println("***********************************************************")
	fmt.Println()

	var exit bool = true

	for exit {
		var choice = routeUser()

		switch choice {
		case 1:
			createNewPerson()
		case 2:
			getListPerson()
		case 3:
			searchPersonInfoById()
		case 4:
			editPersonInfoById()
		case 5:
			deletePersonInfoById()
		case 0:
			exit = false
		}
	}

	fmt.Println()
	fmt.Println("Exited Program.")
	fmt.Println("***********************************************************")
}

// route user to function
func routeUser() int32 {
	fmt.Println("1\t\t\tCreate new Person")
	fmt.Println("2\t\t\tGet Person's information list")
	fmt.Println("3\t\t\tSearch Person's information by id")
	fmt.Println("4\t\t\tEdit Person's information by id")
	fmt.Println("5\t\t\tDelete Person's information by id")
	fmt.Println("0\t\t\tExit program.")

	var choice int32
	fmt.Scanln(&choice)

	return choice
}

// Person define Person structure
type Person struct {
	id   int32
	name string
	age  int32
}

// define an array to store Person's info
var personInfoList = []Person{}

// mapping key:id and value:person
var personInfoMap = map[int32]Person{}

// Create new Person object
func createNewPerson() {
	fmt.Println("Enter your id: ")
	var id int32
	fmt.Scanln(&id)

	fmt.Println("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter your age: ")
	var age int32
	fmt.Scanln(&age)

	var person Person
	person.id = id
	person.name = name
	person.age = age

	personInfoList = append(personInfoList, person)

	personInfoMap[person.id] = person
}

// Get list persons
func getListPerson() {
	for i := 0; i < len(personInfoList); i++ {
		showPersonInfo(personInfoList[i])
	}
}

func showPersonInfo(person Person) {
	fmt.Println("Person's id: " + strconv.Itoa(int(person.id)))
	fmt.Println("Person's name: " + person.name)
	fmt.Println("Person's age: " + strconv.Itoa(int(person.age)))
}

// Get Person's info by id
func searchPersonInfoById() {
	fmt.Println("Enter person's id you want to search: ")
	var id int32
	fmt.Scanln(&id)

	showPersonInfo(personInfoMap[id])
}

// Edit Person's info
func editPersonInfoById() {
	fmt.Println("Enter person's id you want to edit: ")
	var id int32
	fmt.Scanln(&id)

	var person Person = personInfoMap[id]

	fmt.Println("Enter new name: ")
	var name string
	fmt.Scanln(&name)
	person.name = name

	fmt.Println("Enter new age: ")
	var age int32
	fmt.Scanln(&age)
	person.age = age

	for i := 0; i < len(personInfoList); i++ {
		if personInfoList[i].id == id {
			personInfoList[i] = person
		}
	}
}

// Delete Person's info object
func deletePersonInfoById() {
	fmt.Println("Enter person's id you want to delete: ")
	var id int32
	fmt.Scanln(&id)

	// delete in map
	delete(personInfoMap, id)

	// delete in array
	var indexOfPerson int32
	for i := 0; i < len(personInfoList); i++ {
		if personInfoList[i].id == id {
			indexOfPerson = int32(i)
			break
		}
	}

	personInfoList = append(
		personInfoList[:indexOfPerson],
		personInfoList[indexOfPerson+1:]...,
	)
}
