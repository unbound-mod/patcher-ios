package main

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