package client

import (
	"time"
	"fmt"
	"os"
	"sort"
	"strings"
	"mp1-server/utils"
	"mp1-server/config"
	"mp1-server/logger"
)

/**
	clientTester.go consists of all distributed tests	
**/

/*
	Performs grep on a rare pattern
*/
func TestCombinedGrepRare(Logger *logger.CustomLogger, config *config.Config){

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
	file, err := os.ReadFile("client/testOutput/rare.txt")
	if err != nil {
		Logger.Error("No file, Please download file. Error opening file: ", err)
		return
	}
	expResult := string(file)
	expNumLines := 41

	if strings.TrimRight(resp,"\r\n") != strings.TrimRight(expResult,"\r\n") || numLines != expNumLines {
		Logger.Info("[TEST] The result is Wrong: " + resp)
		Logger.Error("Grep for rare pattern occurences failed",nil)
		return
	}
	Logger.Info("[TEST] Distributed Test for rare pattern successful")
}

/*
	Performs grep on a frequent pattern
*/
func TestCombinedGrepFrequent(Logger *logger.CustomLogger, config *config.Config){
	
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
	file, err := os.ReadFile("client/testOutput/frequentNew.txt")
	if err != nil {
		Logger.Error("No file, Please download file. Error opening file: ", err)
		return
	}
	expResult := string(file)
	expNumLines := 1081851

	if strings.TrimRight(resp,"\r\n") != strings.TrimRight(expResult,"\r\n") || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + resp)
		Logger.Error("Grep for frequent pattern occurences failed", nil)
		return
	}

	Logger.Info("[TEST] Distributed Test for frequent pattern successful")
}

/*
	Performs grep on a somewhat frequent pattern
*/
func TestCombinedGrepSomewhatFreq(Logger *logger.CustomLogger, config *config.Config){
	
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
	file, err := os.ReadFile("client/testOutput/somewhatfrequentNew.txt")
	if err != nil {
		Logger.Error("No file, Please download file. Error opening file: ", err)
	}
	expResult := string(file)
	expNumLines := 337979

	if strings.TrimRight(resp,"\r\n") != strings.TrimRight(expResult,"\r\n") || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + resp)
		Logger.Error("Grep for somewhat pattern occurences failed", nil)
		return
	}

	Logger.Info("[TEST] Distributed Test for somewhat frequent pattern successful")
}

/*
	Performs grep on a pattern present on a single machine
*/
func TestCombinedGrepSingle(Logger *logger.CustomLogger, config *config.Config){
	
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

	if strings.TrimRight(resp,"\r\n") != strings.TrimRight(expResult,"\r\n") || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + resp)
		Logger.Error("Grep when pattern present only in a single machine failed", nil)
		return
	}

	Logger.Info("[TEST] Distributed Test for Grep on a single file successful")
}

/*
	Performs grep with -c 
*/
func TestCombinedGrepCount(Logger *logger.CustomLogger, config *config.Config){
	
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
		Logger.Error("Grep for -c failed", nil)
		return
	}

	Logger.Info("[TEST] Distributed Test for grep -c option file successful")

}

/*
	Performs grep on a regec pattern
*/
func TestCombinedGrepRegex(Logger *logger.CustomLogger, config *config.Config){
	
	addresses := utils.AddressParser(Logger, config)
	command := "grep -E \"(DELETE|PUT)\\b\""

	Logger.Info("Distributed Testing a regex pattern, command is: " + command)

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)


	sort.Slice(result, func(i, j int) bool { return result[i].NumMach < result[j].NumMach })
	resp := ""
	for _, r := range result {
		resp = resp + r.Lines
	}
	file, err := os.ReadFile("client/testOutput/regex.txt")
	if err != nil {
		Logger.Error("No file, Please download file. Error opening file: ", err)
	}
	expResult := string(file)
	expNumLines := 812631

	if strings.TrimRight(resp,"\r\n") != strings.TrimRight(expResult,"\r\n") || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + resp)
		Logger.Error("Grep for regex pattern occurences failed", nil)
		return
	}

	Logger.Info("[TEST] Distributed Test for regex pattern successful")
}

/*
	Performs grep to check fault tolerance. Assumes machine 10 dies before this test is performed
*/
func TestCombinedGrepFT(Logger *logger.CustomLogger, config *config.Config){
	
	addresses := utils.AddressParser(Logger, config)
	command := "grep \"Macintosh; U; Intel Mac OS X 10_5_7 rv:6.0; en-US\""

	Logger.Info("Distributed Testing a pattern with fault tolerance, command is: " + command)

	startTime := time.Now().UnixMilli()
	result, numLines, endTime := PerformCombinedGrep(addresses,command,Logger,config)
	latency := endTime - startTime
	fmt.Println("Time taken for the query (in ms): ",latency)


	sort.Slice(result, func(i, j int) bool { return result[i].NumMach < result[j].NumMach })
	resp := ""
	for _, r := range result {
		resp = resp + r.Lines
	}
	file, err := os.ReadFile("client/testOutput/ft.txt")
	if err != nil {
		Logger.Error("No file, Please download file. Error opening file: ", err)
	}
	expResult := string(file)
	expNumLines := 38

	if strings.TrimRight(resp,"\r\n") != strings.TrimRight(expResult,"\r\n") || numLines != expNumLines {
		Logger.Info("[Test] The result is Wrong: " + resp)
		Logger.Error("Grep for fault tolerance failed", nil)
		return
	}

	Logger.Info("[TEST] Distributed Test for pattern with fault tolerance (machine 10 down) successful")
}
