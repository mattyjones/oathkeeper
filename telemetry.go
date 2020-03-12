/*
Copyright © 2020 Matt Jones <Matt Jones <urlugal@gmail.com>>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"fmt"
	"time"
)

// incrementService will bump the count of services.
func (t *Telemetry) incrementService() {
	t.ServiceTotal++
}

// incrementFilesTotal will bump the count of hosts.
func (t *Telemetry) incrementHost() {
	t.HostTotal++
}

// initTelemetry will set the initial values
func (s *Collection) initTelemetry() {
	if s.Telemetry != nil {
		return
	}
	s.Telemetry = &Telemetry{
		StartedAt:    time.Now(),
		ServiceTotal: 0,
		HostTotal:    0,
	}
}

// printTelemtery will print a summary of the collected telemetry data
func printTelemetry(c *Collection) {

	fmt.Println("-------Telemetry-------")
	fmt.Println("Oathkeeper Version......:", AppVersion())
	fmt.Println("Total Services..........:", c.Telemetry.ServiceTotal)
	fmt.Println("Total Hosts.............:", c.Telemetry.HostTotal)
	fmt.Println("Elapsed Time............:", time.Since(c.Telemetry.StartedAt))
}
