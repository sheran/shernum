package modules

import (
	"encoding/xml"
	"net"
)

type NmapRun struct {
	XMLName	xml.Name	`xml:"nmaprun"`
	Hosts	[]NmapHost 	`xml:"host"`
}

func (n *NmapRun) GetHosts() []net.IP{
	var hostList []net.IP
	for _, host := range n.Hosts{
		hostList = append(hostList, net.ParseIP(host.Address.Addr))
	}
	return hostList
}




type NmapHost struct {
	XMLName xml.Name `xml:"host"`
	Address NmapAddress `xml:"address"`
	Ports	NmapPorts	`xml:"ports"`
}

type NmapAddress struct {
	XMLName xml.Name `xml:"address"`
	Addr string `xml:"addr,attr"`
	AddrType string `xml:"addrtype,attr"`
}

type NmapPorts struct {
	XMLName xml.Name `xml:"ports"`
	Ports []NmapPort `xml:"port"`
}

type NmapPort struct {
	XMLName xml.Name	`xml:"port"`
	Protocol string `xml:"protocol,attr"`
	PortId	string	`xml:"portid,attr"`
	State	NmapState		`xml:"state"`
	Service NmapService 	`xml:"service"`
}

type NmapState struct {
	State 	string 	`xml:"state,attr"`
	Reason	string	`xml:"reason,attr"`
}

type NmapService struct {
	Name 	string 	`xml:"name,attr"`
}