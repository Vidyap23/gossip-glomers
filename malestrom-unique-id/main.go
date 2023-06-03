package main

import (
    "encoding/json"
    "log"
	"time"
	"fmt"
	
    maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main(){
	node := maelstrom.NewNode()
	node.Handle("generate", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
	     
		// Update the message type to return back.
		currentTime := time.Now()
		str := fmt.Sprintf("%v", body["msg_id"])
		msg_id := str + currentTime.Format("2006-01-02 15:04:05.000000")
		body["type"] = "generate_ok"
		body["id"] = msg_id
	
		// Echo the original message back with the updated message type.
		return node.Reply(msg, body)
	})
	if err := node.Run(); err != nil {
		log.Fatal(err)
	}
}