package main

import (
	"GoTipBot/Internal/CommManage"
	"GoTipBot/Internal/DataManage"
	"GoTipBot/Internal/FileManage"
	"GoTipBot/Structs"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Global instance of the current config
var config Structs.JSONData

func main() {

	// Get command-line JSON struct file
	fileStruct := os.Args[1]

	// Check if JSON struct file exists
	boolFileStruct, strFileStruct := FileManage.CheckFile(fileStruct)

	// Read struct file and assign it to current env.
	if boolFileStruct && DataManage.ReadStruct(fileStruct, &config) {

		// Create a channel to receive commands
		comms := make(chan string)

		// Start the bot in a separate goroutine
		go bot()

		// Start a goroutine to read commands
		go waitNewComm(comms)

		// Processing all the commands
		CommManage.ManageComms(comms, config)

	} else {
		fmt.Printf(strFileStruct)
	}
}

func bot() {
	for {
		// Add any desired actions or logic here
	}
}

func waitNewComm(comms chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		comm, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading command:", err)
			continue
		}

		comms <- strings.TrimSpace(comm)
	}
}
