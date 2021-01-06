package domain

import "os"

type (
	HostProvider int
	FlowType     int
)

const (
	GoogleCloudBucket HostProvider = iota
	SourceFlow        FlowType     = iota
	SinkFlow
)

type CloudStorage struct {
	HostProvider HostProvider
	ProjectId    string
	FlowType     FlowType
}

type Bucket struct {
	Name string `json:"name"`
}

type Buckets struct {
	Buckets []Bucket `json:"buckets"`
}

type LocalRepository interface {
	FileExists(fileName string) (fileExists bool, err error)
}

type BucketRepository interface {
	List(projectId string) (buckets []Bucket, err error)
	FileExists(fileName string) (fileExists bool, err error)
	DownloadFile(fileName string) (success bool, err error)
	UploadFile(fileName string) (success bool, err error)
}

type FileRepository interface {
	Store(file os.File)
	FindById(id int) os.File
}
