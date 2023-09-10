package main

import (
	"compress/flate"
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
)

func extract() {
	logger.Infof("Attempting to extract \"%s\"", ipa)
	format := archiver.Zip{}
	directory = fileNameWithoutExtension(filepath.Base(ipa))

	if _, err := os.Stat(directory); err == nil {
		logger.Info("Detected previously extracted directory, cleaning it up...")

		err := os.RemoveAll(directory)
		if err != nil {
			logger.Fatalf("Failed to clean up previously extracted directory: %s", err)
			exit()
		}

		logger.Info("Previously extracted directory cleaned up. ")
	}

	err := format.Unarchive(ipa, directory)
	if err != nil {
		logger.Fatalf("Failed to extract %s: **%v**", ipa, err)
		os.Exit(1)
	}

	logger.Infof("Successfully extracted to \"%s\"", directory)
}

func archive() {
	logger.Infof("Attempting to archive \"%s\"", directory)

	format := archiver.Zip{ CompressionLevel: flate.BestCompression }
	zip := directory + ".zip"

	if _, err := os.Stat(zip); err == nil {
		logger.Info("Detected previous archive, cleaning it up...")

		err := os.Remove(zip)
		if err != nil {
			logger.Fatalf("Failed to clean up previous archive: %s", err)
			exit()
		}

		logger.Info("Previous archive cleaned up.")
	}

	logger.Infof("Archiving \"%s\" to \"%s\"", directory, zip)
	err := format.Archive([]string{filepath.Join(directory, "Payload")}, zip)
	if err != nil {
		logger.Fatalf("Failed to archive \"%s\": %v", zip, err)
		exit()
	}

	if _, err := os.Stat("Unbound.ipa"); err == nil {
		logger.Info("Detected previous Unbound IPA, cleaning it up...")

		err := os.Remove("Unbound.ipa")
		if err != nil {
			logger.Fatalf("Failed to clean up previous Unbound IPA: %s", err)
			exit()
		}

		logger.Info("Previous Unbound IPA cleaned up.")
	}

	err = os.Rename(zip, "Unbound.ipa")
	if err != nil {
		logger.Fatalf("Failed to rename \"%s\": %v", zip, err)
		exit()
	}

	logger.Infof("Successfully archived \"%s\" to \"Unbound.ipa\"", zip)
}