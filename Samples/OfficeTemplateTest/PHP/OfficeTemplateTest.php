<?php
//------------------------------------------------------------------------------
// Copyright (c) 2001-2021 by PDFTron Systems Inc. All Rights Reserved.
// Consult legal.txt regarding legal and license information.
//------------------------------------------------------------------------------
namespace Pdftron;
include("../../../PDFNetC/Lib/PDFNetPHP.php");
include("../../LicenseKey/PHP/LicenseKey.php");

// Relative path to the folder containing the test files.
$input_path = getcwd()."/../../TestFiles/";
$output_path = $input_path."Output/";

//------------------------------------------------------------------------------
// The following sample illustrates how to use the PDF::Convert utility class
// to convert MS Office files to PDF and replace templated tags present in the document
// with content supplied via json
//
// For a detailed specification of the template format and supported features,
// see: https://www.pdftron.com/documentation/core/guides/generate-via-template/data-model/
//
// This conversion is performed entirely within the PDFNet and has *no*
// external or system dependencies -- Conversion results will be
// the same whether on Windows, Linux or Android.
//
// Please contact us if you have any questions.
//------------------------------------------------------------------------------

function main()
{
	// The first step in every application using PDFNet is to initialize the
	// library. The library is usually initialized only once, but calling
	// Initialize() multiple times is also fine.
	global $LicenseKey;
	PDFNet::Initialize($LicenseKey);
	PDFNet::SetResourcesPath("../../../Resources");

	global $input_path, $output_path;
	$input_filename = 'SYH_Letter.docx';
	$output_filename = 'SYH_Letter.pdf';

	$json = '
	{
		"dest_given_name": "Janice N.",
        "dest_street_address": "187 Duizelstraat",
        "dest_surname": "Symonds",
        "dest_title": "Ms.",
        "land_location": "225 Parc St., Rochelle, QC ",
        "lease_problem": "According to the city records, the lease was initiated in September 2010 and never terminated",
        "logo": { "image_url": "' . $input_path . 'logo_red.png", "width" : 64, "height":  64 },
        "sender_name": "Arnold Smith"
	}';

	// Start with a PDFDoc (the conversion destination)
	$pdfdoc = new PDFDoc();

	$options = new OfficeToPDFOptions();
	$options->SetTemplateParamsJson($json);

    // perform the conversion with template delimiters and content dictionary
	Convert::OfficeToPDF($pdfdoc, $input_path.$input_filename, $options);

	// save the result
	$pdfdoc->Save($output_path.$output_filename, SDFDoc::e_linearized, NULL);

	// And we're done!
	echo nl2br("Saved ".$output_filename . "\n");
}

main()

?>