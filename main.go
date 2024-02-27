package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/emersion/go-vcard"
)

func main() {
	// Prompt user for input vCard file path
	fmt.Print("Enter the path to your vCard file: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	vcardFilePath := strings.TrimSpace(scanner.Text())

	// Prompt user for output CSV file path
	fmt.Print("Enter the path for the output CSV file: ")
	scanner.Scan()
	csvOutputFile := strings.TrimSpace(scanner.Text())

	// Open and parse the vCard file
	file, err := os.Open(vcardFilePath)
	if err != nil {
		fmt.Println("Error opening vCard file:", err)
		return
	}
	defer file.Close()

	// Create a vCard decoder
	decoder := vcard.NewDecoder(file)

	// Create a CSV file for writing
	csvFile, err := os.Create(csvOutputFile)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Name", "Email", "Phone"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing CSV header:", err)
		return
	}

	// Decode the vCard file
	for {
		card, err := decoder.Decode()
		if err != nil {
			// If there are no more cards, break out of the loop
			if err == io.EOF {
				break
			}
			fmt.Println("Error decoding vCard:", err)
			return
		}

		// Extract contact information from the vCard and write to CSV
		name := card.PreferredValue(vcard.FieldFormattedName)
		email := card.PreferredValue(vcard.FieldEmail)
		phone := card.PreferredValue(vcard.FieldTelephone)

		err = writer.Write([]string{name, email, phone})
		if err != nil {
			fmt.Println("Error writing contact to CSV:", err)
			return
		}
	}

	fmt.Println("Contacts exported to", csvOutputFile)
}
