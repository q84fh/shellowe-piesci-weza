# Init

$ go mod init viper-demo;
$ go get github.com/spf13/viper;
$ go install github.com/spf13/cobra-cli@latest;
$ cobra-cli init --author "Wielki Szu" --license apache --viper
$ ls
cmd  go.mod  go.sum  LICENSE  main.go
$ go build
$ viper-demo
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

# Zaimplementuj program
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Test: %s\n", viper.Get("test"))
	},

# Flagi
	rootCmd.PersistentFlags().StringP("test", "t", "wartość domyślna", "to jest testowa flaga")
	viper.BindPFlag("test", rootCmd.PersistentFlags().Lookup("test"))
$ viper test
Test: wartość domyślna
$ viper-test -t "To jest wartość z CLI"
Test: To jest wartość z CLI

# Zmienne środowiskowe
Dodaj 
    viper.SetEnvPrefix("DEMO")

$ DEMO_TEST="To jest wartość ze środowiska" ./viper-test -t "To jest wartość z CLI" 
Test: To jest wartość z CLI

$ DEMO_TEST="To jest wartość ze środowiska" ./viper-test 
Test: To jest wartość ze środowiska

$ export DEMO_TEST="To jest wartość ze środowiska"
$  ./viper-test

# Plik konfiguracyjny
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

config.yaml:
   test: To jest wartość z pliku konfiguracyjnego

$ ./viper-test
Test: To jest wartość ze środowiska

$ unset DEMO_TEST
$ ./viper-test
Using config file: /home/q84fh/Projekty/Goats/shellowe-piesci-weza/viper-test/test.yaml
Test: To jest wartość z pliku konfiguracyjnego