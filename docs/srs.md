# Oathkeeper


## 1. Introduction
This is a software requirements sheet for a pilot version of the tools. All asumptions and dependencies are for a deliverable pilot version only. The final production version will be developed seperately.

The keywords "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY",and "OPTIONAL" in this document are to be interpreted as described in RFC 2119.
These requirements are directed towards a functioning pilot program. The final production tool may have additional or modified requirements and may be documented and discussed separately.

### 1.1 Purpose
Oathkeeper is designed to scrape aws service documentation and create a yaml format document.

### 1.2 Intended Audience
- Public Proxy Team
- Infrastructure Security Assurance

### 1.3 Intended Use
The tool is designed to scrape sites for the information detailed below and organize it into a hierarchical yaml format for the purposes of being a whitelist.

### 1.4 Scope
A single cross-platform golang binary that will meet all functional and non-functional requirements for the pilot as written in this document. Documentation describing its use and design will also be made available as part of an initial deliverable.

### 1.5 Definitions and Acronyms
N/A

## 2. Overall Description
This is a tool to crawl and scrape dynamic aws endpoint service documentation sites, built as a statically linked binary to run on Linux. It will output data in a machine-readable and validated yaml format described below that can then be used by other tools. 

### 2.1 User Needs
- The ability to gather AWS service endpoints consisting of a host and port with minimal need for user intervention.
- The ability to produce usable data in a machine-readable yaml format.
  
### 2.2 Assumptions and Dependencies
A user will be running one of the following operating systems:
- OSX
- Amazon Linux
- ParrotOS
- CentOS 7
- The AWS service endpoint documentation url will be accessible to Oathkeeper
- This tool will be developed outside of Salesforce using open source tools and be released under an MIT license granting the user unrestricted freedom in its usage continued development

## 3. System Features and Requirements

### 3.1 Functional Requirements
- The tool must scrape the AWS service endpoint documentation site with the need for minimal user intervention once a job has been initiated.
- The tool must output machine-readable data derived from the parsing of the scraped data in this format.
- The tool must not store or leave any trace of the sites it discovers or scrapes.
- The tool must not disrupt or otherwise leave a trace of itself on the site it is scraping.

### 3.2 Nonfunctional Requirements
The tool will have full functionality on the following systems:
- OSX
- Amazon Linux
- ParrotOS
- CentOS 7
- The tool will be installed using standard packages or a compressed archive
- The tool will generate and output metrics related to the following measures:
- Time to parse the data set (parse the scraped data)
- Time to generate the data set (connect and scrape the site)
- Valid data points found services and endpoints cataloged)
- The lead developer will make available documentation describing the design and operation.
- The tool will ensure system integrity and stability in its use of system resources.
- The tool will generate usable debug info to help diagnose performance, usability, and functional issues.

### 3.3 External Interface Requirements
- An Amazon Linux, Parrot OS or CentOS 7 system
- A terminal program capable of 256 colors and displaying UTF-8 characters

