package utils

import (
	"testing"
	"mp1-server/logger"
)

/**
	Unit tests for grep and running a bash command
**/

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