package main

import "os"

func commandExit(config *Config) error {
	os.Exit(0)
	return nil
}
