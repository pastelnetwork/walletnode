package cli

import (
	"context"
	"io/ioutil"

	"github.com/pastelnetwork/go-commons/cli"
	"github.com/pastelnetwork/go-commons/configer"
	"github.com/pastelnetwork/go-commons/errors"
	"github.com/pastelnetwork/go-commons/log"
	"github.com/pastelnetwork/go-commons/log/hooks"
	"github.com/pastelnetwork/go-commons/version"
	"github.com/pastelnetwork/walletnode/config"
	"github.com/pastelnetwork/walletnode/internal/common"
	"github.com/pastelnetwork/walletnode/internal/restserver"
	"github.com/pastelnetwork/walletnode/internal/ticket"
	"github.com/pastelnetwork/walletnode/pastel"
)

const (
	appName  = "walletnode"
	appUsage = "Pastel Wallet Node"

	defaultConfigFile = ""
)

// NewApp inits a new command line interface.
func NewApp() *cli.App {
	configFile := defaultConfigFile
	config := config.New()

	app := cli.NewApp(appName)
	app.SetUsage(appUsage)
	app.SetVersion(version.Version())

	app.AddFlags(
		// Main
		cli.NewFlag("config-file", &configFile).SetUsage("Set `path` to the config file.").SetValue(configFile).SetAliases("c"),
		cli.NewFlag("log-level", &config.LogLevel).SetUsage("Set the log `level`.").SetValue(config.LogLevel),
		cli.NewFlag("log-file", &config.LogFile).SetUsage("The log `file` to write to."),
		cli.NewFlag("quiet", &config.Quiet).SetUsage("Disallows log output to stdout.").SetAliases("q"),
		// Rest
		//cli.NewFlag("swagger", &config.Rest.Swagger).SetUsage("Enable Swagger UI."),
	)

	app.SetActionFunc(func(args []string) error {
		if configFile != "" {
			if err := configer.ParseFile(configFile, config); err != nil {
				return err
			}
		}

		if config.Quiet {
			log.SetOutput(ioutil.Discard)
		} else {
			log.SetOutput(app.Writer)
		}

		if config.LogFile != "" {
			fileHook := hooks.NewFileHook(config.LogFile)
			log.AddHook(fileHook)
		}

		if err := log.SetLevelName(config.LogLevel); err != nil {
			return errors.Errorf("--log-level %q, %s", config.LogLevel, err)
		}

		return run(config)
	})

	return app
}

func run(config *config.Config) error {
	log.Debug("[app] start")
	defer log.Debug("[app] end")

	log.Debugf("[app] config: %s", config)

	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	// sys.RegisterInterruptHandler(cancel, func() {
	// 	log.Info("[app] Interrupt signal received. Gracefully shutting down...")
	// })

	if err := pastel.Init(config.Pastel); err != nil {
		return err
	}

	internalApp := common.NewApplication(appUsage)
	restServer := restserver.New(pastel.Client, config.REST)

	ticketProc := ticket.NewTicketProc(pastel.Client)
	restServer.AddGetHandlers(map[string]interface{}{
		"/ws": ticketProc.RegisterArtTicket,
	})

	//p2pServer := fileserver.P2PServer{}

	internalApp.Run([]func(ctx context.Context, a *common.Application) func() error{
		// Start REST Server
		restServer.Start,
		// Start p2p Listener
		//p2pServer.Start,
	})

	// err := rest.New(config.Rest).Run(ctx)
	// return err
	return nil
}
