//---------------------------------------------------------------------------------------
// Copyright (c) 2001-2021 by PDFTron Systems Inc. All Rights Reserved.
// Consult LICENSE.txt regarding license information.
//---------------------------------------------------------------------------------------

package main
import (
	"fmt"
	. "pdftron"
)

import  "pdftron/Samples/LicenseKey/GO"

//---------------------------------------------------------------------------------------
// The following sample illustrates how to use the PDF::Convert utility class to convert 
// documents and files to Excel.
//
// The Structured Output module is an optional PDFNet Add-on that can be used to convert PDF
// and other documents into Word, Excel, PowerPoint and HTML format.
//
// The PDFTron SDK Structured Output module can be downloaded from
// https://www.pdftron.com/documentation/core/info/modules/
//
// Please contact us if you have any questions.
//---------------------------------------------------------------------------------------

// Relative path to the folder containing the test files.
var inputPath = "../../TestFiles/"
var outputPath = "../../TestFiles/Output/"

//---------------------------------------------------------------------------------------

func catch(err *error) {
    if r := recover(); r != nil {
        *err = fmt.Errorf("%v", r)
    }
}

//---------------------------------------------------------------------------------------

func ConvertToExcelTest() (err error) {
	defer catch(&err)

	// Convert PDF document to Excel
	fmt.Println("Converting PDF to Excel")

	inputFile := inputPath + "paragraphs_and_tables.pdf"
	outputFile := outputPath + "paragraphs_and_tables.xlsx"

	// Convert to Excel
	ConvertToExcel(inputFile, outputFile)

	fmt.Println("Result saved in " + outputFile)
	return nil
}

//---------------------------------------------------------------------------------------

func ConvertToExcelWithOptionsTest() (err error) {
	defer catch(&err)

	// Convert PDF document to Excel with options
	fmt.Println("Converting PDF to Excel with options")

	inputFile := inputPath + "paragraphs_and_tables.pdf"
	outputFile := outputPath + "paragraphs_and_tables_second_page.xlsx"

	excelOutputOptions := NewExcelOutputOptions()

	// Convert only the second page
	excelOutputOptions.SetPages(2, 2)

	// Convert to Excel
	ConvertToExcel(inputFile, outputFile, excelOutputOptions)

	fmt.Println("Result saved in " + outputFile)
	return nil
}

//---------------------------------------------------------------------------------------

func main() {
    // The first step in every application using PDFNet is to initialize the 
    // library. The library is usually initialized only once, but calling 
    // Initialize() multiple times is also fine.
    PDFNetInitialize(PDFTronLicense.Key)

	//-----------------------------------------------------------------------------------

	PDFNetAddResourceSearchPath("../../../PDFNetC/Lib/")

	if !StructuredOutputModuleIsModuleAvailable() {
		fmt.Println("")
		fmt.Println("Unable to run the sample: PDFTron SDK Structured Output module not available.")
		fmt.Println("-----------------------------------------------------------------------------")
		fmt.Println("The Structured Output module is an optional add-on, available for download")
		fmt.Println("at https://www.pdftron.com/documentation/core/info/modules/. If you have already")
		fmt.Println("downloaded this module, ensure that the SDK is able to find the required files")
		fmt.Println("using the PDFNet::AddResourceSearchPath() function.")
		fmt.Println("")
		return
	}

	//-----------------------------------------------------------------------------------

	// Convert PDF document to Excel
	err := ConvertToExcelTest()
	if err != nil {
		fmt.Println(fmt.Errorf("Unable to convert PDF document to Excel, error: %s", err))
	}

	//-----------------------------------------------------------------------------------

	// Convert PDF document to Excel with options
	err = ConvertToExcelWithOptionsTest()
	if err != nil {
		fmt.Println(fmt.Errorf("Unable to convert PDF document to Excel, error: %s", err))
	}

	//-----------------------------------------------------------------------------------

    PDFNetTerminate()
    fmt.Println("Done.")
}