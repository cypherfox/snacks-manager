/*
Copyright © 2025 Lutz Behnke <lutz.behnke@gmx.de>
This file is part of {{ .appName }}
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/cypherfox/snacks-manager/api/oapi"
	"github.com/cypherfox/snacks-manager/internal/backend"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// backendCmd represents the backend command
var backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr := "0.0.0.0:24771"

		runListener(serverAddr)
	},
}

func runListener(addr string) {

	backend := backend.NewSnackBackEnd()

	strict := oapi.NewApiHandler(backend)

	r := mux.NewRouter()

	// get an `http.Handler` that we can use
	h := oapi.HandlerFromMuxWithBaseURL(strict, r, "")

	fmt.Printf("starting server and listening on port %s", addr)

	s := &http.Server{
		Handler: h,
		Addr:    addr,
	}

	err := s.ListenAndServe()
	fmt.Printf("\n ended server \n")
	if err != nil {
		if err != nil {
			fmt.Printf("with error: %s", err.Error())
		}
	}
}

func init() {
	rootCmd.AddCommand(backendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
