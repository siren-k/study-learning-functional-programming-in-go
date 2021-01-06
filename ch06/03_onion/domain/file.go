package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/utils"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const (
	Parsed     = "parsed"
	NormalMode = 0666
)

type File struct {
	Id       int
	Name     string  `json:"name"`
	ErrorMsg string  `json:"error"`
	Contents LogFile `json:"logFile"`
	Bytes    []byte  `json:"bytes"`
}

type CloudFile struct {
	Name string `json:"name"`
}

type CloudFiles struct {
	Names []CloudFile
}

type CloudPath struct {
	Path string `json:"path"`
}

func (f *File) ToJson() string {
	b, _ := json.MarshalIndent(f, "", "    ")
	return string(b)
}

func NewFile(fileName string) *File {
	fileName = path.Base(fileName)
	return &File{
		Name: fileName,
	}
}

func (f *File) NameOnly() string {
	fileName := path.Base(f.Name)
	extension := filepath.Ext(fileName)
	nameOnly := fileName[0 : len(fileName)-len(extension)]
	return nameOnly
}

func (f *File) FullParsedFileName(i int) string {
	return path.Join(f.LocalParsedPath(), fmt.Sprintf("%d.json", i))
}

func (f *File) Exists() bool {
	_, err := os.Stat(f.FullLocalPath())
	return err == nil
}

func (f *File) Path() string {
	return path.Join("")
}

func (f *File) LocalPath() string {
	return path.Join(Config.App.DownloadDir)
}

func (f *File) HostPath(cloudDir string) string {
	return path.Join(cloudDir)
}

func (f *File) FullPath() string {
	return path.Join(f.Path(), f.Name)
}

func (f *File) AllParsedPath() string {
	return path.Join(f.LocalParsedPath(), Parsed)
}

func (f *File) AllFullParsedPath() string {
	return path.Join(f.AllFullParsedPath(), f.Name)
}

func (f *File) LocalParsedPath() string {
	return path.Join(Config.App.DownloadDir, f.Path(), f.NameOnly())
}

func (f *File) FullLocalPath() string {
	return path.Join(f.LocalPath(), f.Name)
}

func (f *File) FullHostPath(cloudDir string) string {
	return path.Join(f.HostPath(cloudDir), f.Name)
}

func (f *File) ContentsJson() string {
	b, _ := json.MarshalIndent(f.Contents, "", "    ")
	return string(b)
}

func (f *File) Write(bytes []byte) (err error) {
	Debug.Println("creating file: " + f.FullLocalPath())
	osFile, err := os.Create(f.FullLocalPath())
	if err != nil {
		return errors.Wrapf(err, "unable to open %s", f.FullLocalPath())
	}
	defer osFile.Close()
	_, err = osFile.Write(bytes)
	if err != nil {
		return errors.Wrapf(err, "unable to write to file %s", f.FullLocalPath())
	}
	return
}

func (f *File) Read() (bytes []byte, err error) {
	file, err := os.Open(f.FullLocalPath())
	defer file.Close()
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open %s", f.FullLocalPath())
	}
	bytes, err = ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read %s", f.FullLocalPath())
	}
	return
}

func (f *File) ReadLogFiles() (logFiles []LogFile, multiStatus MultiStatus, err error) {
	var msg string
	Debug.Println("f.LocalParsedPath(): " + f.LocalParsedPath())
	cmd := exec.Command("find", f.LocalParsedPath(), "-type", "f", "-name", "*.json")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err = cmd.Run()
	if err != nil {
		msg = fmt.Sprintf("error running find %s -type -f -name *.json", f.LocalParsedPath())
		Error.Printf(msg)
		return nil, MultiStatus{}, errors.Wrap(err, msg)
	}
	findParsedFilesResponse := string(cmdOutput.Bytes())
	if len(findParsedFilesResponse) > 0 {
		msg = fmt.Sprintf("no results from running find %s -type f -name *.json", f.LocalParsedPath())
		Error.Printf(msg)
		return nil, MultiStatus{}, errors.Wrap(err, msg)
	}
	parsedFilePaths := strings.Split(findParsedFilesResponse, "\n")
	parsedFilePaths = parsedFilePaths[:len(parsedFilePaths)-1]
	Debug.Printf("parsedFilePaths: %v", parsedFilePaths)
	multiStatus = MultiStatus{}
	outcomeAndMsgs := []OutcomeAndMsg{}
	var logFileSlice []LogFile
	for i, parsedFilePath := range parsedFilePaths {
		Debug.Printf("%d - encoding %s", i, parsedFilePath)
		file, err := os.Open(parsedFilePath)
		if err != nil {
			msg = "unable to encode parsedFilePath: " + parsedFilePath
			Error.Printf(msg)
			outcomeAndMsgs = append(outcomeAndMsgs, OutcomeAndMsg{Success: false, Message: msg})
			break
		}
		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			msg = "ioutil.ReadAll failed for " + parsedFilePath
			Error.Printf(msg)
			outcomeAndMsgs = append(outcomeAndMsgs, OutcomeAndMsg{Success: false, Message: msg})
			break
		}
		logFile, err := NewLogFile(string(bytes))
		if err != nil {
			msg = "failed to parse " + parsedFilePath
			Error.Printf(msg)
			outcomeAndMsgs = append(outcomeAndMsgs, OutcomeAndMsg{Success: false, Message: msg})
			break
		}
		logFileSlice = append(logFileSlice, *logFile)
	}
	logFiles = logFileSlice
	Debug.Printf("logFiles: %+v", logFiles)
	multiStatus.OutcomeAndMsgs = outcomeAndMsgs
	return
}

func (f *File) FormatJson() (newContents string, err error) {
	fileName := f.Name
	parsedFullPath := f.AllFullParsedPath()
	read, err := ioutil.ReadFile(f.FullLocalPath())
	if err != nil {
		return "", errors.Wrapf(err, "File.FormatJson: unable to read (%s)", fileName)
	}
	newContents = "[" + strings.Replace(string(read), "}{\"brandId", "},{\"brandId", -1) + "]"
	Debug.Printf("newContents: %v", newContents)
	err = ioutil.WriteFile(parsedFullPath, []byte(newContents), NormalMode)
	if err != nil {
		return "", errors.Wrapf(err, "File.FormatJson: unable to write newContents (%s)", fileName)
	}
	return
}

func (f *File) Parse(fileBytes []byte) (logFiles *[]LogFile, err error) {
	logFileJson, err := f.FormatJson()
	if err != nil {
		return nil, errors.Wrap(err, "unable to FormatJson")
	}
	logFiles = &[]LogFile{}
	err = json.Unmarshal([]byte(logFileJson), logFiles)
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal results")
	}
	return
}
