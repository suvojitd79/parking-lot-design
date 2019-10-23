package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"parking_lot/src/model"
	"strconv"
	"strings"
)

func main() {

	args := os.Args

	// read input from a file if possible
	if len(args) > 1 {
		processData("file", args[1]) // read from file
	}
	processData("cmd", "") // read from console
}

// provided file must be in bin directory
func processData(readFrom, fileName string) {

	var scanner *bufio.Scanner

	// read from file
	if readFrom == "file" {

		filePath, _ := os.Getwd()
		filePath = filePath[:len(filePath)-3] + "bin/" + fileName

		file, e := os.Open(filePath)

		if e != nil {
			log.Fatal(e) // error
			return
		}

		fmt.Println("================================")
		fmt.Println("READING DATA FROM THE PATH " + filePath)
		fmt.Println("================================")

		defer file.Close()

		// read from file
		scanner = bufio.NewScanner(file)

	} else if readFrom == "cmd" { // read from console

		//read from std in
		scanner = bufio.NewScanner(os.Stdin)

		// by default the data should be read from console

		fmt.Println("================================")
		fmt.Println("READING DATA FROM COMMAND PROMPT")
		fmt.Println("================================")

	}

	if scanner == nil {
		return // corner case - if no specific option is provided
	}

	var parkingLot *model.ParkingLot
	var msg string
	var err error

	for scanner.Scan() {
		inputLine := strings.Split(scanner.Text(), " ")
		operation := strings.TrimSpace(inputLine[0])

		switch operation {

		case "create_parking_lot":
			size, _ := strconv.ParseInt(strings.TrimSpace(inputLine[1]), 10, 64)
			parkingLot, msg = model.CreateParkingLot(uint64(size))
			fmt.Println(msg)
			break

		case "park":
			if parkingLot == nil {
				break
			}
			carID := strings.TrimSpace(inputLine[1])
			carColor := strings.TrimSpace(inputLine[2])
			msg, err = parkingLot.Park(model.Vehicle{RgNumber: carID, Col: carColor})
			if err != nil {
				fmt.Println(err.Error()) // slot not available
			} else {
				fmt.Println(msg) // slot available
			}
			break
		case "status":
			if parkingLot == nil {
				break
			}
			parkingLot.Status()
			break

		case "leave":
			if parkingLot == nil {
				break
			}
			ID, _ := strconv.ParseInt(strings.TrimSpace(inputLine[1]), 10, 64)
			msg, err = parkingLot.Leave(uint64(ID))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(msg)
			}
			break

		case "registration_numbers_for_cars_with_colour":
			if parkingLot == nil {
				break
			}
			color := strings.TrimSpace(inputLine[1])
			ans, _ := parkingLot.GetRegistrationNumbes(color)
			if len(ans) < 2 {
				fmt.Println(ans)
			} else {
				// format the output
				for i := 0; i < len(ans)-1; i++ {
					fmt.Print(ans[i] + ", ")
				}
				fmt.Println(ans[len(ans)-1])
			}

			break

		case "slot_numbers_for_cars_with_colour":
			if parkingLot == nil {
				break
			}
			color := strings.TrimSpace(inputLine[1])
			ans, _ := parkingLot.GetSlotNumbers(color)
			if len(ans) < 2 {
				fmt.Println(ans)
			} else {
				// format the output
				for i := 0; i < len(ans)-1; i++ {
					fmt.Print(strconv.FormatUint(ans[i], 10) + ", ")
				}
				fmt.Println(ans[len(ans)-1])
			}
			break

		case "slot_number_for_registration_number":
			if parkingLot == nil {
				break
			}
			rgNo := strings.TrimSpace(inputLine[1]) // registration number
			ans, err := parkingLot.GetSlotNumber(rgNo)
			if err != nil {
				fmt.Println(err.Error()) // not found
			} else {
				fmt.Println(ans)
			}
			break
		case "exit":
			os.Exit(0) //success
		default:
			break
		}

	}

	if er := scanner.Err(); er != nil {
		fmt.Println(er)
	}
}
