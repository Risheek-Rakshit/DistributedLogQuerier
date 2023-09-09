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

func PerformCombinedGrep(addresses []net.TCPAddr, input string, logger *logger.CustomLogger,config *config.Config) (string, int, int64){
		
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

	endTime := time.Now().UnixMilli()
	totLine := 0
	resp := ""
	for message := range channel {
		fmt.Println("Message from server i: ",message.NumMach)
		fmt.Println(message.Lines)
		totLine += message.LineCount
		resp = resp + message.Lines +"\n"
	}
	fmt.Println("Total Line matched:", totLine)

	return resp, totLine, endTime
}

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

func connectionHandler(machNumber int, addr net.TCPAddr, wg *sync.WaitGroup, pattern string, ch chan utils.Results, logger *logger.CustomLogger, config *config.Config) {
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
