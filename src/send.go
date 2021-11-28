package main

import (
	"encoding/binary"
	"log"
	"net"

	"gopkg.in/yaml.v3"
)

func SendPacket(conn net.Conn, command int, metadata HeaderValues, body []byte) error {
	// version
	version := make([]byte, 2)
	binary.BigEndian.PutUint16(version, ProtocolVersion)

	// header encode
	metadataBytes, err := yaml.Marshal(&metadata)
	if err != nil {
		return err
	}

	// header length
	headerLength := uint32(4 + len(metadataBytes) + 32)
	headerLengthBytes := make([]byte, 3)
	binary.BigEndian.PutUint32(headerLengthBytes, headerLength)

	// command
	commandBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(commandBytes, uint32(command))

	// checksum
	checksum := make([]byte, 32)
	for i := 0; i < 32; i++ {
		checksum[i] = 0
	}

	// body length send
	bodyLength := len(body) + 32
	bodyLengthBytes := make([]byte, 3)
	binary.BigEndian.PutUint32(bodyLengthBytes, uint32(bodyLength))

	// send everything
	// TODO put everything into one array
	conn.Write(version)
	conn.Write(headerLengthBytes)
	conn.Write(commandBytes)
	conn.Write(metadataBytes)
	conn.Write(checksum)
	conn.Write(bodyLengthBytes)
	conn.Write(body)
	conn.Write(checksum)

	log.Println("sent packet to", conn.RemoteAddr().String())

	return nil
}
