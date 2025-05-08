package cmd

import (
	"log"
	"os"

	"github.com/gvcastellain/go-driver/pkg/requests"
	"github.com/spf13/cobra"
)

func authenticate() *cobra.Command {
	var (
		user string
		pass string
	)

	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Autentica usuario com API",
		Run: func(cmd *cobra.Command, args []string) {
			if user == "" || pass == "" {
				log.Println("usuario e senha obrigatorios")
				os.Exit(1)
			}

			err := requests.Auth("/auth", user, pass)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&user, "user", "u", "", "nome do usuario")
	cmd.Flags().StringVarP(&pass, "pass", "p", "", "senha do usuario")

	return cmd
}
