package shared

import (
	"time"
)

type GatewayStatus struct {
	Eui               string    `json:"eui"`
	Time              time.Time `json:"time"`
	Latitude          *float64  `json:"latitude,omitempty"`
	Longitude         *float64  `json:"longitude,omitempty"`
	Altitude          *float64  `json:"altitude,omitempty"`
	RxCount           *uint     `json:"rxCount,omitempty"`
	RxOk              *uint     `json:"rxOk,omitempty"`
	RxForwarded       *uint     `json:"rxForwarded,omitempty"`
	AckRatio          *float64  `json:"ackRatio,omitempty"`
	DatagramsReceived *uint     `json:"datagramsReceived,omitempty"`
	DatagramsSent     *uint     `json:"datagramsSent,omitempty"`
}

type RxPacket struct {
	GatewayEui string    `json:"gatewayEui"`
	NodeEui    string    `json:"nodeEui"`
	Time       time.Time `json:"time"`
	RawData    string    `json:"rawData"`
	Data       string    `json:"data,omitempty"`
        Rssi       *int      `json:"rssi,omitempty"`
        Snr        *float64  `json:"lsnr,omitempty"`
}

type ConsumerQueues struct {
	GatewayStatuses chan *GatewayStatus
	RxPackets       chan *RxPacket
}
