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
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strings"
)

// fetchHosts will scrape the AWS documentation page of each service and pass the response to parseHosts to
// parse and assemble the data into a usable structure.
func (c *Destination) fetchHosts() error {

	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}

	for _, s := range c.Services {
		// create a request
		req, err := http.NewRequest("GET", s.Link, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, "can't create request:", err)
			os.Exit(2)
		}
		// use the http client to fetch the page
		resp, err := httpClient.Do(req)
		if err != nil {
			fmt.Fprintln(os.Stderr, "can't GET page:", err)
			os.Exit(3)
		}
		defer resp.Body.Close()

		// parse the response into a usable data structure
		s.parseHosts(resp)

		// Add the hosts to the running total
		for s.Endpoint.HostCount > 0 {
			c.Telemetry.IncrementHost()
			s.Endpoint.HostCount--
		}
	}
	return nil
}

// NewDestination will initialize a Destination data structure
func newDestination() (*Destination, error) {
	var destination Destination

	destination.start()

	return &destination, nil
}

// Initialize the tool for gathering telemetry
func (s *Destination) start() error {
	s.initTelemetry()

	return nil
}

// finish will cleanup and close anything remaining open
func (c *Destination) finish() error {

	return nil

}

// fetchServices will scrape the AWS documentation page and pass the response to parseServices to
// parse and assemble the data into a usable structure for discovering and cataloging endpoints.
func (c *Destination) fetchServices() error {

	documentationAddress := "https://docs.aws.amazon.com/general/latest/gr/aws-service-information.partial.html"

	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}

	// create a request
	req, err := http.NewRequest("GET", documentationAddress, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create request:", err)
		os.Exit(2)
	}
	// use the http client to fetch the page
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't GET page:", err)
		os.Exit(3)
	}
	defer resp.Body.Close()

	// parse the response into a usable data structure
	c.parseServices(resp)

	return nil
}

// Parse the the endpoint and quota documentation page for a service for specific data points and then assemble
// them into a usable data structure.
func (c *Destination) parseServices(resp *http.Response) error {

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the li items. The services are organized into an unordered list with their respective relative
	// links to their detailed documentation pages.
	doc.Find("li").Each(func(i int, g *goquery.Selection) {

		// Get the service name in the list item and remove any excess white space
		// BUG TrimSpace does not always work
		svc := strings.TrimSpace(g.Text())

		// Get the relative link for the service page
		doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
			href, _ := item.Attr("href")
			if svc == item.Text() {

				c.Telemetry.IncrementService()

				// Create the link to the service page
				strParts := strings.Split(href, ".")
				linkHeader := "https://docs.aws.amazon.com/general/latest/gr"
				linkFooter := ".partial.html"
				link := linkHeader + strParts[1] + linkFooter

				// Create our basic service data structure. This will be used for scraping the actual services
				var service Service
				service.Name = svc
				service.Link = link
				c.Services = appendServiceIfMissing(c.Services, &service)
			}
		})
	})

	return nil
}

// Parse the service documentation page for specific data points and then assemble them into a
// usable data structure.
func (s *Service) parseHosts(resp *http.Response) error {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// slice of all hosts
	var h []string

	// Count of all the hosts
	var hCount int

	// Find the td items. The data points we need are stored in a table. We pull each column in the table row
	// and look for specific matches to the pattern we need.
	doc.Find("td").Each(func(i int, g *goquery.Selection) {

		if strings.Contains(g.Text(), ".com") {

			// BUG TrimSpace does not always work
			h = appendHostIfMissing(h, strings.TrimSpace(g.Text()))
			hCount++

		}
	})
	s.Endpoint.HostCount = hCount
	s.Endpoint.Host = h

	// BUG not every endpoint is HTTPS, we need to figure out how to get the protocol for each url. In the
	// BUG cases where there is no protocol we need to find some other way or a method to call it out for
	// BUG manual checking. This could only be done on the first pass but that would require loading an existing
	// BUG endpoint yaml file and then diffing and combining it with what we have which is beyond the scope
	// BUG of the pilot.
	s.Endpoint.Port = "443"

	return nil
}
