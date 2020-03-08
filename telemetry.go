package main

import (
	"fmt"
	"time"
)

// IncrementService will bump the count of services.
func (t *Telemetry) incrementService() {
	t.ServiceTotal++
}

// IncrementFilesTotal will bump the count of hosts.
func (t *Telemetry) incrementHost() {
	t.HostTotal++
}

// InitTelemetry will set the initial values
func (s *Destination) initTelemetry() {
	if s.Telemetry != nil {
		return
	}
	s.Telemetry = &Telemetry{
		StartedAt:    time.Now(),
		ServiceTotal: 0,
		HostTotal:    0,
	}
}

// Print a summary of telemetry data
func printTelemetry(c *Destination) {

	fmt.Println("-------Telemetry-------")
	fmt.Println("Oathkeeper Version......:", AppVersion())
	fmt.Println("Total Services..........:", c.Telemetry.ServiceTotal)
	fmt.Println("Total Hosts.............:", c.Telemetry.HostTotal)
	fmt.Println("Elapsed Time............:", time.Since(c.Telemetry.StartedAt))
}
