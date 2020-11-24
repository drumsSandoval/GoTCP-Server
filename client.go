package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"./Data"
	"time"
)

func GET_PROCESS(c net.Conn, data *Data.Data) {
	err := gob.NewDecoder(c).Decode(data)
	if err != nil {
		fmt.Println("ERROR GET PROCESS", err)
		return
	}
	c.Close()
}

func POST_VALUE(data *Data.Data) {
	c, err := net.Dial("tcp", ":9998")
	if err != nil {
		fmt.Println("Client Error: ", err)
		return
	}
	err = gob.NewEncoder(c).Encode(data)
	if err != nil {
		fmt.Println("POST err", err)
	}
	c.Close()
}

func _clientProcess(data *Data.Data) {
	for {
		fmt.Printf("id: %d, value: %d \n",data.Id, data.Value)
		data.Value++
		time.Sleep(time.Millisecond*500)
	}
}


func main() {
	c, err := net.Dial("tcp", ":9999")
	var data Data.Data
	if err != nil {
		fmt.Println("Client Error: ", err)
		return
	}
	defer POST_VALUE(&data)
	GET_PROCESS(c,&data)
	go _clientProcess(&data)
	fmt.Scanln()
}