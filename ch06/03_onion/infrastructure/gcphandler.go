package infrastructure

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/pkg/errors"
	"github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/domain"
	"github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/interfaces"
	"github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/usecases"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/utils"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io"
	"io/ioutil"
	"os"
)

type GcpHandler struct {
	Client *storage.Client
}

var GcpInteractor *usecases.GcpInteractor

func NewGcpHandler(keyFile string) (gcpHandler *GcpHandler, err error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithServiceAccountFile(keyFile))
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new storage client")
	}
	gcpHandler = new(GcpHandler)
	gcpHandler.Client = client
	return
}

func (handler *GcpHandler) ListBuckets(flowType domain.FlowType, projectId string) (buckets []domain.Bucket, err error) {
	Debug.Printf("Running: ListBuckets(%v, %v)", flowType, projectId)
	client := handler.Client
	ctx := context.Background()
	it := client.Buckets(ctx, projectId)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "bucket iterator error")
		}
		buckets = append(buckets, domain.Bucket{battrs.Name})
	}
	return
}

func (handler *GcpHandler) FileExists(fileName string) (fileExists bool, err error) {
	ctx := context.Background()
	bucketName := Config.Gcp.Source.BucketName
	newFile := domain.NewFile(fileName)
	fullPath := newFile.FullHostPath(Config.Gcp.Source.Path)
	Debug.Printf("fullPath: %s", fullPath)
	br, err := handler.Client.Bucket(bucketName).Object(fullPath).NewReader(ctx)
	if err != nil {
		return false, errors.Wrapf(err, "bucket reader error for %s", fullPath)
	} else {
		data, err := ioutil.ReadAll(br)
		defer br.Close()
		if err != nil {
			return false, errors.Wrapf(err, "ioutil.ReadAll error for %s", fullPath)
		} else if len(data) == 0 {
			return false, errors.Wrapf(err, "file size must be greater than 0 for %s", fullPath)
		}
	}
	return true, err
}

func (handler *GcpHandler) GetBucketObject(flowType domain.FlowType, projectId string, bucketName string, fileName string) storage.ObjectHandle {
	client := handler.Client
	fileObject := client.Bucket(bucketName).Object(fileName)
	return *fileObject
}

func (handler *GcpHandler) DownloadFile(fileName string) (success bool, err error) {
	newFile := domain.NewFile(fileName)
	fullFilePath := newFile.FullHostPath(Config.Gcp.Source.Path)
	Debug.Printf("fullFilePath: %s", fullFilePath)
	ctx := context.Background()

	Debug.Printf("Config.Gcp.Source.ProjectId: %s", Config.Gcp.Source.ProjectId)
	Debug.Printf("Config.Gcp.Source.BucketName: %s", Config.Gcp.Source.BucketName)
	Debug.Printf("fullFilePath: %s", fullFilePath)

	bucketObject := handler.GetBucketObject(domain.SourceFlow, Config.Gcp.Source.ProjectId, Config.Gcp.Source.BucketName, fullFilePath)
	fr, err := bucketObject.NewReader(ctx)
	if err != nil {
		return false, errors.Wrapf(err, "unable to get file (%s) from bucket(%s)", fullFilePath, Config.Gcp.Source.BucketName)
	}
	defer fr.Close()
	fileBytes, err := ioutil.ReadAll(fr)
	if err != nil {
		return false, errors.Wrap(err, "ioutil.ReadAll failed")
	}
	logFiles, err := newFile.Parse(fileBytes)
	if err != nil {
		return false, errors.Wrap(err, "newFile.Parse failed")
	}
	success = true
	var logFileName string
	var cachedLogFiles []string
	for i, logFile := range *logFiles {
		logFileName = newFile.FullParsedFileName(i)
		Info.Println("encoding, caching and saving logFileName: " + logFileName)
		logFileJson, err := logFile.ToJson()
		if err != nil {
			Error.Printf("unable to encode logFileName (%s) - ERROR: %v", logFileName, err)
			break
		}
		cachedLogFiles = append(cachedLogFiles, logFileJson)
		logFile.Write(logFileName, logFileJson)
	}
	return
}

func (handler *GcpHandler) UploadFile(fileName string) (success bool, err error) {
	ctx := context.Background()
	newFile := domain.NewFile(fileName)
	newFullPath := newFile.FullLocalPath()
	f, err := os.Open(newFullPath)
	if err != nil {
		return false, errors.Wrapf(err, "unable to open local file: %s", newFullPath)
	}
	defer f.Close()
	bucketObject := handler.GetBucketObject(domain.SinkFlow, Config.Gcp.Sink.ProjectId, Config.Gcp.Sink.BucketName, newFile.FullHostPath(Config.Gcp.Sink.Path))
	wc := bucketObject.NewWriter(ctx)
	if _, err := io.Copy(wc, f); err != nil {
		return false, errors.Wrapf(err, "io.Copy failed for %s", newFullPath)
	}
	if err := wc.Close(); err != nil {
		return false, errors.Wrapf(err, "io.Copy failed for %s", newFullPath)
	}
	success = true
	return
}

func GetGcpInteractor() (gcpInteractor *usecases.GcpInteractor, err error) {
	if GcpInteractor == nil {
		sourceHandler, err := NewGcpHandler(Config.Gcp.Source.KeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create new source gcp handler")
		}
		sinkHandler, err := NewGcpHandler(Config.Gcp.Sink.KeyFile)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create new sink gcp handler")
		}
		handlers := make(map[string]interfaces.GcpHandler)
		handlers["SourceBucketRepo"] = sourceHandler
		handlers["SinkBucketRepo"] = sinkHandler
		gcpInteractor = new(usecases.GcpInteractor)
		gcpInteractor.SourceBucketRepository = interfaces.NewSourceBucketRepo(handlers)
		gcpInteractor.SinkBucketRepository = interfaces.NewSinkBucketRepo(handlers)
		GcpInteractor = gcpInteractor
	}
	return GcpInteractor, err
}
