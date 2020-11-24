package main

import (
	"encoding/gob"
	"fmt"
	"./Data"
	"net"
	"time"
)

func _serverPOST(s *[]Data.Data) {
	server, err := net.Listen("tcp", ":9998")
	if err != nil {
		fmt.Println("GET ERROR on create server", err)
	}
	defer server.Close()
	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Println("Listen server error",err)
			continue
		}
		go _handleClient(client, s)

	}
}

func _handleClient(c net.Conn, s *[]Data.Data  ) {
	var data Data.Data
	err := gob.NewDecoder(c).Decode(&data)
	if err != nil {
		fmt.Println("Handle client: ", err)
		return
	}
	*s = append(*s, data)
}

func _serverGET(s *[]Data.Data) {
	server, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error On create sever: ",err)
		return
	}
	defer server.Close()
	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Println("Listen Server Error: ", err)
			continue
		}
		//_handleClient(client, *s)
		err = gob.NewEncoder(client).Encode((*s)[0])
		if err != nil {
			fmt.Println("Serialization error ", err)
			continue
		}
		(*s) = _removeItem(*s)
	}
}



func _process(s *[]Data.Data) {
	for {
		for index, _ := range *s {
			if index < len(*s) {
				fmt.Printf("id: %d, value: %d \n", (*s)[index].Id, (*s)[index].Value)
				(*s)[index].Value++
				time.Sleep(time.Millisecond * 500)
			}
		}
	}
}

func _removeItem(s []Data.Data) []Data.Data {
	return append(s[:0], s[1:]...)
}

func main() {
	data := make([] Data.Data, 5)
	for i:=0; i<5; i++{
		data[i] = Data.Data{Id: i+1, Value: 0}
	}
	go _process(&data)
	go _serverGET(&data)
	go _serverPOST(&data)
	fmt.Scanln()
}