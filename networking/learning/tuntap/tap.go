package main

import (
	"log"
	"os/exec"

	"github.com/songgao/packets/ethernet"
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
	cmd := exec.Command("ip", "addr", "add", "10.1.1.100/24", "dev", ifaceName)
	_ = cmd.Run()

	cmd = exec.Command("ip", "link", "set", "dev", ifaceName, "up")
	_ = cmd.Run()

	var frame ethernet.Frame
	for {
		frame.Resize(1500)
		n, err := iface.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		log.Printf("Received: % x\n", frame)
		log.Println("####################\tDecoded content\t####################")
		log.Printf("Dst: %s\n", frame.Destination())
		log.Printf("Src: %s\n", frame.Source())
		log.Printf("Ethertype: % x\n", frame.Ethertype())
		log.Printf("Payload: % x\n", frame.Payload())
	}
}
