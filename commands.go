package gopen

import (
	"os/exec"
)

func Redox(url string) error {
	command := "launcher"
	return exec.Command(command, url).Start()
}

func Windows(url string) error {
	return exec.Command("cmd", "/c", "start", url).Start()
}

func Unix(url string) error {
	var browsers = []string{
		"wslview",          // Windows Subsystem for Linux
		"xdg-open",         // Generic X11
		"sensible-browser", // Debian/Ubuntu
		"firefox",          // Mozilla Firefox
		"google-chrome",    // Google Chrome
		"chromium",         // Chromium
		"chromium-browser", // Chromium (alternative name)
		"safari",           // Safari
		"opera",            // Opera
		"epiphany",         // GNOME Web
	}

	for _, browser := range browsers {
		cmd := exec.Command(browser, url)
		if err := cmd.Start(); err == nil {
			return nil
		}
	}

	return ErrOpenFailed
}

func Darwin(url string) error {
	return exec.Command("/usr/bin/open", url).Start()
}

func Haiku(url string) error {
	return exec.Command("/bin/open", url).Start()
}

func IOS(url string) error {
	return exec.Command("uiopen", "--url", url).Start()
}

func Android(url string) error {
	return exec.Command("am", "start", "--user", "0", "-a", "android.intent.action.VIEW", "-d", url).Start()
}
