package main

import (
    "bufio"
    "fmt"
    "gopkg.in/zeromq/goczmq.v4"
    "log"
    "os"
    "strings"
    "time"
)

func main() {
    ep := "tcp://34.71.52.251:40000"
    sample := goczmq.NewReqChanneler(ep)
    if sample == nil {
        log.Fatal("Failed to subscribe to endpoint: ", ep)
    }
    defer sample.Destroy()

    // Prompt the user for input
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter proof data (or press Enter to use default): ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    // Let the socket connect
    time.Sleep(5 * time.Second)

    // Prepare the proof data
    var proof [][]byte
    if input == "" {
        // Use default value if no input is provided
        proof = [][]byte{[]byte("proofblock"), []byte("timestamp :" + time.Now().String()), []byte("0x0000000000000000000")}
    } else {
        // Use user-provided input
        proof = [][]byte{[]byte("proofblock"), []byte(input), []byte("!!!!!")}
    }

    // Send the proof
    sample.SendChan <- proof
    fmt.Printf("Proof sent: %s\n", proof)

    // Receive the response
    resp := <-sample.RecvChan
    fmt.Printf("Response received: %s\n", resp)
}
