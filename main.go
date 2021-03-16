package main

import (
	"fmt"
	"os"
	"time"

	terminal "github.com/Baozisoftware/qrcode-terminal-go"
	whatsapp "github.com/Rhymen/go-whatsapp"
)

func main() {

	var wac, _ = whatsapp.NewConn(5 * time.Second)

	qr := make(chan string)
	go func() {
		terminal := terminal.New()
		terminal.Get(<-qr).Print()
	}()

	session, err := wac.Login(qr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
	}
	fmt.Printf("login successful, session: %v\n", session)

	text := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: "51977727632@s.whatsapp.net",
		},
		Text: "Hello Whatsapp",
	}

	err2, _ := wac.Send(text)
	fmt.Println(err2)

}
