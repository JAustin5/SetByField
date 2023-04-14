package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scanFile := bufio.NewScanner(file)
	act2 := new(Actors)

	// converting 2nd command line argument into an int to be used through the program
	secondField := os.Args[2]
	x, _ := strconv.ParseUint(secondField, 10, 32)
	sortField := int(x)

	for scanFile.Scan() {
		lines := scanFile.Text()

		parts := strings.Split(lines, ",")

		// separating the parts of the Actor class fields for Actors slice
		act := Actor{
			lastName:  parts[0],
			firstName: parts[1],
			series:    parts[2],
		}
		// loading file data into Actors
		act2.actor_slice = append(act2.actor_slice, act)
		// go to Sort interfaces method to do operation on Actors
		act2.sortingPos(sortField)
	}

	// looping through Actors in sorted slice to print contents via use of String() method
	for _, actings := range act2.actor_slice {
		fmt.Printf("%s", actings)
		fmt.Println()
	}
}

// Actor class
type Actor struct {
	lastName, firstName, series string
}

// Actors class
type Actors struct {
	actor_slice []Actor
}

// Sort interface containing one method
type Sort interface {
	sortingPos() []Actors
}

// Sort interface method taking 2nd command-line arg. as a parameter to sort Actors class
func (acts *Actors) sortingPos(num int) {
	if num == 0 {
		sort.SliceStable(acts.actor_slice, func(i, j int) bool {
			return acts.actor_slice[i].lastName < acts.actor_slice[j].lastName
		})
	}
	if num == 1 {
		sort.SliceStable(acts.actor_slice, func(i, j int) bool {
			return acts.actor_slice[i].firstName < acts.actor_slice[j].firstName
		})
	}
	if num == 2 {
		sort.SliceStable(acts.actor_slice, func(i, j int) bool {
			return acts.actor_slice[i].series < acts.actor_slice[j].series
		})
	}
}

// implemenation of String() method on Actor class to return preferred format
func (parts Actor) String() string {
	return fmt.Sprintf("%s, %s, %s", parts.lastName, parts.firstName, parts.series)
}
