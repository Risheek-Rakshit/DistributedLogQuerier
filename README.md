# Distributed Log Querier

The project focuses on developing a log querier which allows the user to execute grep commands on multiple machines in parallel.


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
port : The default port servers (currently 4040)
addresspath: The JSON file containing the machine names
timeout: The time for which the client waits for response from server
logpath: The directory that contains the machine logs.
```
* The program invokes the client and server implementations on the machines it is run.
* The client implementation of the program waits for the input from the user. 

## Installation and Running
* Install Golang version 1.19
* Clone the repository
```bash
https://gitlab.engr.illinois.edu/sl203/mp1-cs425-group58.git
```
* Build the project
```bash
cd mp1-server
go build
```
* Run the program in the required mode in mp1-server directory. 0 represents default mode and 1 represents test mode. loglev can be info, debug or error. by default, loglev is info, and mode is 0
```bash
./mp1-server -loglev=<LogLevel> -mode=<0 or 1>
```

* The above will launch the program which waits for a user input (will work on every machine it is launched). Type in your grep query command. The command must be of the form.
```bash
grep <option> <pattern or regEx>
```

* Refer to the Testing section to see how to use the testing mode


## Testing
* We have standard go unit tests for grep and go-exec in utils package. To run them:
```bash
cd utils
go test
```
This will run unit tests for utils, regardless of the network being up.

* We have the test mode (mode 1) for distributed testing, This assumes all 10 machines are up initially. We are using the provided Demo logs. Before running in test mode please download the expected outputs from the given drive link and place the folder 'testOuput' inside the client directory for proper results.
```bash
https://drive.google.com/drive/folders/1mk_2KE3NFMRh3sHG_Trzw48-H192q0c2?usp=sharing
```
* After the above step, perform distributed tests, run the program like this (mode 1):
```bash
./mp1-server -loglev=debug -mode=1
```
* You will be propted by the program to enter a number between 0-6. They correspond to the following grep specific predefined command. Type in the number for test to perform.
```bash
0 : grep on Rare pattern
1 : grep on Frequent pattern
2 : grep on Somewhat Frequent pattern
3 : grep on Pattern found only in one machine
4 : grep to see working of -c: the count of the lines matched
5 : grep to see working of -E: regEx
6 : grep when machine 10 fails (Note: you must kill machine 10's process manually to pass this test)
```

Note: Distributed tests categories 0-5 assume logs are coming from all 10 machines, whereas category 6 assumes machine 10 is down. You must kill the process on machine 10 for this test to be successful.


## Report
Please refer to mp1-report.pdf for further insights into design and performance.
