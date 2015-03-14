package main

import (
//	"code.google.com/p/gopacket/layers"
    "code.google.com/p/gopacket/pcap"
	"code.google.com/p/gopacket"
	"code.google.com/p/gopacket/tcpassembly"
	"code.google.com/p/gopacket/tcpassembly/tcpreader"
	"fmt"
	"bufio"
//	"flag"
	"io"
	"log"
	"net/http"
//	"time"
)

//func handlePacket(p pcap.PacketSource){
//	fmt.Println(p)
//}


// httpStreamFactory implements tcpassembly.StreamFactory
type httpStreamFactory struct{}

// httpStream will handle the actual decoding of http requests.
type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}

func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hstream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}
	go hstream.run() // Important... we must guarantee that data from the reader stream is read.

	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &hstream.r
}

func (h *httpStream) run() {
	buf := bufio.NewReader(&h.r)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF {
			// We must read until we see an EOF... very important!
			return
		} else if err != nil {
			log.Println("Error reading stream", h.net, h.transport, ":", err)
		} else {
			bodyBytes := tcpreader.DiscardBytesToEOF(req.Body)
			req.Body.Close()
			log.Println("Received request from stream", h.net, h.transport, ":", req, "with", bodyBytes, "bytes in request body")
		}
	}
}

func main() {
    ifs, err := pcap.FindAllDevs()
    if err == nil {
            for _, ife := range ifs {
                fmt.Println(ife.Name, ife.Description)
            }
    }
		
	// Set up assembly
	//streamFactory := &httpStreamFactory{}
	//streamPool := tcpassembly.NewStreamPool(streamFactory)
	//assembler := tcpassembly.NewAssembler(streamPool)
	//ticker := time.Tick(time.Minute)
	if handle, err := pcap.OpenLive(ifs[1].Name, 1600, true, 0); err != nil {
		panic(err)
	} else if err := handle.SetBPFFilter("tcp and port 80 or 443"); err != nil {  // optional
		panic(err)
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			//select{

			//handlePacket(packet)  // Do something with a packet here.
			fmt.Println(packet);
			//if(packet.ApplicationLayer() != nil){
			//	fmt.Println("%x\n", packet.ApplicationLayer().Payload())
			//}
			//tcp := packet.TransportLayer().(*layers.TCP)
			//fmt.Println(tcp);
			//assembler.AssembleWithTimestamp(packet.NetworkLayer().NetworkFlow(), tcp, packet.Metadata().Timestamp)
		}
	}
}