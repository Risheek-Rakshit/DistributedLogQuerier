# CS425 - MP1 - Group 58

The MP1 focuses on developing a log querier which allows the user to execute grep commands on multiple machines in parallel.


## Overview

The log querier is built using Golang. It leverages Go routines to concurrently send requests to multiple machines and collect them.

## Usage
* The 'addresses.json' file will serve as a repository for specifying the machine addresses and ports in our network.
* Place the log files on the machines with the naming convention 'machine.[i].log' where i corresponds to the system's position in the 'addresses.json' file.
```bash
log/machine.[i].log
```
* You can also use your custom json file, but ensure that the content is of similar types. Update the config.yaml file to point to the desired file.
```bash
Constants to be defined in config.yaml file
port : The default port servers
addresspath: The JSON file containing the machine names
timeout: The time for which the client waits for response from server
logpath: The directory that contains the machine logs.
```
* Run the program 'main.go' in all the systems. The program invokes the client and server implementations on the machines it is run.
* The client implementation of the program waits for the input from the user. Use any of the active machines to enter the command. The command must be of the form
```bash
grep <option> <pattern or regEx>
```

## Installation
* Install Golang version 1.19
* Clone the repository
```bash
https://gitlab.engr.illinois.edu/sl203/mp1-cs425-group58.git
```
* Build the project
```bash
go build
```
* Run the program in the required mode. 0 represents default mode and 1 represents test mode. 
```bash
./mp1-server -loglev=<LogLevel> -mode=<0 or 1>
```
* When run in the test mode, mention the category type to grep a specific predefined command for each of the following category. Look at the Testing section for more details.
```bash
0 : Rare pattern
1 : Frequent pattern
2 : Somewhat Frequent pattern
3 : Pattern found only in one machine
4 : Only the count of the lines matched in each file
5 : Regular Expression 
```

## Testing
* We have unit tests for grep and go-exec in utils package. To run them:
```bash
cd utils
go test
```
This will run unit tests for utils
* We have test mode for distributed testing, This assumes all 10 machines are up. We are using the provided Demo logs. Before running in test mode please download the expected outputs from the given drive link and place the folder 'testOuput' inside the client directory for proper results.
```bash
https://drive.google.com/drive/folders/1mk_2KE3NFMRh3sHG_Trzw48-H192q0c2?usp=sharing
```
## Installation
* Install Golang version 1.19
* Clone the repository
```bash
https://gitlab.engr.illinois.edu/sl203/mp1-cs425-group58.git
```
* Build the project
```bash
go build
```
* Run the program in the required mode. 0 represents default mode and 1 represents test mode. 
```bash
./mp1-server -loglev=<LogLevel> -mode=<0 or 1>
```
* When run in the test mode, mention the category type to grep a specific predefined command for each of the following category.
```bash
0 : Rare pattern
1 : Frequent pattern
2 : Somewhat Frequent pattern
3 : Pattern found only in one machine
4 : Only the count of the lines matched in each file
5 : Regular Expression 
```
* Before running in test mode please unzip the test.zip file and place the folder 'testOuput' inside the client folder for proper results.
## Report
Please refer to < > for further insights into design and performance.
