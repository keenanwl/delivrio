// usb - Self contained USB and HID library for Go
// Copyright 2019 The library Authors
//
// This library is free software: you can redistribute it and/or modify it under
// the terms of the GNU Lesser General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option) any
// later version.
//
// The library is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
// A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License along
// with the library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"context"
	"github.com/pteich/hid"
	"github.com/pteich/usbsymbolreader/scanner"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// create main context
	ctx, done := context.WithCancel(context.Background())

	// listen for system signals
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		select {
		case <-signalChannel:
			log.Print("shutdown signal received")
			done()
		}
	}()

	// main loop to enumerate USB devices and start reading from it
	for {

		log.Print("searching for USB devices...")
		devices := hid.Enumerate(0, 0)
		log.Printf("found %d devices", len(devices))

		for _, deviceInfo := range devices {
			log.Printf("found %s by %s . VendorID %d - ProductId %d", deviceInfo.Product, deviceInfo.Manufacturer, deviceInfo.VendorID, deviceInfo.ProductID)

			if deviceInfo.VendorID == 1504 || deviceInfo.ProductID == 4864 {

				symbolScanner, err := scanner.New(deviceInfo)
				if err != nil {
					log.Print(err)
					break
				}

				log.Printf("connected to %s (Serial No. %s)", deviceInfo.Product, deviceInfo.Serial)

				codes := symbolScanner.ReadCodes(ctx)
				for code := range codes {
					// TODO safe to file
					log.Printf("scanned code: %s", code.String())
				}

				break
			}
		}

		select {
		case <-ctx.Done():
			log.Print("shutting down")
			return
		case <-time.After(1 * time.Second):
		}

	}

}
