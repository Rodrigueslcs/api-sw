package cmd

import (
	"api-sw/src/server"
	"api-sw/src/server/config"
	"api-sw/src/shared/providers/logger"
	"os/signal"

	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	environment string
	version     string
	port        string
	apiCmd      = &cobra.Command{
		Use:   "api",
		Short: "",
		Long:  "",
		Run:   run,
	}
)

func init() {
	godotenv.Load()

	apiCmd.PersistentFlags().StringVarP(&version, "version", "v", os.Getenv("VERSION"), "")
	apiCmd.PersistentFlags().StringVarP(&environment, "environment", "e", os.Getenv("ENV"), "")

	rootCmd.AddCommand(apiCmd)
}

func run(cmd *cobra.Command, args []string) {
	log := logger.New()
	srv := server.New()

	cfg, err := config.ReadConfigFromEnv(environment, version)
	if err != nil {
		panic(err)
	}

	if err := srv.Run(cfg, log); err != nil {
		panic(err)
	}

	chanExit := make(chan os.Signal, 1)
	signal.Notify(chanExit, os.Interrupt)
	<-chanExit

}
