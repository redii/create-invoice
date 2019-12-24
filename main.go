package main

import (
	"log"
	"github.com/desertbit/fillpdf"
)

func main() {
	// Create the form values.
	form := fillpdf.Form{
		// "field_1": "Hello",
		// "field_2": "World",
		"Anschrift": "Region Hannover\nTeam 51.18\nHildesheimer Str. 20\n30199 Hannover",
		"Zeichen": "Rechnung 123456\nTeambesprechung",
		"Datum": "Garbsen, den 01.01.2020",
	}

	// Fill the form PDF with our values.
	err := fillpdf.Fill(form, "invoice.pdf", "./filled.pdf", true)
	if err != nil {
		log.Fatal(err)
	}
}