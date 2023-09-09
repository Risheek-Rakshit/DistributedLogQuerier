package client

import (
	"testing"
	"time"
	"fmt"
	"mp1-server/utils"
	"mp1-server/config"
	"mp1-server/logger"
)

func TestCombinedGrepRare(t *testing.T){

	Logger := logger.NewLogger("debug")
	config := &config.Config{
		Port: "4040",
		AddressPath: "../addresses.json",
		TimeOut: 5,
		LogPath: "log",
	}
	addresses := utils.AddressParser(Logger, config)
	command := "grep stars"

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)

	expResult := ""
	expNumLines := 10

	if result != expResult || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + result)
		t.Errorf("Grep for rare pattern occurences failed")
	}
}

func TestCombinedGrepFrequent(t *testing.T){
	Logger := logger.NewLogger("debug")
	config := &config.Config{
		Port: "4040",
		AddressPath: "../addresses.json",
		TimeOut: 5,
		LogPath: "log",
	}
	addresses := utils.AddressParser(Logger, config)
	command := "grep stars"

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)

	expResult := ""
	expNumLines := 10

	if result != expResult || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + result)
		t.Errorf("Grep for frequent pattern occurences failed")
	}
}

func TestCombinedGrepSomewhatFreq(t *testing.T){
	Logger := logger.NewLogger("debug")
	config := &config.Config{
		Port: "4040",
		AddressPath: "../addresses.json",
		TimeOut: 5,
		LogPath: "log",
	}
	addresses := utils.AddressParser(Logger, config)
	command := "grep stars"

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)

	expResult := ""
	expNumLines := 10

	if result != expResult || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + result)
		t.Errorf("Grep for somewhat pattern occurences failed")
	}
}

func TestCombinedGrepSingle(t *testing.T){
	Logger := logger.NewLogger("debug")
	config := &config.Config{
		Port: "4040",
		AddressPath: "../addresses.json",
		TimeOut: 5,
		LogPath: "log",
	}
	addresses := utils.AddressParser(Logger, config)
	command := "grep 180.34.166.169"

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)

	expResult := "180.34.166.169 - - [30/Oct/2022:14:55:48 -0500] \"PUT /wp-admin HTTP/1.0\" 200 5024 \"http://www.mason.com/index/\" \"Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_7; rv:1.9.2.20) Gecko/2016-05-17 12:28:05 Firefox/3.6.20\"\n"
	expNumLines := 1

	if result != expResult || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + result)
		t.Errorf("Grep when pattern present only in a single machine failed")
	}
}

func TestCombinedGrepMulti(t *testing.T){
	Logger := logger.NewLogger("debug")
	config := &config.Config{
		Port: "4040",
		AddressPath: "../addresses.json",
		TimeOut: 5,
		LogPath: "log",
	}
	addresses := utils.AddressParser(Logger, config)
	command := "grep stars"

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)

	expResult := ""
	expNumLines := 10

	if result != expResult || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + result)
		t.Errorf("Grep for rare pattern occurences failed")
	}
}