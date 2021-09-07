package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	if name == "" {
		name = "Christiane"
	}

	return fmt.Sprintf("Welcome to my party, %v!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and stands out his age.
func HappyBirthday(name string, age int) string {
	if name == "" {
		name = "Christiane"
	}

	if age == 0 {
		age = 58
	}

	return fmt.Sprintf("Happy birthday %v! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour string, direction string, distance float64) string {
	message := fmt.Sprintf("%v", Welcome(name))
	message += "\n" + fmt.Sprintf("You have been assigned to table %X. Your table is %v, exactly %.1f meters from here.", table, direction, distance)
	message += "\n" + fmt.Sprintf("You will be sitting next to %v.", neighbour)

	return message
}
