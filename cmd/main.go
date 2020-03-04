package main

import (
	"context"
	"flag"
	"os"

	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/mgutz/ansi"
	"github.com/pkg/errors"
)

var (
	confPath string
	statusOK = true
	planMode bool
)

func init() {

	// Parse flags
	flag.StringVar(&confPath, "conf", "configs/config.yaml", "The path of the configuration file")
	flag.BoolVar(&planMode, "plan", false, "Plan mode")
	flag.Parse()

	// Print information on current update
	fmt.Println(ansi.Color("BigQuery update information", "yellow+b"))

	// Print plan mode status
	fmt.Println(fmt.Sprintf("  Plan mode:            %t", planMode))

	// Print creation date
	fmt.Println(fmt.Sprintf("  Created at:           %s", time.Now().UTC()))

	// Print GCP credentials file
	credEnv, ok := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS")
	if ok {
		fmt.Println(fmt.Sprintf("  Credentials:          %s", credEnv))
	} else {
		fmt.Println("  Credentials:          default")
	}

	// Print configuration file

	if confPath == "config.yaml" {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(ansi.Color(fmt.Sprint(errors.Wrap(err, "cannot get current working directory")), "red"))
		}
		fmt.Print(fmt.Sprintf("  Configuration file:   %s/%s", dir, confPath))
	} else {
		fmt.Print(fmt.Sprintf("  Configuration file:   %s", confPath))
	}
}

func main() {

	if err := run(); err != nil {
		fmt.Println(ansi.Color(fmt.Sprintf("%v", err), "red"))
		statusOK = false
	}

	fmt.Println(ansi.Color("\n\nBigQuery update result", "yellow+b"))

	if statusOK {
		if planMode {
			fmt.Println("  Status:               " + ansi.Color("plan_only", "blue"))
		} else {
			fmt.Println("  Status:               " + ansi.Color("success", "green"))
		}
	} else {
		fmt.Println("  Status:               " + ansi.Color("failure", "red"))
		os.Exit(1)
	}
}

func run() error {

	var conf Config
	err := conf.LoadFromFile(confPath)
	if err != nil {
		return errors.Wrap(err, "\ncannot load configuration")
	}

	client, err := bigquery.NewClient(context.Background(), conf.Project)
	if err != nil {
		return errors.Wrap(err, "\ncannot create client")
	}

	err = updateAccessControl(client, conf)
	if err != nil {
		return errors.Wrap(err, "\ncannot update accesses")
	}

	return nil
}
