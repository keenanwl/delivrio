package printers

import (
	"fmt"
	"net"
)

func FireNetworkPrintJob(networkAdr string, printData []byte) error {
	conn, err := net.Dial("tcp", networkAdr)
	if err != nil {
		return fmt.Errorf("network printer: fire job: %w", err)
	}
	defer conn.Close()

	_, err = conn.Write(printData)
	if err != nil {
		return fmt.Errorf("network printer: fire job: %w", err)
	}

	return nil
}
