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
	"time"
)

// The names of the struct's below are used to build the yaml output of a specific format.
// Changing the names without ensuring that the tags are correct will break the desired format

// TODO this should be a pointer
// Endpoints contains the required host data for services
type Endpoints struct {
	Host      []string `yaml:"host"`
	Port      string   `yaml:"port"`
	HostCount int      `yaml:"-"`
}

// Service contains the required data points for each service
type Service struct {
	Link     string    `yaml:"-"`
	Name     string    `yaml:"name"`
	Endpoint Endpoints `yaml:"endpoints"`
	Action   string    `yaml:"action"`
}

// Destination is a collection of services and will hold all known services and their associated endpoints.
type Destination struct {
	Services     []*Service `yaml:"destination"`
	ServiceCount int        `yaml:"-"`
	Telemetry    *Telemetry `yaml:"-"`
	Config       Config     `yaml:"-"`
}

// Telemetry hold various runtime statistics used for perf data
type Telemetry struct {
	StartedAt    time.Time `yaml:"-"`
	ServiceTotal int       `yaml:"-"`
	HostTotal    int       `yaml:"-"`
}

// TODO this should be a pointer
// Config stores various configuration details
type Config struct {
	OutputType []string `yaml:"-"`
	OutputFile string   `yaml:"-"`
	Debug      bool     `yaml:"-"`
}
