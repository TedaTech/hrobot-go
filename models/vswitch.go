package models

// VSwitchResponse wraps a single vSwitch response.
type VSwitchResponse struct {
	VSwitch VSwitch `json:"vswitch"`
}

// VSwitch represents a Hetzner vSwitch.
type VSwitch struct {
	ID           int                `json:"id"`
	Name         string             `json:"name"`
	Vlan         int                `json:"vlan"`
	Cancelled    bool               `json:"cancelled"`
	Server       []VSwitchServer    `json:"server,omitempty"`
	Subnet       []VSwitchSubnet    `json:"subnet,omitempty"`
	CloudNetwork []VSwitchCloudNet  `json:"cloud_network,omitempty"`
}

// VSwitchServer represents a server attached to a vSwitch.
type VSwitchServer struct {
	ServerIP      string `json:"server_ip"`
	ServerIPv6Net string `json:"server_ipv6_net"`
	ServerNumber  int    `json:"server_number"`
	Status        string `json:"status"`
}

// VSwitchSubnet represents a subnet attached to a vSwitch.
type VSwitchSubnet struct {
	IP      string `json:"ip"`
	Mask    string `json:"mask"`
	Gateway string `json:"gateway"`
}

// VSwitchCloudNet represents a cloud network attached to a vSwitch.
type VSwitchCloudNet struct {
	ID      int    `json:"id"`
	IP      string `json:"ip"`
	Mask    string `json:"mask"`
	Gateway string `json:"gateway"`
}

// VSwitchCreateInput is the input for creating a vSwitch.
type VSwitchCreateInput struct {
	Name   string `json:"name"`
	VlanID int    `json:"vlan"`
}
