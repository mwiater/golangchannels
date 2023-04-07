package network_test

import (
	"net"
	"testing"

	"github.com/mattwiater/golangchannels/network"
	"github.com/stretchr/testify/assert"
)

func TestGetOutboundIP(t *testing.T) {
	ipAddress := network.GetOutboundIP()
	testIpAddress := net.ParseIP("192.0.2.1/24")
	assert.IsType(t, testIpAddress, ipAddress)
}
