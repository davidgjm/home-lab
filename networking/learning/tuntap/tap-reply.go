package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/songgao/water"
)

const ifaceName = "tap0"

func main() {
	config := water.Config{
		DeviceType: water.TAP,
	}
	config.Name = ifaceName
	iface, err := water.New(config)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("ip", "addr", "add", "10.1.1.200/24", "dev", ifaceName)
	_ = cmd.Run()

	cmd = exec.Command("ip", "link", "set", "dev", ifaceName, "up")
	_ = cmd.Run()

	sourceMACAddr, _ := net.ParseMAC("00:00:00:00:00:01")
	sourceIPAddr := net.ParseIP("10.1.1.200")

	handle, err := pcap.OpenLive(ifaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		arpLayer := packet.Layer(layers.LayerTypeARP)
		if arpLayer != nil {
			arpPacket, _ := arpLayer.(*layers.ARP)
			if arpPacket.Operation == layers.ARPRequest {
				fmt.Println("ARP Operation:", arpPacket.Operation)
				fmt.Println("ARP Source Hardware Address:", arpPacket.SourceHwAddress)
				fmt.Println("ARP Source Protocol Address:", arpPacket.SourceProtAddress)
				fmt.Println("ARP Target Hardware Address:", arpPacket.DstHwAddress)
				fmt.Println("ARP Target Protocol Address:", arpPacket.DstProtAddress)
				fmt.Println("-----------------------------")

				arpReply := &layers.ARP{
					AddrType:          layers.LinkTypeEthernet,
					Protocol:          layers.EthernetTypeIPv4,
					HwAddressSize:     6,
					ProtAddressSize:   4,
					Operation:         layers.ARPReply,
					SourceHwAddress:   sourceMACAddr,
					SourceProtAddress: sourceIPAddr.To4(),
					DstHwAddress:      arpPacket.SourceHwAddress,
					DstProtAddress:    arpPacket.SourceProtAddress,
				}

				ethernetLayer := &layers.Ethernet{
					SrcMAC:       sourceMACAddr,
					DstMAC:       arpPacket.SourceHwAddress,
					EthernetType: layers.EthernetTypeARP,
				}

				frame1 := gopacket.NewSerializeBuffer()
				gopacket.SerializeLayers(frame1, gopacket.SerializeOptions{}, ethernetLayer, arpReply)
				_, err = iface.Write(frame1.Bytes())
			}
			continue
		}

		icmpLayer := packet.Layer(layers.LayerTypeICMPv4)
		if icmpLayer != nil {
			icmpPacket, _ := icmpLayer.(*layers.ICMPv4)

			fmt.Printf("ICMP Type: %d\n", icmpPacket.TypeCode.Type())
			fmt.Printf("ICMP Code: %d\n", icmpPacket.TypeCode.Code())
			fmt.Printf("ICMP Checksum: %d\n", icmpPacket.Checksum)
			fmt.Printf("ICMP Payload: %v\n", icmpPacket.Payload)
			fmt.Println("-----------------------------")

			if icmpPacket.TypeCode.String() == "EchoRequest" {
				icmpReplyPacket := &layers.ICMPv4{
					TypeCode: layers.ICMPv4TypeEchoReply,
					Id:       icmpPacket.Id,
					Seq:      icmpPacket.Seq,
				}

				ipPacket := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
				ipPacket.DstIP, ipPacket.SrcIP = ipPacket.SrcIP, ipPacket.DstIP
				fmt.Printf("ICMP SIP: %v\n", ipPacket.SrcIP)
				fmt.Printf("ICMP DIP: %v\n", ipPacket.DstIP)

				ethernetPacket := packet.Layer(layers.LayerTypeEthernet).(*layers.Ethernet)
				ethernetPacket.DstMAC, ethernetPacket.SrcMAC = ethernetPacket.SrcMAC, ethernetPacket.DstMAC

				frame2 := gopacket.NewSerializeBuffer()
				gopacket.SerializeLayers(frame2, gopacket.SerializeOptions{
					FixLengths: true,
					ComputeChecksums: true,
				}, ethernetPacket, ipPacket, icmpReplyPacket, gopacket.Payload(icmpPacket.Payload))

				_, err = iface.Write(frame2.Bytes())
			}
			continue
		}
	}
}
