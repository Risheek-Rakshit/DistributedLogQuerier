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

type MachineDetails struct {
	MachineName string `json:"machine"`
	PortNumber  int    `json:"port"`
}

type Mode int64
const (
	DEFAULT Mode = iota
	TEST
)


type Payload struct {
	Name    string
	Pattern string
}

type Results struct {
	Lines string
	LineCount int
	NumMach		int
}

func command(cmd string, logger *logger.CustomLogger) (string, error) {
	output, err := exec.Command("bash","-c",cmd).Output()
	strOutput, err2 := strconv.Atoi(string(output[:]))
	if err != nil && err2!= nil && strOutput!=0{
		logger.Error("Error while executing system cmd",err)
		return "",err
	}

	return string(output[:]), nil
}

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


func GrepFileNew(filePath string, cmd string, logger *logger.CustomLogger) (string, int, error) {
	cmdLines := cmd + " " + filePath

	outputLines, err := command(cmdLines, logger)
	if err != nil {
		logger.Error("Error while executing lines grep system cmd",err)
		return "",-1,err
	}

	lines := strings.Split(outputLines, "\n")

	numLines := len(lines) - 1

	result := outputLines

	return result, numLines, nil
}

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
