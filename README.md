# getmacaddress

## What is getmacaddress? 

getmacaddress is a utility by entering given MAC address user can retrieve OUI vendor information, detect virtual machines, possible applications, read the information encoded in the MAC, and get our research's results regarding the MAC address or the OUI.

## Deployment

### Clone this repo

git clone git@github.com:gaurav1der/getmacaddress.git

### Create Docker image

Make sure docker is install and running. To install docker refer to https://hub.docker.com

docker build .

### Run Docker container with interactive mode

docker run -it {image ID} bash

### Once inside the container to get MAC address related details follow below steps

1) export API_KEY={YOUR API KEY}

To generate API key refer https://macaddress.io/api/documentation/making-requests

2) export MACADRESS={Sample MAC Address}

Example 44:38:39:ff:ef:57

### Run go file

go run main.go

### Output

Entered MAC Address = 44:38:39:ff:ef:57

{"vendorDetails":{"oui":"443839","isPrivate":false,"companyName":"Cumulus Networks, Inc","companyAddress":"650 Castro Street, suite 120-245 Mountain View CA 94041 US","countryCode":"US"},"blockDetails":{"blockFound":true,"borderLeft":"443839000000","borderRight":"443839FFFFFF","blockSize":16777216,"assignmentBlockSize":"MA-L","dateCreated":"2012-04-08","dateUpdated":"2015-09-27"},"macAddressDetails":{"searchTerm":"44:38:39:ff:ef:57","isValid":true,"virtualMachine":"Not detected","applications":["Multi-Chassis Link Aggregation (Cumulus Linux)"],"transmissionType":"unicast","administrationType":"UAA","wiresharkNotes":"No details","comment":""}}
