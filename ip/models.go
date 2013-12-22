package ip

import (
//"time"
)

// Type OF IP
var IpType = []string{"cdn", "dedicated", "hosted_ssl", "loadBalancing", "mail", "pcc", "pci", "vpn", "vps", "xdsl"}

// IP
type IpBlock struct {
	// IP (eg 91.121.78.23/32")
	IP   string
	Type string // IpType
}

// ip.FirewallIp
type IpFirewallIp struct {
	IpOnFirewall string `json:"ipOnFirewall"`
	Enabled      bool   `json:"enabled"`
	State        string `json:"state"`
}

// Firewall rules

// destinationPort
type DestinationPort struct {
}

// tcpOption
type TcpOption struct {
	Urg         bool `json:"urg"`
	Psh         bool `json:"psh"`
	Ack         bool `json:"ack"`
	Established bool `json:"established"`
	Syn         bool `json:"syn"`
	Fin         bool `json:"fin"`
	Rst         bool `json:"rst"`
}

// udpOption
type udpOption struct {
	Fragments bool `json:"fragment"`
}

type FirewallRule struct {
	Protocol        string   `json:"protocol"`
	Source          string   `json:"source"`
	DestinationPort string   `json:"destinationPort"`
	Sequence        int      `json:"sequence"`
	Options         []string `json:"options"`
	Destination     string   `json:"destination"`
	Rule            string   `json:"rule"`
	SourcePort      string   `json:"sourcePort"`
	State           string   `json:"state"`
	CreationDate    string   `json:"creationDate"`
	Action          string   `json:"action"`
}