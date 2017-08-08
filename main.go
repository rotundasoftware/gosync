package main

import (
	"fmt"
	"os"

	"github.com/rotundasoftware/gosync/gosync"
	"github.com/rotundasoftware/gosync/version"

	log "github.com/cihub/seelog"
	"github.com/codegangsta/cli"
	"github.com/mitchellh/goamz/aws"
)

func main() {
	app := cli.NewApp()
	app.Name = "gosync"
	app.Usage = "gosync OPTIONS SOURCE TARGET"
	app.Version = version.Version()
	app.Flags = []cli.Flag{
		cli.IntFlag{Name: "concurrent, c", Value: 20, Usage: "number of concurrent transfers"},
		cli.StringFlag{Name: "log-level, l", Value: "info", Usage: "log level"},
		cli.StringFlag{Name: "aws-access-key-id", Value: "", Usage: "AWS Access Key Id"},
		cli.StringFlag{Name: "aws-access-key-secret", Value: "", Usage: "AWS Access Key Secret"},
		cli.StringFlag{Name: "aws-security-token", Value: "", Usage: "AWS Security Token"},
		cli.StringFlag{Name: "aws-region", Value: "", Usage: "AWS Region"},
		cli.StringFlag{Name: "aws-acl", Value: "private", Usage: "AWS ACL"},
	}

	app.Action = func(c *cli.Context) {
		defer log.Flush()
		setLogLevel(c.String("log-level"))

		err := validateArgs(c)
		exitOnError(err)

		key := c.String("aws-access-key-id")
		secret := c.String("aws-access-key-secret")
		token := c.String("aws-security-token")
		region := c.String("aws-region")
		acl := c.String("aws-acl")
		concurrent := c.Int("concurrent")

		auth, err := aws.GetAuth(key, secret)
		exitOnError(err)
		if token != "" {
			auth.Token = token
		}

		source := c.Args()[0]
		target := c.Args()[1]

		log.Infof("Setting source to '%s'.", source)
		log.Infof("Setting target to '%s'.", target)
		log.Infof("Setting concurrent transfers to '%d'.", concurrent)

		syncPair := gosync.NewSyncPair(auth, source, target, region, acl, concurrent)

		err = syncPair.Sync()
		exitOnError(err)

		log.Infof("Syncing completed successfully.")
	}
	app.Run(os.Args)
}

func validateArgs(c *cli.Context) error {
	if len(c.Args()) != 2 {
		return fmt.Errorf("Source and target required.")
	}
	return nil
}

func exitOnError(e error) {
	if e != nil {
		log.Errorf("Received error '%s'", e.Error())
		log.Flush()
		os.Exit(1)
	}
}

func setLogLevel(level string) {
	if level != "error" && level != "warn" {
		log.Infof("Setting log level '%s'.", level)
	}
	logConfig := fmt.Sprintf("<seelog minlevel='%s'>", level)
	logger, _ := log.LoggerFromConfigAsBytes([]byte(logConfig))
	log.ReplaceLogger(logger)
}
