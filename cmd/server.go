package cmd

import (
	"fmt"
	"net/http"
	"cobra/internal/routes"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: startServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

 func startServer(cmd *cobra.Command, args []string) {
	router := routes.NewRouter()
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on %s...\n", addr)
	err := http.ListenAndServe(addr, router)
	
	if err != nil {
		fmt.Println(err)
	}
}