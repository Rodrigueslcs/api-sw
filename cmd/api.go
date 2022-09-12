package cmd

import (
	"api-sw/server"
	"api-sw/src/shared/providers/logger"
	"os/signal"

	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	brandName   string
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

	apiCmd.PersistentFlags().StringVarP(&brandName, "brand", "b", os.Getenv("BRAND"), "")
	apiCmd.PersistentFlags().StringVarP(&version, "version", "v", os.Getenv("VERSION"), "")
	apiCmd.PersistentFlags().StringVarP(&environment, "environment", "e", os.Getenv("ENV"), "")
	apiCmd.PersistentFlags().StringVarP(&port, "port", "p", "0000", "")

	rootCmd.AddCommand(apiCmd)
}

func run(cmd *cobra.Command, args []string) {
	log := logger.New()
	srv := server.New()

	if err := srv.Run(log); err != nil {
		panic(err)
	}

	chanExit := make(chan os.Signal, 1)
	signal.Notify(chanExit, os.Interrupt)
}
