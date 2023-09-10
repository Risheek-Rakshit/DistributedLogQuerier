package main

import (
	"net"
	"fmt"
	"flag"

	"encoding/gob"

	"mp1-server/client"
	"mp1-server/logger"
	"mp1-server/utils"
	"mp1-server/config"
)

/*
	conn: client-Server connection
	logger: a custom logger

	This is run as a gouroutine, it receives client requests, performs combined grep from all peers, and returns results on client
*/
func handleConnection(conn net.Conn, logger *logger.CustomLogger){

	logger.Info("New client connection, serving request")

	dec := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	p := &utils.Payload{}
	dec.Decode(p)

	logger.Debug("File to be searched:" + p.Name)
	logger.Debug("Command to be executed:" + p.Pattern)

	defer conn.Close()

	result, num, err := utils.GrepFileNew(p.Name, p.Pattern, logger)
	if err!=nil {
		logger.Error("Some error while grep", err)
		return
	}

	payload := utils.Results{Lines: result, LineCount: num}
	logger.Debug("Message writing to client")
	encoder.Encode(payload)
	logger.Debug("Message written to client")
	
}


func main(){
	
	//Flag to set log level and testing
	help := flag.Bool("h", false, "help")

	logLevel := "info"
	mode := 0
	flag.IntVar(&mode, "mode", 0, "sets the mode to use: 0 is default mode, 1 is testing mode, where optionally all distributed tests can be performed")
	flag.StringVar(&logLevel, "loglev", "info", "Set log level: can be debug, info, error. Any other value returns info, set to info by default")
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}
	
	//declaring logger and configs
	Logger := logger.NewLogger(logLevel)
	config := config.NewConfig(Logger)


	Logger.Info("Welcome to CS425 MP1, Using log level: "+logLevel)

	//Launching Server in main thread
	listner, err := net.Listen("tcp", "0.0.0.0:" + config.Port)
	if err != nil {
		Logger.Error("Error while starting server", err)
		return
	}
	Logger.Info("Server Started at Port: " + config.Port)

	//Launch client subroutine as per the mode
	if utils.Mode(mode) == utils.TEST {
		go client.ClientImplNewTest(Logger, config)
	} else {
		go client.ClientImplNew(Logger, config)
	}
	
	//Accept client connections
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("Error while accepting listers", err)
		}
		go handleConnection(conn, Logger)
	}

}