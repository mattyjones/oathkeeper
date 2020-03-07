package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func (c *Destination) writeYaml() {

	// FIXME Better file names
	foo, _ := yaml.Marshal(&c)

	// FIXME We should alpha sort the yaml file to make searching easier
	writeOutput(string(foo))

}

func writeOutput(out string) {
	f, _ := os.Create("./aws_endpoints.yaml")

	w := bufio.NewWriter(f)
	n4, _ := w.WriteString(out)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
