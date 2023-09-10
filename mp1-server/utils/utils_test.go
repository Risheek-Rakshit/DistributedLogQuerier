package utils

import (
	"testing"
	"mp1-server/logger"
	// "mp1-server/config"
	// "net"
)

func TestGrepFileNew(t *testing.T){

	testFileName := "test.log"
	testCommand := "grep it"

	Logger := logger.NewLogger("debug")
	Logger.Info("Unit testing simple grep on a test.log file")
	result, numLines, err := GrepFileNew(testFileName,testCommand, Logger)
	if err!= nil {
		Logger.Error("[Test] Some error while unit test grep", err)
		t.Errorf("Exiting...")
	}

	expResult := "Nobody said it was easy\nNobody said it was easy\nNo one ever said it would be this hard\nNobody said it was easy\nOh, it's such a shame for us to part\nNobody said it was easy\nNo one ever said it would be so hard\n"
	expNumLines := 7

	if result != expResult || numLines != expNumLines {
		Logger.Info("[Test] This is Wrong: "+result)
		t.Errorf("Grep is not returning correct result")
	}

}

func TestCommand(t *testing.T){

	testCommand := "echo HelloCovfefeInc"

	Logger := logger.NewLogger("debug")
	result, err := command(testCommand, Logger)
	if err!= nil {
		Logger.Error("[Test] Some error while unit test exec command", err)
		t.Errorf("Exiting...")
	}

	expResult := "HelloCovfefeInc\n"

	if result != expResult {
		Logger.Info("[Test] This is Wrong: "+result)
		t.Errorf("Exec Command is not returning correct result")
	}
}

//Deprecated: Since IPlookup for google.com, facebook.com is machine dependent, this doesn't make sense
/*
func TestAddressParser(t *testing.T){

	var testAddresses []net.TCPAddr
	googleAddr := &net.TCPAddr{
		IP: net.ParseIP("172.217.2.46"),
		Port: 4040,
	}
	facebookAddr := &net.TCPAddr{
		IP: net.ParseIP("157.240.249.35"),
		Port: 4040,
	}
	testAddresses = append(testAddresses,*googleAddr)
	testAddresses = append(testAddresses,*facebookAddr)
	
	Logger := logger.NewLogger("debug")
	config := &config.Config{
		Port: "4040",
		AddressPath: "testAddr.json",
		TimeOut: 5,
		LogPath: "log",
	}

	result := AddressParser(Logger, config)
	if result[0].IP.String() != testAddresses[0].IP.String() || result[1].IP.String() != testAddresses[1].IP.String(){
		Logger.Info("[Test] The result is Wrong: " + result[0].IP.String())
		t.Errorf("AddressParser is not returning correct result")
	}
}
*/