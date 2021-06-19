package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//array of cars
var lot []car

//the length of the parking lot
var lotSize int

//the number of cars actually parked
var numberOfCars int

//map that stores colour as key and array of number plates as its value.
var colourWithRegistration = make(map[string][]string)

//map that stores colour as key and array of slots as its value.
var colourWithSlot = make(map[string][]string)

type car struct {
	numberPlate string
	colour      string
}

//decodeComment decodes each line into specific commands.
func DecodeComment(s string) string {
	words := strings.Split(s, " ")
	//condition to check for a case insensitive first letter
	if strings.Contains(words[0], "create") || strings.Contains(words[0], "Create") {
		lotSize, _ = strconv.Atoi(words[1])
		lot = make([]car, lotSize)
		numberOfCars = 0
		return "Created a parking lot with " + words[1] + " slots"
	}
	if strings.Contains(words[0], "park") || strings.Contains(words[0], "Park") {
		if numberOfCars >= lotSize {
			return "Sorry, the parking lot is full"
		}
		for i := 0; i < lotSize; i++ {
			if (lot[i] == car{}) {
				//the number plates are converted to uppercase and stored to compensate for typos
				np := strings.ToUpper(words[1])
				//the colours are converted to lowercase and stored to compensate for typos
				c := strings.ToLower(words[2])
				lot[i] = car{
					numberPlate: np,
					colour:      c,
				}
				numberOfCars = numberOfCars + 1

				//adding the same entry to both the maps
				colourWithRegistration[c] = append(colourWithRegistration[c], np)
				colourWithSlot[c] = append(colourWithSlot[c], strconv.Itoa(i+1))
				return "Allocated slot number: " + strconv.Itoa(i+1)

			}
		}
	}
	if strings.Contains(words[0], "Leave") || strings.Contains(words[0], "leave") {
		slotNumber, _ := strconv.Atoi(words[1])
		if slotNumber > lotSize {
			return "Please specify lot number within the lot size limit"
		}
		if (lot[slotNumber-1] == car{}) {
			return "There is no car parked in that slot"
		}
		registrations := colourWithRegistration[lot[slotNumber-1].colour]

		//remove the car from both the maps
		for i, registration := range registrations {
			if registration == lot[slotNumber-1].numberPlate {
				registrations[i] = registrations[len(registrations)-1]
				colourWithRegistration[lot[slotNumber-1].colour] = registrations[:len(registrations)-1]
				break
			}
			if i == len(registrations)-1 {
				return "Wrong slot number"

			}
		}

		slots := colourWithSlot[lot[slotNumber-1].colour]
		for i, slot := range slots {
			if slot == words[1] {
				slots[i] = slots[len(slots)-1]
				colourWithSlot[lot[slotNumber-1].colour] = slots[:len(slots)-1]
				break
			}
		}

		lot[slotNumber-1] = car{}
		numberOfCars = numberOfCars - 1
		return "Slot number " + words[1] + " is free"

	}
	if strings.Contains(words[0], "status") || strings.Contains(words[0], "Status") {
		fmt.Printf("%-15s%-20s%s\n", "Slot Number", "Registration Number", "Colour")
		for i := 0; i < lotSize; i++ {
			if (lot[i] != car{}) {
				fmt.Printf("%-15s%-20s%s\n", strconv.Itoa(i+1), lot[i].numberPlate, lot[i].colour)
			}
		}
		return ""
	}
	if strings.Contains(words[0], "colour") || strings.Contains(words[0], "Colour") {
		//to search for registration numbers for a given colour
		if strings.Contains(words[0], "registration") || strings.Contains(words[0], "Registration") {
			v, ok := colourWithRegistration[strings.ToLower(words[1])]
			if ok && len(v) > 0 {
				return strings.Join(v, ",")
			} else {
				return "No car registrations found with colour " + words[1]
			}
			//to search for slots having car of a given colour
		} else {
			v, ok := colourWithSlot[strings.ToLower(words[1])]
			if ok && len(v) > 0 {
				return strings.Join(v, ",")
			} else {
				return "No slots found with car colour " + words[1]
			}
		}
	}
	//to search for the slot of a car with a given registration number.
	for i := 0; i < lotSize; i++ {
		if (lot[i] != car{}) {
			if lot[i].numberPlate == strings.ToUpper(words[1]) {
				return strconv.Itoa(i + 1)

			}
		}
	}
	return "car with registration " + words[1] + " not found"

}

func main() {
	//for interactive command prompt
	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			str := scanner.Text()
			if str == "exit" {
				break
			}
			result := DecodeComment(str)
			if result != "" {
				fmt.Println(result)
			}
		}
		//for file input
	} else {
		in := os.Args[1]
		file, err := os.Open(in)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			result := DecodeComment(scanner.Text())
			if result != "" {
				fmt.Println(result)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
