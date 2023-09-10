package utils

import (
	"encoding/json"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"mp1-server/logger"
	"mp1-server/config"
)

/*
	Stores the Machine name and port number (currently, always 4040)
*/
type MachineDetails struct {
	MachineName string `json:"machine"`
	PortNumber  int    `json:"port"`
}

/*
	Mode enables running distributed tests on the program.
	0 means DEFAULT mode: normal grep commands
	1 means TEST mode: it asks users to perform tests
*/
type Mode int64
const (
	DEFAULT Mode = iota
	TEST
)

/*
	Payload is what server receives from the client as a request
	Name is the filename to search onthat server
	Pattern is the grep command
*/
type Payload struct {
	Name    string
	Pattern string
}

/*
	Result is what server sends to the client as a response
	Lines is the grep command's result
	LineCount is the line count of the command, (in case of a -c, it is the absolute count)
	NumMach is the machine number
*/
type Results struct {
	Lines string
	LineCount int
	NumMach		int
}

/*
	runs a bash command given by cmd
*/
func command(cmd string, logger *logger.CustomLogger) (string, error) {
	output, err := exec.Command("bash","-c",cmd).Output()
	strOutput, err2 := strconv.Atoi(string(output[:]))
	if err != nil && err2!= nil && strOutput!=0{
		logger.Error("Error while executing system cmd",err)
		return "",err
	}

	return string(output[:]), nil
}

/*
	[DEPRECATED]: use GrepFileNew. This assumes that pattern is not a grep command, and exhaustively performs -c and -n options.
	Use GrepFileNew instead, which is more generic
*/
func GrepFile(filePath string, pattern string, logger *logger.CustomLogger) (string, int, error) {
	cmdCount := "grep -c \"" + pattern + "\" " + filePath
	cmdLines := "grep -n \"" + pattern + "\" " + filePath

	outputCount, err := command(cmdCount, logger)
	if err != nil {
		logger.Error("Error while executing count grep system cmd", err)
		return "",-1,err
	}
	
	outputLines, err := command(cmdLines, logger)
	if err != nil {
		logger.Error("Error while executing lines grep system cmd",err)
		return "",-1,err
	}

	numLines, err := strconv.Atoi(outputCount[:len(outputCount)-1])
	if err !=nil {
		logger.Error("Error while converting into integer",err)
		return "",-1,err
	}

	result := "Number of matching lines: " + outputCount + "---------------------------------------\n" + outputLines

	return result, numLines, nil
}

/*
	performs grep command represented by cmd on the filePath.
	This is a server side utils function
	returns the grep result and number of lines
*/
func GrepFileNew(filePath string, cmd string, logger *logger.CustomLogger) (string, int, error) {
	option := strings.Split(cmd, " ")[1]
	
	cmdLines := cmd + " " + filePath

	outputLines, err := command(cmdLines, logger)
	if err != nil {
		logger.Error("Error while executing lines grep system cmd",err)
		return "",-1,err
	}

	numLines := 0

	//Handle -c
	if strings.Contains(option, "-") && strings.Contains(option, "c") {
		numLines, err = strconv.Atoi(outputLines[:len(outputLines)-1])
		if err !=nil {
			logger.Error("Error while converting into integer",err)
			return "",-1,err
		}
	} else {
		lines := strings.Split(outputLines, "\n")
		numLines = len(lines) - 1
	}
	

	result := outputLines

	return result, numLines, nil
}

/*
	Performs an IPLookup on machine names present in the location as mentioned in the config file
	This is a client side util function
	returns a list of resolved TCPAddr
*/
func AddressParser(logger *logger.CustomLogger, config *config.Config) []net.TCPAddr {
	var names []MachineDetails
	var addresses []net.TCPAddr
	file, err := os.ReadFile(config.AddressPath)
	if err != nil {
		logger.Error("Error opening file: ", err)
		return nil
	}

	logger.Info("opening file")

	json.Unmarshal(file, &names)

	for _, address := range names {
		logger.Debug(string(address.MachineName))
		logger.Debug(strconv.Itoa(address.PortNumber))
	}
	for _, machine := range names {
		ip, err := net.LookupIP(machine.MachineName)
		if err != nil {
			logger.Error("Error fetching IP details for" + machine.MachineName, err)
			return nil
		}

		for _, ip := range ip {
			ipv4 := ip.To4()
			if ipv4 == nil {
				continue
			}
			tcpAddr := &net.TCPAddr{
				IP:   ipv4,
				Port: machine.PortNumber,
			}
			addresses = append(addresses, *tcpAddr)
		}

	}

	return addresses
}
