package utils

import (
	"testing"
	"mp1-server/logger"
)

func TestGrepFileNew(t *testing.T){
	testFileName := "test.log"
	testCommand := "grep it"

	Logger := logger.NewLogger("debug")
	result, numLines, err := GrepFileNew(testFileName,testCommand, Logger)
	if err!= nil {
		Logger.Error("[Test] Some error while unit test grep", err)
		t.Errorf("Exiting...")
	}

	expResult := "Nobody said it was easy\nNobody said it was easy\nNo one ever said it would be this hard\nNobody said it was easy\nOh, it's such a shame for us to part\nNobody said it was easy\nNo one ever said it would be so hard"
	expNumLines := 7

	if result != expResult && numLines != expNumLines {
		Logger.Info(result)
		t.Errorf("Grep is not returning correct result")
	}

}