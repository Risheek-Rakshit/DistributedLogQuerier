package client

import (
	"mp1-server/utils"
	"mp1-server/logger"
	"mp1-server/config"
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

/*
	Sending grep requests to the machines.
	Prints the grep responses as received
	Returns grep responses as results arrays, total lines mached and the absolute end time when query was served
*/
func PerformCombinedGrep(addresses []net.TCPAddr, input string, logger *logger.CustomLogger,config *config.Config) ([]utils.Results, int, int64){
		
	var wg sync.WaitGroup
	wg.Add(len(addresses))

	channel := make(chan utils.Results)
	for index, address := range addresses {
		go connectionHandler(index, address, &wg, input, channel, logger, config)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	
	totLine := 0
	resp := ""
	var respArr []utils.Results
	for message := range channel {
		totLine += message.LineCount
		resp = resp + message.Lines +"\n"
		respArr = append(respArr,message)
	}

	endTime := time.Now().UnixMilli()

	for _, message := range respArr {
		fmt.Println("Message from server i: ",message.NumMach)
		fmt.Println(message.Lines)
	}

	fmt.Println("Total Line matched:", totLine)

	return respArr, totLine, endTime
}

/*
	Run as a goroutine
	Gives user option to either enter command, or perform Tests as per config/clientTester.go
	Test inputs (type this when prompted to perform the individual tests)- 
	0: Grep on a rare pattern
	1: Grep on a frequent pattern
	2: Grep on a somewhat frequent pattern
	3: Grep on a pattern present on only 1 machine
	4: Grep with -c option
	5: Grep on a regex
	6: Grep when machine 10 fails
*/
func ClientImplNewTest(logger *logger.CustomLogger,config *config.Config){
	for {
		
		fmt.Println("Enter your grep command, or Test option(0-6) [Refer to the README about what each test option performs]:")
		scanner := bufio.NewScanner(os.Stdin)
		line := ""
		if scanner.Scan() {
    		line = scanner.Text()
		}
		
		switch line {
		case "0":
			TestCombinedGrepRare(logger, config)
		case "1":
			TestCombinedGrepFrequent(logger, config)
		case "2":
			TestCombinedGrepSomewhatFreq(logger, config)
		case "3":
			TestCombinedGrepSingle(logger, config)
		case "4":
			TestCombinedGrepCount(logger, config)
		case "5":
			TestCombinedGrepRegex(logger, config)
		case "6":
			TestCombinedGrepFT(logger, config)
		default:
			startTime := time.Now().UnixMilli()
			addresses := utils.AddressParser(logger, config)

			for _, address := range addresses {
				logger.Debug(string(address.IP))
				logger.Debug(strconv.Itoa(address.Port))
			}
		
			_,_, endTime := PerformCombinedGrep(addresses, line, logger, config)

			latency := endTime - startTime
			fmt.Println("Time taken for the query (in ms): ",latency)
		}
	}
}

/*
	Run as a goroutine
	takes grep command as a user input, and calls PerformCommbinedGrep. Displays the total time taken by the query
*/
func ClientImplNew(logger *logger.CustomLogger,config *config.Config){
	for {
		
		fmt.Println("Enter your grep command:")
		scanner := bufio.NewScanner(os.Stdin)
		line := ""
		if scanner.Scan() {
    		line = scanner.Text()
		}

		startTime := time.Now().UnixMilli()
		addresses := utils.AddressParser(logger, config)

		for _, address := range addresses {
			logger.Debug(string(address.IP))
			logger.Debug(strconv.Itoa(address.Port))
			
		}
		
		_,_, endTime := PerformCombinedGrep(addresses, line, logger, config)

		latency := endTime - startTime
		fmt.Println("Time taken for the query (in ms): ",latency)

	}
	
}

/*
	performs connections to a machine with addr IPv4 address, to send a grep request, and receive for the response.
	store the response in a channel ch
*/
func connectionHandler(machNumber int, addr net.TCPAddr, wg *sync.WaitGroup, pattern string, ch chan utils.Results, logger *logger.CustomLogger, config *config.Config) {
	machNumber 	= machNumber + 1
	logFile := config.LogPath + "/machine." + strconv.Itoa(machNumber) + ".log"
	logger.Info("Contacting machine:" + strconv.Itoa(machNumber))
	address := addr.IP.String() + ":" + strconv.Itoa(addr.Port)
	conn, err := net.Dial("tcp", address)

	
	defer wg.Done()

	if err != nil {
		logger.Error("Error in contacting the Machine: "+ strconv.Itoa(machNumber), err)
		return
	}

	defer conn.Close()
	

	payload := utils.Payload{Name:logFile,Pattern: pattern}
	logger.Info("Message writing")
	encoder := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)
	encoder.Encode(payload)
	logger.Info("Message written")

	err = conn.SetReadDeadline(time.Now().Add(time.Duration(config.TimeOut) * time.Second))
	if err != nil {
		logger.Error("Error in read time", err)
		return
	}

	p := &utils.Results{}
	err = dec.Decode(p)
	if err!= nil {
		logger.Error("Error decoding message: ", err)
		return
	}
	p.NumMach = machNumber

	ch <- *p
}
