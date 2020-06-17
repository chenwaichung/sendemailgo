package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

var (
	appName    = "sendemailgo"
	appVersion = "0.0.1"
)

var v = flag.Bool("v", false, "display the version and exit.")
var ver = flag.Bool("version", false, "display the version and exit.")
var h = flag.Bool("h", false, "display the usage and exit.")
var help = flag.Bool("help", false, "display the usage and exit.")
var enTLS = flag.Bool("tls", false, "enable the tls option.")
var server = flag.String("s", "", "the mail server address.")
var from = flag.String("f", "", "from (sender) email address.")
var port = flag.Int("p", 25, "the mail server port, default 25.")
var user = flag.String("xu", "", "username for SMTP authentication.")
var pass = flag.String("xp", "", "password for SMTP authentication.")
var receiver = flag.String("t", "", "to email address(receiver).")
var subject = flag.String("u", "", "message subject.")
var msg = flag.String("m", "", "message body.")

func main() {
	flag.Parse()

	if *v || *ver {
		fmt.Printf("%s %s", appName, appVersion)
		os.Exit(0)
	}

	if *h || *help {
		flag.Usage()
		os.Exit(0)
	}

	if *from == "" ||
		*server == "" ||
		*receiver == "" ||
		*msg == "" {
		fmt.Printf("params can not be empty.\n\n")
		flag.Usage()
		os.Exit(1)
	}

	d := gomail.NewDialer(*server, *port, *user, *pass)
	if *enTLS {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	m := gomail.NewMessage()
	m.SetHeader("From", *from)
	m.SetHeader("To", *receiver)
	m.SetHeader("Subject", *subject)
	m.SetBody("text/html", *msg)

	// Send the email to receiver
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Successfully sent '%s' to %s\n", *msg, *receiver)
}
