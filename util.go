package main

import (
	"os"
	"path/filepath"

	"howett.net/plist"
)

func fileNameWithoutExtension(name string) string {
	return name[:len(name) - len(filepath.Ext(name))]
}

func exit() {
	logger.Info("Cleaning up...")

	e := os.RemoveAll(directory)
	if e != nil {
		logger.Fatalf("Failed to clean up: %v", e)
	}

	logger.Info("Cleaned up.")

	os.Exit(0)
}

func loadInfo() () {
	if info != nil {
		return
	}

	path := filepath.Join(directory, "Payload", "Discord.app", "Info.plist")
	file, err := os.Open(path)

	if err != nil {
		logger.Fatal("Couldn't find Info.plist. Is the provided zip an IPA file?")
		exit()
	}

	decoder := plist.NewDecoder(file)
	if err := decoder.Decode(&info); err != nil {
		logger.Fatal("Couldn't find Info.plist. Is the provided zip an IPA file?")
		exit()
	}
}

func saveInfo() {
	path := filepath.Join(directory, "Payload", "Discord.app", "Info.plist")
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0600)

	if err != nil {
		logger.Fatalf("Failed to open Info.plist for saving: %s", err)
		exit()
	}

	logger.Infof("Saving Info.plist data...")
	encoder := plist.NewEncoder(file)
	err = encoder.Encode(info)

	if err != nil {
		logger.Fatalf("Failed to save Info.plist. %s", err)
		exit()
	}

	logger.Infof("Saved Info.plist data.")
}
