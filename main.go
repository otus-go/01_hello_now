package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/beevik/ntp"
)

// for tests
var ntpTime = ntp.Time

func PrintNow(out io.Writer) error {
	res, err := ntpTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, res.Format("2006-01-02 15:04:05"))
	return nil
}

func main() {
	err := PrintNow(os.Stdout)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
