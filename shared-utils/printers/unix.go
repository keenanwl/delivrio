//go:build linux

package printers

import (
	"delivrio.io/go/utils"
	"delivrio.io/print-client/ent/printjob"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func (c Client) ListPrinters() ([]Printer, error) {
	out, err := exec.Command("lpstat", "-e").Output()
	if err != nil {
		return nil, err
	}

	allLines := strings.TrimSpace(string(out))

	printers := make([]Printer, 0)

	for _, line := range strings.Split(allLines, "\n") {
		printers = append(
			printers,
			Printer{Name: strings.TrimSpace(line)},
		)
	}

	return printers, nil
}

// Printer must have Raw driver
// https://stackoverflow.com/questions/58185853/zebra-printer-prints-zpl-code-instead-of-label
func FireUSBPrintJob(printerSystemName string, printFile []byte, fileExtension printjob.FileExtension, useShell bool) error {

	switch fileExtension {
	case printjob.FileExtensionPdf:
		if !IsValidPDF(printFile) {
			return fmt.Errorf("invalid pdf file")
		}
	}

	fullPath, err := utils.MemoryDataToTmpFile(
		printFile,
		"delivrio-usb-print-job",
		fmt.Sprintf("label.%s", fileExtension),
	)
	if err != nil {
		return fmt.Errorf("running USB print: %w", err)
	}
	// Potentially move up?
	// Restart should remove any leftover directories in the case
	// that errors occur and
	defer os.RemoveAll(path.Dir(fullPath))

	if fileExtension == printjob.FileExtensionPdf {
		cmd := exec.Command("lp", "-d", printerSystemName, fullPath)
		log.Println("running USB print cmd: ", cmd.String())
		out, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("error running USB print: %w", err)
		}

		log.Println("output: ", string(out))

	} else {
		cmd := exec.Command("lp", "-d", printerSystemName, "-o", "raw", fullPath)
		log.Println("running USB print cmd: ", cmd.String())
		out, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("error running USB print: %w", err)
		}

		log.Println("output: ", string(out))

	}

	return nil
}
