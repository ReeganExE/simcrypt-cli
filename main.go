package main

import (
	"bufio"
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

	input := read()
	password := preparePassword()

	if conf.isDescrypt {
		fmt.Println(decrypt(hashPassword(password), string(input)))
	} else {
		fmt.Println(encrypt(hashPassword(password), input))
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
	fmt.Print("Enter: ")
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
