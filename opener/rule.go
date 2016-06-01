// Package opener package contains all the logic to open access to an EC2 instance or instances
package opener

import "errors"

// RuleProtocol represents the valid protocols for an EC2 port listener
type RuleProtocol int

// String implements stringer interface for RuleProtocol
func (r RuleProtocol) String() string {
	switch r {
	case ALL:
		return "-1"
	case TCP:
		return "tcp"
	case UDP:
		return "udp"
	case ICMP:
		return "icmp"
	default:
		return "Unknown"
	}
}

// ParseRuleProtocol parses a valid protocol
func ParseRuleProtocol(proto string) (RuleProtocol, error) {
	switch proto {
	case "all", "ALL":
		return ALL, nil
	case "tcp", "TCP":
		return TCP, nil
	case "udp", "UDP":
		return UDP, nil
	case "icmp", "ICMP":
		return ICMP, nil
	default:
		return 0, errors.New("not a valid protocol")
	}
}

const (
	// ALL protocols
	ALL RuleProtocol = iota
	// TCP protocol
	TCP
	// UDP Protocol
	UDP
	// ICMP Protocol
	ICMP
)

const (
	// AllCIDR is the CIDR that describes all addresses
	AllCIDR = "0.0.0.0/0"
)

// Rule defines the rule where the instances need to be accessed
type Rule struct {
	CIDR     string
	Port     int
	Protocol RuleProtocol
}

// NewRule creates a regular Rule
func NewRule(cidr string, port int, protocol RuleProtocol) *Rule {
	r := &Rule{
		CIDR:     cidr,
		Port:     port,
		Protocol: protocol,
	}

	return r
}

// NewOpenRule creates a rule to a port accessible with all supported protocols and from anywhere
func NewOpenRule(port int) *Rule {
	r := &Rule{
		CIDR:     AllCIDR,
		Port:     port,
		Protocol: ALL,
	}

	return r
}