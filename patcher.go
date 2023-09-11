package main

import (
	"path/filepath"

	"github.com/mholt/archiver"
)

func setSupportedDevices() {
	logger.Info("Setting supported devices...")

	delete(info, "UISupportedDevices")

	logger.Info("Supported devices set.")
}

func setAppName() {
	logger.Info("Setting app name...")

	info["CFBundleName"] = "Unbound"
	info["CFBundleDisplayName"] = "Unbound"

	logger.Info("App name set.")
}

func setFileAccess() {
	logger.Info("Setting file access...")

	info["UISupportsDocumentBrowser"] = true
	info["UIFileSharingEnabled"] = true

	logger.Info("File access enabled.")
}

func setIcons() {
	logger.Info("Downloading app icons...")

	icons := filepath.Join(assets, "icons.zip");
	download("https://assets.unbound.rip/icons/ios.zip", icons)

	logger.Info("Downloaded app icons.")


	logger.Info("Applying app icons...")

	info["CFBundleIcons"].(map[string]interface{})["CFBundlePrimaryIcon"].(map[string]interface{})["CFBundleIconName"] = "UnboundIcon"
	info["CFBundleIcons"].(map[string]interface{})["CFBundlePrimaryIcon"].(map[string]interface{})["CFBundleIconFiles"] = []string{"UnboundIcon60x60"}
	info["CFBundleIcons~ipad"].(map[string]interface{})["CFBundlePrimaryIcon"].(map[string]interface{})["CFBundleIconName"] = "UnboundIcon"
	info["CFBundleIcons~ipad"].(map[string]interface{})["CFBundlePrimaryIcon"].(map[string]interface{})["CFBundleIconFiles"] = []string{"UnboundIcon60x60", "UnboundIcon76x76"}

	zip := archiver.Zip{OverwriteExisting: true}
	discord := filepath.Join(directory, "Payload", "Discord.app")

	if err := zip.Unarchive(icons, discord); err == nil {
		logger.Info("Applied app icons.")
	} else {
		logger.Errorf("Failed to apply app icons: %v", err)
		exit()
	}
}