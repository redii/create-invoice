package main

import (
	"log"
  "io/ioutil"
  "strings"
	"strconv"
	"github.com/jung-kurt/gofpdf"
)

const (
  devMode bool = true
)

func main() {
  log.Println("Generating Invoice...")

  // load input & template
  template := getData("template.txt")
  input := getData("input.txt")

  // check debug mode
  borderMode := "0"
  if (devMode == true) {
	  template = getData("template.txt.dev")
	  input = getData("input.txt.dev")
    borderMode = "1"
  }

  // create pdf object
  pdf := gofpdf.New("P", "mm", "A4", "")
  tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetLeftMargin(20)
  pdf.AddPage()

  // set header
  pdf.SetFont("Arial", "B", 20)
  pdf.CellFormat(170, 10, tr(template[0]), borderMode, 1, "L", false, 0, "")

  // set subheaders
  pdf.SetFont("Arial", "", 12)
  pdf.MultiCell(170, 6, tr(template[1]), borderMode, "L", false)
  pdf.Ln(6)

  // set address
	xBeforeAddress := pdf.GetX()
	yBeforeAddress := pdf.GetY()
  pdf.SetFont("Arial", "", 10)
  pdf.CellFormat(110, 5, tr(template[2]), borderMode, 0, "L", false, 0, "")

  // set contact
  pdf.MultiCell(60, 5, tr(template[3]), borderMode, "R", false)
  pdf.Ln(16)

	// set customer address (INPUT)
	pdf.SetFont("Arial", "", 10)
  xBeforeMeta := pdf.GetX()
	yBeforeMeta := pdf.GetY()
	pdf.SetXY(xBeforeAddress, yBeforeAddress + 10)
  pdf.MultiCell(60, 5, tr(input[0]), borderMode, "L", false)
	pdf.SetXY(xBeforeMeta, yBeforeMeta)

  // set invoice metadata (INPUT)
	pdf.SetFont("Arial", "B", 10)
	pdf.Line(xBeforeMeta, yBeforeMeta, xBeforeMeta + 170, yBeforeMeta)
	pdf.CellFormat(110, 10, tr(input[1]), borderMode, 0, "L", false, 0, "")
  pdf.CellFormat(60, 10, tr(input[2]), borderMode, 0, "R", false, 0, "")
  pdf.Ln(12)

  // set invoice text1
	pdf.SetFont("Arial", "", 10)
  pdf.MultiCell(170, 5, tr(template[4]), borderMode, "L", false)
  pdf.Ln(6)

	// set invoice table (INPUT)
	var total float64 = 0
	headerBool := true
	lines := strings.Split(string(input[3]), "\n")
	for _, line := range lines {
		items := strings.Split(string(line), ";")
		if (headerBool == true) {
	  	pdf.SetFont("Arial", "B", 12)
			pdf.CellFormat(90, 6, tr(items[0]), "1", 0, "L", false, 0, "")
			pdf.CellFormat(20, 6, tr(items[1]), "1", 0, "C", false, 0, "")
			pdf.CellFormat(20, 6, tr(items[2]), "1", 0, "C", false, 0, "")
			pdf.CellFormat(20, 6, tr(items[3]), "1", 0, "C", false, 0, "")
			pdf.CellFormat(20, 6, tr(items[4]), "1", 1, "C", false, 0, "")
			headerBool = false
		} else {
	  	pdf.SetFont("Arial", "", 10)
			pdf.CellFormat(90, 6, tr(items[0]), "1", 0, "L", false, 0, "")
			pdf.CellFormat(20, 6, tr(items[1]), "1", 0, "C", false, 0, "")
			pdf.CellFormat(20, 6, tr(items[2]), "1", 0, "L", false, 0, "")
			pdf.CellFormat(20, 6, tr(items[3] + " €"), "1", 0, "C", false, 0, "")
			amount, _ := strconv.ParseFloat(items[1], 64)
			cost, _ := strconv.ParseFloat(items[3], 64)
			tempTotalFloat := amount * cost
			total = total + tempTotalFloat
			tempTotalString := strconv.FormatFloat(tempTotalFloat, 'f', 2, 64)
			pdf.CellFormat(20, 6, tr(tempTotalString + " €"), "1", 1, "C", false, 0, "")
		}
	}
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(90, 6, tr(""), "1", 0, "L", false, 0, "")
	pdf.CellFormat(20, 6, tr(""), "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 6, tr(""), "1", 0, "L", false, 0, "")
	pdf.CellFormat(20, 6, tr(""), "1", 0, "C", false, 0, "")
	totalString := strconv.FormatFloat(total, 'f', 2, 64)
	pdf.CellFormat(20, 6, tr(totalString + " €"), "1", 1, "C", false, 0, "")
  pdf.Ln(6)

	// set invoice text2
	pdf.SetFont("Arial", "", 10)
  pdf.MultiCell(170, 5, tr(template[5]), borderMode, "L", false)
	pdf.Ln(6)

	// set footer data 1 & 2
  pdf.SetFont("Arial", "", 8)
	xBeforeFooter := pdf.GetX()
	yBeforeFooter := pdf.GetY()
  pdf.MultiCell(115, 5, tr(template[6]), borderMode, "L", false)
	pdf.SetXY(xBeforeFooter + 115, yBeforeFooter)
  pdf.MultiCell(55, 5, tr(template[7]), borderMode, "L", false)

  // output
	path := strings.Replace(tr(input[1]), " ", "_", -1) + ".pdf"
	log.Printf("Ouput: %v\n", path)
	pdf.OutputFileAndClose(path)
  log.Println("Invoice generated.")
}

func getData(path string) ([]string) {
	content, _ := ioutil.ReadFile(path)
	result := strings.Split(string(content), "\n;\n")
	return result
}