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

* Use any of the active machines to enter the command. The command must be of the form.
```bash
grep <option> <pattern or regEx>
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

## Report
Please refer to mp1-report.pdf for further insights into design and performance.
