package main

import (
	"fmt"

	"path/filepath"

	utils "github.com/pthomison/golang-utils"
	"github.com/spf13/cobra"
)

var (
	SSMparameter string
	K8Ssecret    string
	K8Snamespace string
	AWSRegion    string

	rootCmd = &cobra.Command{
		Use:   "ssm-k8s-injector",
		Short: "ssm-k8s-injector",
		Run:   run,
	}
)

func main() {

	rootCmd.PersistentFlags().StringVarP(&SSMparameter, "ssm-parameter", "p", "", "Parameter to inject")
	rootCmd.PersistentFlags().StringVarP(&AWSRegion, "aws-region", "r", "us-east-2", "")
	rootCmd.PersistentFlags().StringVarP(&K8Ssecret, "k8s-secret", "s", "", "Secret to inject into")
	rootCmd.PersistentFlags().StringVarP(&K8Snamespace, "k8s-namespace", "n", "default", "Namespace of the secret")

	rootCmd.MarkPersistentFlagRequired("ssm-parameter")
	rootCmd.MarkPersistentFlagRequired("k8s-secret")

	err := rootCmd.Execute()

	utils.Check(err)
}

func run(cmd *cobra.Command, args []string) {
	value, err := utils.AWSGetParameter(SSMparameter, AWSRegion)
	utils.Check(err)

	cs, err := utils.GetClientSet()
	utils.Check(err)

	sec := make(utils.Secret)

	sec[filepath.Base(SSMparameter)] = []byte(value)

	_, err = utils.ApplySecret(cs, K8Ssecret, K8Snamespace, sec)
	utils.Check(err)

	fmt.Println(value)
}
