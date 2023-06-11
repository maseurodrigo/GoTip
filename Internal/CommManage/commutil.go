package CommManage

import (
	"GoTipBot/Internal/API"
	"GoTipBot/Structs"

	"fmt"
	"strings"
)

func ManageComms(comms chan string, config Structs.JSONData) {

	for comm := range comms {
		// Process the command
		if strings.HasPrefix(comm, config.Trigger) {

			// Command splitted by spaces
			splitedComm := strings.Split(strings.Replace(comm, config.Trigger, "", 1), " ")

			// Exit command trigger
			if splitedComm[0] == "quit" || splitedComm[0] == "exit" {
				fmt.Println("Bye")
				break
			}

			// Exclude the comm trigger from the array
			remaining := splitedComm[1:]

			// Concatenate the remaining elements into a single string
			strJoin := strings.Join(remaining, " ")

			// Get response from api and print it
			fmt.Println(fmt.Sprintf("%s\n", API.NewApiCall(config.Key, fmt.Sprintf("%s %s %s: %s\n", config.CallStart, splitedComm[0], config.CallEnd, strJoin))))
		}
	}
}
