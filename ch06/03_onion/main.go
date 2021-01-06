package main

import (
	"github.com/pkg/errors"
	"github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/infrastructure"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/interfaces"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/utils"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	defaultFileName = "eventset1.jsonl"
)

var (
	fileName string
	wsh      WebserviceHandler
)

func init() {
	GetOptions()
	if Config.Log.DebugInfo {
		InitLog("trace-debug-log.txt", os.Stdout, os.Stdout, os.Stderr)
	} else {
		InitLog("trace-log.txt", ioutil.Discard, os.Stdout, os.Stderr)
	}

	fileName = os.Getenv("TEST_FILENAME")
	if len(fileName) == 0 {
		fileName = defaultFileName
	}
	Debug.Printf("application root directory: %s", PadRight(Config.App.Root, " ", 20))
	Debug.Printf("application profile: %s", PadRight(Config.App.Profile, " ", 20))
	Debug.Printf("loging debug info: %v", Config.Log.DebugInfo)
	HandlePanic(os.Chdir(Config.App.Root))
}

type endpoint struct {
	Api
	uriExample string
}

func printApiExample(url, uriExample string) {
	if len(uriExample) == 0 {
		Info.Printf("http://localhost:%s%s", Config.App.Port, url)
	} else {
		Info.Printf("http://localhost:%s%s?%s", Config.App.Port, url, uriExample)
	}
}

func main() {
	gcpi, err := infrastructure.GetGcpInteractor()
	HandlePanic(errors.Wrap(err, "unable to get gcp interactor"))
	li, err := infrastructure.GetLocalInteractor()
	HandlePanic(errors.Wrap(err, "unable to get local interactor"))

	wsh = WebserviceHandler{}
	wsh.GcpInteractor = gcpi
	wsh.LocalInteractor = li

	var endpoints = []endpoint{
		{Api{wsh.Health, "/health"}, ""},
		{Api{wsh.ListSourceBuckets, "/list-source-buckets"}, "projectId=" + Config.Gcp.Source.ProjectId},
		{Api{wsh.ListSinkBuckets, "/list-sink-buckets"}, "projectId=" + Config.Gcp.Sink.ProjectId},
		{Api{wsh.SourceFileExists, "/source-file-exists"}, "fileName=" + fileName},
		{Api{wsh.DownloadFile, "/download-file"}, "fileName=" + fileName},
		{Api{wsh.UploadFile, "/upload-file"}, "fileName=" + fileName},
		{Api{wsh.LocalFileExists, "/local-file-exists"}, "fileName=" + fileName},
	}
	Info.Println("Example API endpoints:")
	{
		for _, ep := range endpoints {
			http.HandleFunc(ep.Api.Url, ep.Api.Handler)
			printApiExample(ep.Api.Url, ep.uriExample)
		}
	}
	http.ListenAndServe(":"+Config.App.Port, nil)
}
