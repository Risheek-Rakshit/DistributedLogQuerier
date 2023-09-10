package client

import (
	"testing"
	"time"
	"fmt"
	"os"
	"sort"
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
		LogPath: "../log",
	}
	addresses := utils.AddressParser(Logger, config)
	command := "grep \"Macintosh; U; Intel Mac OS X 10_5_7 rv:6.0; en-US\""

	Logger.Info("Distributed Testing a Rare pattern, command is: " + command)

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)


	sort.Slice(result, func(i, j int) bool { return result[i].NumMach < result[j].NumMach })
	resp := ""
	for _, r := range result {
		resp = resp + r.Lines
	}
	file, err := os.ReadFile("testOutput/rare.txt")
	if err != nil {
		Logger.Error("Error opening file: ", err)
		t.Errorf("No File, please download the file as per Readme")
	}
	expResult := string(file)
	expNumLines := 41

	if resp != expResult || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + resp)
		t.Errorf("Grep for rare pattern occurences failed")
	}
}

func TestCombinedGrepFrequent(t *testing.T){
	Logger := logger.NewLogger("debug")
	config := &config.Config{
		Port: "4040",
		AddressPath: "../addresses.json",
		TimeOut: 5,
		LogPath: "../log",
	}
	addresses := utils.AddressParser(Logger, config)
	command := "grep Safari"

	Logger.Info("Distributed Testing a Frequent pattern, command is: " + command)

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)


	sort.Slice(result, func(i, j int) bool { return result[i].NumMach < result[j].NumMach })
	resp := ""
	for _, r := range result {
		resp = resp + r.Lines
	}
	file, err := os.ReadFile("testOutput/frequent.txt")
	if err != nil {
		Logger.Error("Error opening file: ", err)
		t.Errorf("No File, please download the file as per Readme")
	}
	expResult := string(file)
	expNumLines := 1081851

	if resp != expResult || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + resp)
		t.Errorf("Grep for frequent pattern occurences failed")
	}
}

func TestCombinedGrepSomewhatFreq(t *testing.T){
	
	Logger := logger.NewLogger("debug")
	config := &config.Config{
		Port: "4040",
		AddressPath: "../addresses.json",
		TimeOut: 5,
		LogPath: "../log",
	}
	addresses := utils.AddressParser(Logger, config)
	command := "grep \"/app/main/posts HTTP/1.0\""

	Logger.Info("Distributed Testing a somewhat Frequent pattern, command is: " + command)

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)


	sort.Slice(result, func(i, j int) bool { return result[i].NumMach < result[j].NumMach })
	resp := ""
	for _, r := range result {
		resp = resp + r.Lines
	}
	file, err := os.ReadFile("testOutput/somewhatfrequent.txt")
	if err != nil {
		Logger.Error("Error opening file: ", err)
		t.Errorf("No File, please download the file as per Readme")
	}
	expResult := string(file)
	expNumLines := 41

	if resp != expResult || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + resp)
		t.Errorf("Grep for somewhat pattern occurences failed")
	}
}

func TestCombinedGrepSingle(t *testing.T){
	Logger := logger.NewLogger("debug")
	config := &config.Config{
		Port: "4040",
		AddressPath: "../addresses.json",
		TimeOut: 5,
		LogPath: "../log",
	}
	addresses := utils.AddressParser(Logger, config)
	command := "grep 180.34.166.169"

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)

	expResult := "180.34.166.169 - - [30/Oct/2022:14:55:48 -0500] \"PUT /wp-admin HTTP/1.0\" 200 5024 \"http://www.mason.com/index/\" \"Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_7; rv:1.9.2.20) Gecko/2016-05-17 12:28:05 Firefox/3.6.20\"\n"
	expNumLines := 1

	resp := ""
	for _, r := range result {
		resp = resp + r.Lines
	}

	if resp != expResult || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + resp)
		t.Errorf("Grep when pattern present only in a single machine failed")
	}
}

func TestCombinedGrepCount(t *testing.T){
	Logger := logger.NewLogger("debug")
	config := &config.Config{
		Port: "4040",
		AddressPath: "../addresses.json",
		TimeOut: 5,
		LogPath: "../log",
	}
	addresses := utils.AddressParser(Logger, config)
	command := "grep -c DELETE"

	Logger.Info("Distributed Testing a count grep command: " + command)

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)

	sort.Slice(result, func(i, j int) bool { return result[i].NumMach < result[j].NumMach })
	resp := ""
	for _, r := range result {
		resp = resp + r.Lines
	}
	expResult := "27976\n26754\n26854\n27144\n27293\n26846\n27031\n27294\n26898\n26437\n"
	expNumLines := 270527

	if resp != expResult || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + resp)
		t.Errorf("Grep for somewhat pattern occurences failed")
	}
}