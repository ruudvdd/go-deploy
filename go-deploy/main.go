package main

import (
	"flag"
	"fmt"
	"go-deploy/config"
	"os"
)

func main() {
	configFile := ""
	help := false

	flag.StringVar(&configFile, "config", "./go-deploy.yml", "Path to the configuration file you want to use")
	flag.BoolVar(&help, "help", false, "Print this help message")
	flag.Parse()

	args := flag.Args()

	if help {
		printUsage()
		os.Exit(0)
	}

	if len(args) != 2 {
		fmt.Println("go-deploy expects 2 arguments:")
		printUsage()
		os.Exit(2)
	}

	cfg, err := config.NewConfig(configFile)

	if err != nil {
		fmt.Println(err.Error())
		printUsage()
		os.Exit(2)
	}

	projectName := args[0]
	envName := args[1]

	project, err := cfg.GetProject(projectName)

	if err != nil {
		fmt.Println(err.Error())
		printUsage()
		os.Exit(2)
	}

	env, err := project.GetEnvironment(envName)

	if err != nil {
		fmt.Println(err.Error())
		printUsage()
		os.Exit(2)
	}

	fmt.Print(*env)

	// TODO start deployment process
}

func printUsage() {
	fmt.Println("usage: go-deploy <project> <environment>")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Printf("  % -20s%s\n", "project", "The name of the project you want to deploy")
	fmt.Printf("  % -20s%s\n", "environment", "The environment of the project you want to deploy")
	fmt.Println()
	fmt.Println("Options:")
	flag.PrintDefaults()
}
