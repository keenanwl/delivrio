package main

import (
	"context"
	"delivrio.io/print-client/ent"
	"delivrio.io/print-client/ent/localdevice"
	"fmt"
	"github.com/Ullaakut/nmap"
	"log"
	"net"
	"strings"
)

type RemoteConnectionData struct {
	WorkstationName string `json:"workstation_name"`
	URL             string `json:"url"`
	LastPing        string `json:"last_ping"`
}

type SubnetSearch struct {
	Results    []NetworkDevice `json:"results"`
	SearchTerm string          `json:"search_term"`
}

type NetworkDevice struct {
	AlreadyActive bool   `json:"already_active"`
	Hostname      string `json:"hostname"`
	Host          string `json:"host"`
	Port          string `json:"port"`
}

type RecentScan struct {
	Created string `json:"created"`
	Code    string `json:"code"`
	Result  string `json:"result"`
}

func (a *App) FindNetworkDevices(port string) ([]SubnetSearch, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, fmt.Errorf("error fetching interface addresses: %w", err)
	}

	output := make([]SubnetSearch, 0)
	for _, addr := range addrs {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			log.Println("Error parsing CIDR:", err)
			continue
		}
		if ip.To4() != nil && !ip.IsLoopback() {
			search := strings.Split(ip.String(), ".")
			pattern, results, err := searchSubnet(a.ctx, search, port)
			if err != nil {
				log.Println(err)
				continue
			}

			output = append(output, SubnetSearch{
				Results:    results,
				SearchTerm: fmt.Sprintf("%s:%s", pattern, port),
			})
		}
	}

	return output, nil
}

func networkPrinterActive(ctx context.Context, ip string, port string) (bool, error) {
	cli := ent.FromContext(ctx)

	return cli.LocalDevice.Query().
		Where(
			localdevice.Archived(false),
			localdevice.Address(fmt.Sprintf("%s:%s", ip, port)),
		).Exist(ctx)
}

func searchSubnet(ctx context.Context, subnet []string, searchPort string) (string, []NetworkDevice, error) {

	// Wildcard
	subnet[3] = "*"
	pattern := strings.Join(subnet, ".")

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(pattern),
		nmap.WithPorts(searchPort),
	)
	if err != nil {
		return "", nil, fmt.Errorf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if len(warnings) > 0 {
		return "", nil, fmt.Errorf("run finished with warnings: %s\n", warnings) // Warnings are non-critical errors from nmap.
	}
	if err != nil {
		return "", nil, fmt.Errorf("unable to run nmap scan: %v", err)
	}

	output := make([]NetworkDevice, 0)
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0], host.Hostnames, host.OS, host.Status)

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)

			if port.State.State == "open" {

				exists, err := networkPrinterActive(ctx, host.Addresses[0].String(), searchPort)
				if err != nil {
					log.Println("nmap: ", err)
					continue
				}

				output = append(output, NetworkDevice{
					Hostname:      host.Hostnames[0].String(),
					Host:          host.Addresses[0].String(),
					Port:          searchPort,
					AlreadyActive: exists,
				})
			}
		}
	}

	fmt.Printf("Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
	return pattern, output, nil
}
