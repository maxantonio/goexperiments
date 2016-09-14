package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var file *os.File
var err error
var path = "F:\\out.txt"

func init_file() {
	/*Following line check whether files exists or not.If file exists then it opens in append mode
	* else it creates new file in path variable.
	 */
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err = os.Create(path)
		checkError(err)
	} else {
		file, err = os.OpenFile(path, os.O_APPEND, 0666)
		checkError(err)
	}

}
func main() {

	//Initilizes the file creti
	init_file()
	//tcp port on which this server listens
	tcp_port := "8085"
	//This variable is used to count incoming messages to the server
	message_count := 0
	//This is used to store the start up time of the server.
	start_time := time.Now()
	//Following method listens for the tcp protocol and specified port ,in this case it is 8085
	ln, _ := net.Listen("tcp", ":"+tcp_port)
	//Following method accepts the connection from the client and returns the connection object.
	conn, _ := ln.Accept()
	// run loop forever
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//After receiving messgae it increments the message_count variable by 1 .
		message_count += 1
		//It calls a method to write the data into the file.
		write_into_file(message, message_count, start_time.String())
	}
	//defer is used to close file when program exits.
	defer file.Close()
}

func write_into_file(message string, message_count int, start_time string) {
	//fmt.Println("Message Received..." + start_time + "\t" + strconv.Itoa(message_count) + "\t" + message + "\n")
	//Following line writes into the file. i.e timestamp,message count,actual message
	_, err := file.WriteString(start_time + "\t" + strconv.Itoa(message_count) + "\t" + message)
	checkError(err)
	//file.Close()
}

/*
This function is for error checking.If there is any error thet throws an exception
*/
func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

}
