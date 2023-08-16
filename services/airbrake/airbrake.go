package airbrake

import (
	"os"
	"strconv"

	"github.com/airbrake/gobrake/v5"
	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"
)

var (
	Notifier *gobrake.Notifier
	log      logger.Logger
)

func init() {
	log = loadLoggerConfig()
	var err error
	if Notifier, err = loadAirbrakeConfig(); err != nil {
		log.Errorf("Skipping Airbrake config due to error: %s", err)
	}
}

func loadLoggerConfig() logger.Logger {
	var c = config.New()
	var loggerFactory = logger.NewFactory(c)
	return loggerFactory.NewLogger()
}

func loadAirbrakeConfig() (*gobrake.Notifier, error) {

	project_id, err := strconv.ParseInt(os.Getenv("AIRBRAKE_PROJECT_ID"), 0, 64)
	var project_key string = os.Getenv("AIRBRAKE_PROJECT_KEY")
	var env string = os.Getenv("GO_ENV")

	if err != nil {
		return nil, err
	}

	//Initialize and use gobrake with your account credentials:
	var notifier = gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
		ProjectId:   project_id,
		ProjectKey:  project_key,
		Environment: env,
	})

	// defer airbrake.NotifyOnPanic()

	return notifier, nil
}

// airbrake.Notify(errors.New("Hello Airbrake!"), nil)// Or the familiar
// Go pattern.if err != nil {	airbrake.Notify(err)}
