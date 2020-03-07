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
	"log"
	"strings"
	//"time"

	//"net"
	"net/http"
	"os"
	//"strings"
	"github.com/PuerkitoBio/goquery"
)

func (c *Destination) fetchHosts() {

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

		s.parseHosts(resp)
	}
}

// newHunt will initialize a hunt data structure
func newCollection() (*Destination, error) {
	var collection Destination

	return &collection, nil
}

// Finish will set the end time and do any necessary cleanup steps and then make the status as necessary
func (c *Destination) fetchServices() {

	collectionAddress := "https://docs.aws.amazon.com/general/latest/gr/aws-service-information.partial.html"

	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}

	// create a request
	req, err := http.NewRequest("GET", collectionAddress, nil)
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

	// scrape the tor page to check if the connection is being proxied
	c.parseServices(resp)
}

// Parse the tor project site to ensure that the proxy is working. This will return a bool and the ip address
func (c *Destination) parseServices(resp *http.Response) {

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("li").Each(func(i int, g *goquery.Selection) {

		// Get the service name in the list item
		svc := strings.TrimSpace(g.Text())

		// Get the relative link for the service page
		doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
			href, _ := item.Attr("href")
			if svc == item.Text() {

				// Create the link to the service page
				strParts := strings.Split(href, ".")
				linkHeader := "https://docs.aws.amazon.com/general/latest/gr"
				linkFooter := ".partial.html"
				link := linkHeader + strParts[1] + linkFooter

				// Create our basic service data structutre. This will be used for scraping the actual services
				var service Service
				service.Name = svc
				service.Link = link
				c.Services = append(c.Services, &service)
			}
		})
	})
}

func (s *Service) parseHosts(resp *http.Response) {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var h []string

	// Find the review items
	doc.Find("td").Each(func(i int, g *goquery.Selection) {

		if strings.Contains(g.Text(), ".com") {

			//h = []string

			h = append(h, strings.TrimSpace(g.Text()))

			//var h Host
			//h.Host = strings.TrimSpace(g.Text())
			//h.Port = "443"
			//
			//s.Hosts = append(s.Hosts, &h)
		}
	})
	s.Endpoint.Host = h
	s.Endpoint.Port = "443"
}
