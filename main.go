package main

import (
	"fmt"

	utils "github.com/pthomison/golang-utils"
	"github.com/spf13/cobra"
)

var (
	message string
	name    string

	rootCmd = &cobra.Command{
		Use:   "ssm-k8s-injector",
		Short: "ssm-k8s-injector",
		Run:   run,
	}
)

func main() {

	rootCmd.PersistentFlags().StringVarP(&message, "message", "m", "hello world", "message the program will output")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "patrick", "name the program will output to")

	err := rootCmd.Execute()

	utils.Check(err)
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println(message + ", " + name)
}
