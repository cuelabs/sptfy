package cmd

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func init() {
	log.Debug("setting up")

	// read in the command line

	// read in configuration variables from the environment

	// read in inputs
}



func main() {

	log.Debugf("ready to the start streaming music at %s", time.Now().UTC().Format(time.RFC3339))

}
