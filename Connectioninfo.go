/*
This file is part of TS3QueryApi For Go.
TS3QueryApi For Go is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package ts3api

import (
	"errors"
	"strconv"
	"strings"
)

type Connectioninfo struct {
	ConnectionFiletransferBandwidthSent        uint64
	ConnectionFiletransferBandwidthReceived    uint64
	ConnectionFiletransferBytesSentTotal       int
	ConnectionFiletransferBytesReceivedTotal   int
	ConnectionPacketsSentTotal                 int
	ConnectionBytesSentTotal                   uint64
	ConnectionPacketsReceivedTotal             int
	ConnectionBytesReceivedTotal               uint64
	ConnectionBandwidthSentLastSecondTotal     int
	ConnectionBandwidthSentLastMinuteTotal     int
	ConnectionBandwidthReceivedLastSecondTotal int
	ConnectionBandwidthReceivedLastMinuteTotal int
	ConnectionConnectedTime                    int
	ConnectionPacketlossTotal                  float32
	ConnectionPing                             float32
}

func (this *Connectioninfo) parseMsg(s string) {
	params := strings.Split(s, " ")
	for _, param := range params {
		if strings.Contains(param, "=") {
			kvpair := strings.SplitN(param, "=", 2)
			this.parseParam(kvpair[0], kvpair[1])
		} else {
			this.parseParam(param, "")
		}
	}
}
func (this *Connectioninfo) parseParam(k, v string) (err error) {
	switch k {
	case "connection_filetransfer_bandwidth_sent":
		this.ConnectionFiletransferBandwidthSent, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.ConnectionFiletransferBandwidthSent = 0
			return
		}

	case "connection_filetransfer_bandwidth_received":
		this.ConnectionFiletransferBandwidthReceived, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.ConnectionFiletransferBandwidthReceived = 0
			return
		}

	case "connection_filetransfer_bytes_sent_total":
		this.ConnectionFiletransferBytesSentTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionFiletransferBytesSentTotal = -1
			return
		}

	case "connection_filetransfer_bytes_received_total":
		this.ConnectionFiletransferBytesReceivedTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionFiletransferBytesReceivedTotal = -1
			return
		}

	case "connection_packets_sent_total":
		this.ConnectionPacketsSentTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionPacketsSentTotal = -1
			return
		}

	case "connection_bytes_sent_total":
		this.ConnectionBytesSentTotal, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.ConnectionBytesSentTotal = 0
			return
		}

	case "connection_packets_received_total":
		this.ConnectionPacketsReceivedTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionPacketsReceivedTotal = -1
			return
		}

	case "connection_bytes_received_total":
		this.ConnectionBytesReceivedTotal, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.ConnectionBytesReceivedTotal = 0
			return
		}

	case "connection_bandwidth_sent_last_second_total":
		this.ConnectionBandwidthSentLastSecondTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBandwidthSentLastSecondTotal = -1
			return
		}

	case "connection_bandwidth_sent_last_minute_total":
		this.ConnectionBandwidthSentLastMinuteTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBandwidthSentLastMinuteTotal = -1
			return
		}

	case "connection_bandwidth_received_last_second_total":
		this.ConnectionBandwidthReceivedLastSecondTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBandwidthReceivedLastSecondTotal = -1
			return
		}

	case "connection_bandwidth_received_last_minute_total":
		this.ConnectionBandwidthReceivedLastMinuteTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBandwidthReceivedLastMinuteTotal = -1
			return
		}

	case "connection_connected_time":
		this.ConnectionConnectedTime, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionConnectedTime = -1
			return
		}

	case "connection_packetloss_total":
		var g float64
		g, err = strconv.ParseFloat(v, 32)
		if err != nil {
			this.ConnectionPacketlossTotal = 0.0
			return
		}
		this.ConnectionPacketlossTotal = float32(g)

	case "connection_ping":
		var g float64
		g, err = strconv.ParseFloat(v, 32)
		if err != nil {
			this.ConnectionPing = 0.0
			return
		}
		this.ConnectionPing = float32(g)

	default:
		logger.Error("%s=%s is not a valid param of Connectioninfo.", k, v)
		err = errors.New(k + "=" + v + " is not a valid param of Connectioninfo.")
	}
	return
}
