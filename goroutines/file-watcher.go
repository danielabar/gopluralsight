package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// folder to watch
const watchedPath = "./source"

// Watch for invoice files to be dropped into a folder.
// When an invoice comes in, print a message including invoice details.
func main() {
	// loop forever until code explicitly breaks out
	for {
		// get handle to watched folder
		d, _ := os.Open(watchedPath)

		// enumerate contents of folder, negative val means return as many files as are found
		// Readdir returns a slice of file info objects
		files, _ := d.Readdir(-1)

		// loop through files found in watch dir
		for _, fi := range files {
			// build up path to file using string concatenation
			filePath := watchedPath + "/" + fi.Name()

			// open file (note open function works for both files and folders)
			f, _ := os.Open(filePath)

			// read all contents of file and tore in a variable
			data, _ := ioutil.ReadAll(f)

			// done with file so close it, not using Defer because that waits until function exits,
			// but this code runs inside main function therefore won't exit until program exits,
			// therefore Defer on file close would cause program to hang on to all the file handles unnecessarily
			f.Close()

			// delete file, in a real production system, would want to wait until confirmation that file
			// was processed correctly, but that requires tracking file names, don't want to overcomplicate
			// this demo
			os.Remove(filePath)

			// asynchronously process file in goroutine with self invoking anonymous function
			// pass in data parameter to keep this isolated from data changes in outer scope
			go func(data string) {
				// get a reader from csv package, note that it does not work with a raw string,
				// it needs a reader object. We don't want to pass in file handle from outer scope,
				// to keep this isolated, therefore use strings package to convert string to reader
				reader := csv.NewReader(strings.NewReader(data))
				// ReadAll returns a slice of all the records in the file
				records, _ := reader.ReadAll()
				// iterate over each recod
				for _, r := range records {
					// create a new invoice object
					invoice := new(Invoice)
					// there is no schema information in csv file, therefore can't use unmarshall
					// like we did with xml in web services example, therefore must parse it out manually
					// to populate invoice
					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoice.PurchaseOrderNumber, _ = strconv.Atoi(r[2])

					// for invoice date, first convert string in csv file to int,
					// use parseInt for more control than Atoi, give it:
					// string to be converted, base, and width of number represented...
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					// ...then int (representing unix timestamp) to time object
					invoice.InvoiceDate = time.Unix(unixTime, 0)

					// print message confirming all data received as expected
					fmt.Printf("Received invoice '%v' for $%.2f and submitted", invoice.Number, invoice.Amount)
				} //for reach record
			}(string(data))
		} //for each file in watch dir
	} //loop forever
} //main

// data type to hold invoice information, note invoice number is a string because it contains dashes
type Invoice struct {
	Number              string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}
