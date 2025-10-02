package main

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/mdns"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func fireNetworkPrintJob(networkAdr string, printData []byte) error {
	conn, err := net.Dial("ipp", networkAdr)
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

func main() {
	/*
		adr := "ipp://localhost/printers/Zebra_Technologies_ZTC_ZD421-203dpi_ZPL"

		zplLabel, err := ioutil.ReadFile("label.zpl")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = fireNetworkPrintJob(adr, zplLabel)
		if err != nil {
			panic(err)
		}*/

	// Replace this URL with the actual IPP endpoint of your printer
	printerURL := "ipp://localhost/printers/Zebra_Technologies_ZTC_ZD421-203dpi_ZPL"

	// Example ZPL code
	zplCode := "^XA^FO100,100^B3N,N,100,Y,N^FD>:123456^FS^FO200,200^GB100,100,100^FS^XZ"

	// Send a POST request with ZPL code to the printer
	response, err := http.Post(printerURL, "application/vnd.cups-raw", bytes.NewBufferString(zplCode))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error:", response.Status)
		return
	}

	// Read and print the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Body:", string(body))

	/*// Get a list of network interfaces on the machine
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Iterate through the network interfaces
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Iterate through the IP addresses associated with each interface
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			// Check if it's an IPv4 address
			if ipNet.IP.To4() != nil {
				// Get the subnet mask
				mask := ipNet.Mask

				// Print the interface name, IP address, and subnet mask
				fmt.Printf("Interface: %s, IP: %s, Subnet Mask: %s\n", iface.Name, ipNet.IP, mask)
			}
		}
	}

	/*ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()*/

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5-minute timeout.
	/*	scanner, err := nmap.NewScanner(
			//ctx,
			nmap.WithTargets("192.168.0.*"),
			nmap.WithPorts("9100"),
			//nmap.WithServiceInfo(),
			//nmap.WithAggressiveScan(),
		)
		if err != nil {
			log.Fatalf("unable to create nmap scanner: %v", err)
		}

		result, warnings, err := scanner.Run()
		if len(warnings) > 0 {
			log.Printf("run finished with warnings: %s\n", warnings) // Warnings are non-critical errors from nmap.
		}
		if err != nil {
			log.Fatalf("unable to run nmap scan: %v", err)
		}

		// Use the results to print an example output
		for _, host := range result.Hosts {
			if len(host.Ports) == 0 || len(host.Addresses) == 0 {
				continue
			}

			fmt.Printf("Host %q:\n", host.Addresses[0], host.Hostnames, host.OS, host.Status)

			for _, port := range host.Ports {
				fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
			}
		}

		fmt.Printf("Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)*/

	/*printerIP := "192.168.0.197" // Replace with your printer's IP address
	labelFile := "label.zpl"     // File containing ZPL label

	// Read the ZPL label content from the file
	zplLabel, err := ioutil.ReadFile(labelFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	conn, err := net.Dial("tcp", printerIP+":9100")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write(zplLabel)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Label sent to printer.")
	*/
}

func discoverPrinters() {

	// Start the query
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			// Print information about discovered printers
			fmt.Printf("Discovered Printer: %s (%s)\n", entry.Name, entry.AddrV4)
		}
	}()

	// Start the mDNS query
	err := mdns.Lookup("_ipp._tcp", entriesCh)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Run the query for a specified duration (e.g., 30 seconds)
	time.Sleep(5 * time.Second)
	close(entriesCh)
}
