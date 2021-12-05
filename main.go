package main

import (
	"path/filepath"

	utils "github.com/pthomison/golang-utils"
	"github.com/spf13/cobra"
)

var (
	SSMparameter string
	SecretName   string
	SecretKey    string
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
	rootCmd.PersistentFlags().StringVarP(&SecretName, "secret-name", "s", "", "Secret to inject into")
	rootCmd.PersistentFlags().StringVarP(&SecretKey, "secret-key", "k", "", "Secret to inject into")
	rootCmd.PersistentFlags().StringVarP(&K8Snamespace, "k8s-namespace", "n", "default", "Namespace of the secret")

	rootCmd.MarkPersistentFlagRequired("ssm-parameter")
	rootCmd.MarkPersistentFlagRequired("secret-name")

	err := rootCmd.Execute()

	utils.Check(err)
}

func run(cmd *cobra.Command, args []string) {
	value, err := utils.AWSGetParameter(SSMparameter, AWSRegion)
	utils.Check(err)

	cs, err := utils.GetClientSet()
	utils.Check(err)

	if SecretKey == "" {
		SecretKey = filepath.Base(SSMparameter)
	}

	sec := make(utils.Secret)
	sec[SecretKey] = []byte(value)

	_, err = utils.GetSecret(cs, SecretName, K8Snamespace)

	if err != nil {
		_, err = utils.ApplySecret(cs, SecretName, K8Snamespace, sec)
		utils.Check(err)
	} else {
		_, err = utils.UpdateSecret(cs, SecretName, K8Snamespace, sec)
		utils.Check(err)
	}
}
