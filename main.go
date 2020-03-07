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
	"github.com/go-yaml/yaml"
	"os"
)

func main() {
	debug := false

	collection, _ := newCollection()

	collection.fetchServices()

	collection.fetchHosts()

	if debug { // TODO move this to a function for generating stdout text (writing it to a file as well)
		for _, s := range collection.Services {
			fmt.Println(s.Name)
			for _, h := range s.Endpoint.Host {
				fmt.Println(h)
			}
			fmt.Println("")
		}
	}

	foo, _ := yaml.Marshal(&collection)

	f, _ := os.Create("./aws_endpoints.yaml")

	w := bufio.NewWriter(f)
	n4, _ := w.WriteString(string(foo))
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
