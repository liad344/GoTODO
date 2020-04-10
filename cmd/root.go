package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	rootCmd = &cobra.Command{
	Use:   "start",
	Short: "GoTODO is a simple service that helps you TODO stuff",
	Long: `GoTODO Parses TODOs on your program (on any language) and save them to different formats`,
	Run: todo,
	}
	notificationCmd = &cobra.Command{
		Use:                        "notification",
		Aliases:                    []string{"notifaction"},
		Short:                      "Will use os notification",
		Run:                        notification,
	}
	cfgPath string
	n string
)

func notification(cmd *cobra.Command, args []string) {

}

func todo(cmd *cobra.Command, args []string) {
	
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgPath, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.AddCommand(notificationCmd)
	//	rootCmd.AddCommand(Keep)
	//	rootCmd.AddCommand(Whatsapp..)


}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgPath != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgPath)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}