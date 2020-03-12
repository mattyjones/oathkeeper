# Oathkeeper

## Overview

Oathkeeper will scrape the AWS Endpoint and Services page and gather relative url's for scraping and parsing services. The 
default output is a formatted yaml file. The code will also allow output to stdout but that is not yet enabled. A cli will 
be the next feature added as time permits to allow users to configure various settings without having to recompile the code.

## Installation

1. Install Golang
2. Clone the project into your go path
3. `go build -o bin/oathkeeper` to build a binary for your architecture

## Usage

1. `oathkeeper`
