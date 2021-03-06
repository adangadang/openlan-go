package schema

type Route struct {
	Prefix  string `json:"prefix"`
	Nexthop string `json:"nexthop"`
}

type Network struct {
	Name    string  `json:"name"`
	IpAddr  string  `json:"ipAddr"`
	IpRange int     `json:"ipRange"`
	Netmask string  `json:"netmask"`
	Routes  []Route `json:"routes"`
}
