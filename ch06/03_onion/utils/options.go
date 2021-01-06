package utils

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	App struct {
		Profile     string `yaml:"env"`
		Root        string `yaml:"root"`
		DownloadDir string `yaml:"download-dir"`
		Port        string `yaml:"port"`
	} `yaml:"app"`

	Gcp struct {
		Source struct {
			Path       string `yaml:"path"`
			KeyFile    string `yaml:"key-file"`
			ProjectId  string `yaml:"project-id"`
			BucketName string `yaml:"bucket-name"`
		} `yaml:"source"`
		Sink struct {
			Path       string `yaml:"path"`
			KeyFile    string `yaml:"key-file"`
			ProjectId  string `yaml:"project-id"`
			BucketName string `yaml:"bucket-name"`
		}
	} `yaml:"gcp"`

	Log struct {
		TimeTrack        bool `yaml:"time-track"`
		FullStackTrace   bool `yaml:"full-stack-trace"`
		DebugInfo        bool `yaml:"debug-info"`
		DebugInfoForTest bool `yaml:"debug-info-for-test"`
	} `yaml:"log"`
}

var Config Conf

func GetOptions() bool {
	var configFile string

	flag.StringVar(&configFile, "config", "", "configuration file")
	flag.StringVar(&Config.App.Profile, "app-profile", "development", "Runtime Profile. Determines to run server scripts or expect env vars")
	flag.StringVar(&Config.App.Root, "app-root", "/Users/benjamin/Lab/github/siren-k/study-learning-functional-programming-in-go/ch06/03_onion", "Application Root Directory(Required). Must be absolute path")
	flag.StringVar(&Config.App.DownloadDir, "app-download-dir", "/Users/benjamin/Lab/github/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/download", "Where files are downloaded")
	flag.StringVar(&Config.App.Port, "app-port", "8080", "Port that the API listens on")

	flag.BoolVar(&Config.Log.TimeTrack, "log-time-track", true, "Enable or disable logging of utils/TimeTrack() (For benchmarking/debugging)")
	flag.BoolVar(&Config.Log.FullStackTrace, "log-full-stack-trace", false, "Print version information and exit")
	flag.BoolVar(&Config.Log.DebugInfo, "log-debug-info", false, "Whether to log debug output to the log (Set to true for debug purposes")
	flag.BoolVar(&Config.Log.DebugInfoForTest, "log-debug-info-for-test", true, "Whether to log debug output to the log when running tests (Set to true for debug purposes)")

	flag.Parse()

	if configFile != "" {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			HandlePanic(err)
		}

		if err := yaml.Unmarshal(data, &Config); err != nil {
			HandlePanic(err)
		}
	}
	return true
}
