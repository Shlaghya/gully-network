# Parking lot ticketing system 

## Run the program

#### The program can be run in 2 ways: Command line and input text file. A sample input text file is included. 
#### The commands included are : 

* `create_parking_lot <size>` is used to create a parking lot with a specific size.
* `park <registration number> <colour>` is used to park a car with a given registration number and colour. A parking slot will be alloted to the car..
* `leave <slot number>` frees the slot for another car.
* `status` prints the parking lot status as a list of cars with its slot numbers, registration numbers and colours.
* `slot_numbers_for_cars_with_colour <colour>` shows a list of slot numbers having car of that particular colour. 
* `registration_numbers_for_cars_with_colour <colour>` shows a list of registration numbers of cars with that particular colour.
* `slot_number_for_registration_number <registration number>` shows the slot number where the car of that partuclar registration number is parked.

#### Typing command `go run gully-network.go input.txt` will run the program with the input.txt as its input file. 
#### Typing command `go run gully-network.go` will run the program and open the terminal for command line inputs. The same commands shown above can be entered one at a time to view its results.
