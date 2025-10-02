//go:build windows

package printers

import (
	"delivrio.io/print-client/ent/printjob"
	"fmt"
	"github.com/alexbrainman/printer"
	"os"
	"os/exec"
	"strings"
)

func (c Client) ListPrinters() ([]Printer, error) {
	out, err := printer.ReadNames()
	if err != nil {
		return nil, err
	}

	printers := make([]Printer, 0)
	for _, line := range out {
		printers = append(
			printers,
			Printer{Name: strings.TrimSpace(line)},
		)
	}

	return printers, nil
}

func FireUSBPrintJob(printerSystemName string, printFile []byte, fileExtension printjob.FileExtension, useShell bool) error {

	if !useShell {
		return fireRawPrint(printerSystemName, printFile, fileExtension)
	}

	return fireShellPrint(printerSystemName, printFile, fileExtension)
}

func fireShellPrint(printerSystemName string, printFile []byte, fileExtension printjob.FileExtension) error {
	// Create a temporary file to save the printFile content
	tempFile, err := os.CreateTemp("", fmt.Sprintf("print_job_*.%s", fileExtension))
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer os.Remove(tempFile.Name()) // Ensure the temporary file is removed after use

	// Write the printFile content to the temporary file
	if _, err := tempFile.Write(printFile); err != nil {
		return fmt.Errorf("failed to write to temporary file: %w", err)
	}

	// Close the file to flush the content to disk
	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("failed to close temporary file: %w", err)
	}

	// Construct the PowerShell command to print the file
	// We use `Start-Process` with the `-Verb Print` to print the document
	// The `-NoNewWindow` flag ensures the process doesn't open a new window
	command := fmt.Sprintf(`Start-Process -FilePath "%s" -PrinterName "%s" -NoNewWindow -Verb Print`, tempFile.Name(), printerSystemName)

	// Execute the PowerShell command
	cmd := exec.Command("powershell", "-Command", command)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute PowerShell command: %w", err)
	}

	return nil
}

func fireRawPrint(printerSystemName string, printFile []byte, fileExtension printjob.FileExtension) error {

	p, err := printer.Open(printerSystemName)
	if err != nil {
		return fmt.Errorf("failed to open printer: %w", err)
	}
	defer p.Close()

	if fileExtension == printjob.FileExtensionPdf {
		// Add this in so we can test
		err = p.StartRawDocument("delivrio")
		if err != nil {
			return fmt.Errorf("failed to start document: %v", err)
		}
		defer p.EndDocument()

		// Start a new page
		err = p.StartPage()
		if err != nil {
			return fmt.Errorf("failed to start page: %v", err)
		}

		// Write ZPL content to the printer
		_, err = p.Write(printFile)
		if err != nil {
			return fmt.Errorf("failed to write to printer: %v", err)
		}

		// End the page
		err = p.EndPage()
		if err != nil {
			return fmt.Errorf("failed to end page: %v", err)
		}
	} else {
		// Don't use StartRawDoc as it doesn't work with ZPL printers?
		err = p.StartDocument("delivrio", "RAW")
		if err != nil {
			return fmt.Errorf("failed to start document: %v", err)
		}
		defer p.EndDocument()

		// Start a new page
		err = p.StartPage()
		if err != nil {
			return fmt.Errorf("failed to start page: %v", err)
		}

		// Write ZPL content to the printer
		_, err = p.Write(printFile)
		if err != nil {
			return fmt.Errorf("failed to write to printer: %v", err)
		}

		// End the page
		err = p.EndPage()
		if err != nil {
			return fmt.Errorf("failed to end page: %v", err)
		}
	}

	return nil
}
