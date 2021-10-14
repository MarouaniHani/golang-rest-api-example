package serve

import (
	"fmt"
	"kuwait-test/svc/configs"
	"kuwait-test/svc/server"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	Cmd *cobra.Command

	argAddress     string
	argDsn         string
	argCORSHosts   string
)

func init() {
	Cmd = &cobra.Command{
		Use:   "serve",
		Short: "Connect to the storage and begin serving requests.",
		Long:  ``,
		Run: func(Cmd *cobra.Command, args []string) {
			if err := serve(Cmd, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}

	Cmd.Flags().StringVarP(&argAddress, "address", "a", ":8080", "address to listen on")
	Cmd.Flags().StringVar(&argDsn, "dsn", "mysql://bits_store_user:bits_store_pwd@localhost:3306/bits_store?parseTime=true", "db url")
	Cmd.Flags().StringVar(&argCORSHosts, "cors-hosts", "*", "cors hosts, separated by comma")

}

func serve(cmd *cobra.Command, args []string) error {
	svr, err := server.NewServer(&configs.Config{
		HostPort:          argAddress,
		Dsn:               argDsn,
		CORSHosts:         argCORSHosts,
	})
	if err != nil {
		return err
	}

	log.Fatalln(svr.Run())

	return nil
}
