package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var conf config

func main() {
	flag.StringVar(&conf.password, "password", "", "password")
	flag.StringVar(&conf.password, "p", "", "password")
	flag.StringVar(&conf.passwordFile, "f", "", "Password from file")
	flag.BoolVar(&conf.isDescrypt, "d", false, "Decrypt mode")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
	}
	flag.Parse()

	password := preparePassword()
	input := read()

	if conf.isDescrypt {
		input, err := base64.StdEncoding.DecodeString(string(input))
		if err != nil {
			log.Fatalln("Input should be base64 encoded", err)
		}

		ct, err := Decrypt(input, password)
		if err != nil {
			log.Fatalln(err)
		}
		// Print the original data without endline
		fmt.Print(string(ct))
	} else {
		ct, err := Encrypt(input, password)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(base64.StdEncoding.EncodeToString(ct))
	}
}

func preparePassword() (p []byte) {
	if conf.password != "" {
		return []byte(conf.password)
	}
	if conf.passwordFile != "" {
		var e error
		if p, e = ioutil.ReadFile(conf.passwordFile); e != nil {
			log.Fatal(e)
		}
		return
	}

	if ev := os.Getenv("SIMCRYPT_PASSWORD"); ev != "" {
		return []byte(ev)
	}

	log.Fatalln("Please provide a password")
	return
}

func read() (d []byte) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			d = append(d, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		return
	}
	fmt.Print("Enter a data: ")
	fmt.Scanf("%s", &d)
	return
}

type config struct {
	password     string
	passwordFile string
	isDescrypt   bool
}

var usage = `Usage:  %s [OPTIONS]

A simple cli to encrypt/decrypt data.

Options:
  -d                    Decrypt mode
  -p, --password string Password for encrypt/decrypt
                        environment SIMCRYPT_PASSWORD
  -f <file>             Read password from a file

Example:
  export SIMCRYPT_PASSWORD='a strong password'
  echo 'test data' | simcrypt | tee /dev/stderr | simcrypt -d
`
