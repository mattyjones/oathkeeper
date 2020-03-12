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
	"bufio"
	"fmt"
	"os"
)

// appendServiceIfMissing will check the service slice before appending it
func appendServiceIfMissing(slice []*Service, s *Service) []*Service {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

// appendHostIfMissing will check the host slice before appending it
func appendHostIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

// outputCollection will output the data structures in the chosen format.
func (c *Collection) outputCollection() error {

	// create a yaml file and write it out
	_, yaml := find(c.Config.OutputType, "yaml")
	if yaml {
		c.createYaml()
	}

	// output the data in a simple text format to stdout
	_, stdout := find(c.Config.OutputType, "stdout")
	if stdout {
		c.writeStdout()
	}

	return nil
}

// find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// writeStdout will write the services to stdout for further user interactions
func (c *Collection) writeStdout() {
	for _, s := range c.Services {
		fmt.Println(s.Name)
		for _, h := range s.Endpoint.Host {
			fmt.Println(h)
		}
		fmt.Println("")
	}
}

// writeFile will write out a string to a given file
func (c *Collection) writeFile(out string) {
	f, _ := os.Create(c.Config.OutputFile)

	w := bufio.NewWriter(f)
	w.WriteString(out)

	w.Flush()
}
