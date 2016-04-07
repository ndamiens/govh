package ip

import (
	"fmt"
	"net"
	"time"

	"github.com/toorop/govh"
)

// IPType enumerates each type of IP
var IPType = [10]string{"cdn", "dedicated", "hosted_ssl", "loadBalancing", "mail", "pcc", "pci", "vpn", "vps", "xdsl"}

// IP is a string representation of an IP
//type IP string

// IPBlock represents represents OVH ipBlock type
type IPBlock string

// GetIPs return IPs in IPblocks
func (i *IPBlock) GetIPs() (IPs []string, err error) {
	ip, ipNet, err := net.ParseCIDR(string(*i))
	if err != nil {
		return
	}
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		IPs = append(IPs, ip.String())
	}
	return
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// RoutedTo represents ip.routedTo OVH type
type RoutedTo struct {
	ServiceName string
}

// MoveTo represents move property of an IP
type MoveTo struct {
	To string `json:"to"`
}

type IPTaskId int

// IpTask represents a task on an IP (answer of MoveTo)
type IpTask struct {
	TaskID      IPTaskId `json:"taskId"`
	Function    string   `json:"function"`
	LastUpdate  string   `json:"lastUpdate"`
	Comment     string   `json:"comment"`
	Status      string   `json:"status"`
	StartDate   string   `json:"startDate"`
	DoneDate    string   `json:"doneDate"`
}

func (it IpTask) String() string {
	return fmt.Sprintf("Task: %d\nFunction: %s\nComment: %s\nStatus: %s\nStart Date: %s\nDone Date: %s\n", it.TaskID, it.Function, it.Comment, it.Status, it.StartDate, it.DoneDate)
}

// IP represents OVH ip.Ip type
type IP struct {
	OrgranisationID string   `json:"organisationId"`
	Country         string   `json:"country"`
	RoutedTo        RoutedTo `json:"routedTo"`
	IPBlock         IPBlock  `json:"ip"`
	CanBeTerminated bool     `json:"canBeTerminated"`
	Type            string   `json:"type"`
	Description     string   `json:"description"`
}

// String return the string representations of IP
func (i IP) String() string {
	return fmt.Sprintf("Block: %s\nDescription: %s\nRouted To: %s\nType: %s\nCountry:%s\nCan be terminated: %t", i.IPBlock, i.Description, i.RoutedTo.ServiceName, i.Type, i.Country, i.CanBeTerminated)
}

// IpUpdatableProperties represents updatable properties of an IP
type IpUpdatableProperties struct {
	Description string `json:"description,omitempty"`
}

//FirewalledIp represents an IP on the Firewall
type FirewalledIp struct {
	Ip      string `json:"ipOnFirewall"`
	Enabled bool   `json:"enabled"`
	State   string `json:"state"`
}

// Firewall rules

// destinationPort
type DestinationPort struct {
	From int `json:"from"`
	To   int `json:"to"`
}

// sourcePort
type SourcePort struct {
	From int `json:"from"`
	To   int `json:"to"`
}

// fwTcpOption represents TCP option for a firewall rule
type FwTcpOption struct {
	Fragments bool   `json:"fragments,omitempty"`
	Option    string `json:"option,omitempty"`
}

// FwFirewallRule2Add
type FwRule2Add struct {
	Action    string       `json:"action"`
	ToPort    int          `json:"destinationPort,omitempty"`
	Protocol  string       `json:"protocol"`
	Sequence  int          `json:"sequence"`
	FromIp    string       `json:"source,omitempty"`
	FromPort  int          `json:"sourcePort,omitempty"`
	TcpOption *FwTcpOption `json:"tcpOption,omitempty"`
}

// Reply
type FirewallRule struct {
	Protocol     string        `json:"protocol"`
	FromIp       string        `json:"source"`
	ToPort       string        `json:"destinationPort"`
	Sequence     int           `json:"sequence"`
	TcpOption    string        `json:"tcpOption"`
	ToIp         string        `json:"destination"`
	Rule         string        `json:"rule"`
	FromPort     string        `json:"sourcePort"`
	State        string        `json:"state"`
	CreationDate govh.DateTime `json:"creationDate"`
	Action       string        `json:"action"`
	Fragments    bool          `json:"fragments"`
}

// ReverseIP represents an OVH reverseIp type
type ReverseIP struct {
	IPReverse string `json:"ipReverse"`
	Reverse   string `json:"reverse"`
}

// String() returns ReverseIP as string
func (r ReverseIP) String() string {
	return fmt.Sprintf("%s %s", r.IPReverse, r.Reverse)
}

//
//// SPAM
//

// SpamIP represents an OVH ip.SpamIp type
type SpamIP struct {
	IP    string        `json:"ipSpamming"` // IP address which is sending spam
	Time  int           `json:"time"`       // Time (in seconds) while the IP will be blocked
	Date  govh.DateTime `json:"date"`       // Last date the ip was blocked
	State string        `json:"state"`      // Current state of the ip. blockedForSpam | unblocked | unblocking
}

// Stringer
func (s SpamIP) String() string {
	out := "IP: " + s.IP + "\n"
	out += fmt.Sprintf("Blocked since: %d seconds\n", s.Time)
	out += fmt.Sprintf("Last blocked: %v\n", s.Date)
	out += "state: " + s.State + "\n"
	return out
}

// SpamTarget represents an OVH ip.SpamTarget type
type SpamTarget struct {
	DestinationIP string `json:"destinationIp"` // IP address of the target
	MessageID     string `json:"messageId"`     // The message-id of the email
	Date          int64  `json:"date"`          // Timestamp when the email was sent
	Spamscore     int    `json:"spamscore"`     // Spam score of the email
	//Spamcause     string `json:"spamcause"`     // Detailled spam cause
}

// SpamTarget stringer
func (s SpamTarget) String() string {
	out := "Destination: " + s.DestinationIP + "\n"
	out += "Date: " + time.Unix(s.Date, 0).String() + "\n"
	out += "Message-id: " + s.MessageID + "\n"
	out += fmt.Sprintf("Score: %d\n", s.Spamscore)
	return out
}

// SpamStats represents an OVH ip.SpamStats type
type SpamStats struct {
	Timestamp        int64 `json:"timestamp"` // Time when the IP address was blocked
	DetectedSpams    []SpamTarget
	AverageSpamScore int `json:"averageSpamscore"` // Average spam score.
	Total            int `json:"total"`            // Number of emails sent
	NumberOfSpams    int `json:"numberOfSpams"`    //Number of spams sent
}

// SpamStats stringer
func (s SpamStats) String() string {
	out := "Blocked: " + time.Unix(s.Timestamp, 0).String() + "\n"
	out += fmt.Sprintf("Email total: %d\n", s.Total)
	out += fmt.Sprintf("Number of spams: %d\n", s.NumberOfSpams)
	out += fmt.Sprintf("Average score: %d\n", s.AverageSpamScore)
	if len(s.DetectedSpams) != 0 {
		out += "Spams:\n"
		for _, spam := range s.DetectedSpams {
			out += spam.String()
		}
	}
	return out
}
