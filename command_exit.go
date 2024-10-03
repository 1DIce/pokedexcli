package main

import "os"

func commandExit(config *Config, arguments []string) error {
	os.Exit(0)
	return nil
}
