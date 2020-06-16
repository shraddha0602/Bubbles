package main

import (
	"fmt"
	"io"
	"log"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal("Error : %s", err)
	}
}

func WriteString(message string, write *io.Writer) error {
	_, err := write.Write([]byte(message))
	HandleError(err)
	return nil
}

func CommandLine(command interface{}) error {
	var write io.Writer
	var err error

	switch cmd := command.(type) {
	case ServerSend:
		err = WriteString(fmt.Sprintf("Send %v\n", cmd.Message), &write)
	case Name:
		err = WriteString(fmt.Sprintf("Name %v\n", cmd.UserName), &write)
	case UserMessage:
		err = WriteString(fmt.Sprintf("Message %v %v\n", cmd.UserName, cmd.Message), &write)
	default:
		err = WriteString(fmt.Sprintf("Invalid Command"), &write)
	}
	return err
}

func main() {

}
