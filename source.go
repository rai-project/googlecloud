package googlecloud

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Unknwon/com"
	homedir "github.com/mitchellh/go-homedir"
)

func wellKnownFile() string {
	const f = "application_default_credentials.json"
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "gcloud", f)
	}
	home, err := homedir.Dir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".config", "gcloud", f)
}

func DefaultConfigurationSource() (string, error) {
	// First, try the environment variable.
	const envVar = "GOOGLE_APPLICATION_CREDENTIALS"
	if filename := os.Getenv(envVar); filename != "" && com.IsFile(filename) {
		return filename, nil
	}

	// Second, try a well-known file.
	if filename := wellKnownFile(); filename != "" && com.IsFile(filename) {
		return filename, nil
	}

	return "", errors.New("unable to find default google cloud credentials file")
}
