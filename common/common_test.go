package common_test

import (
	"testing"

	"github.com/mattwiater/golangchannels/common"
	"github.com/stretchr/testify/assert"
)

func TestConsoleColumnWidth(t *testing.T) {
	columnWidth := common.ConsoleColumnWidth("12345", 30)
	testColumnWidth := 30 - len("12345")
	assert.Equal(t, testColumnWidth, columnWidth)
}

func TestSplitStringLines(t *testing.T) {
	splitString := common.SplitStringLines("12345\n67890")
	testSplitString := []string{"12345", "67890"}
	assert.IsType(t, testSplitString, splitString)
	assert.Equal(t, testSplitString, splitString)
}
