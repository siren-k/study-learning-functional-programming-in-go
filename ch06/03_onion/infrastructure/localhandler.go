package infrastructure

import (
	"fmt"
	"github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/interfaces"
	"github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/usecases"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch06/03_onion/utils"
	"os"
)

type LocalHandler struct{}

var LocalInteractor *usecases.LocalInteractor

func NewLocalHandler() *LocalHandler {
	gcpHandler := new(LocalHandler)
	return gcpHandler
}

func (handler *LocalHandler) FileExists(fileName string) (fileExists bool, err error) {
	_, err = os.Stat(fmt.Sprintf("%s/%s", Config.App.DownloadDir))
	if !os.IsNotExist(err) {
		fileExists = true
	}
	return
}

func GetLocalInteractor() (localInteractor *usecases.LocalInteractor, err error) {
	if LocalInteractor == nil {
		localHandler := NewLocalHandler()
		localHandlers := make(map[string]interfaces.LocalHandler)
		localHandlers["LocalFileSystemRepo"] = localHandler
		localInteractor = new(usecases.LocalInteractor)
		localInteractor.LocalRepository = interfaces.NewLocalRepo(localHandlers)
		LocalInteractor = localInteractor
	}
	return LocalInteractor, nil
}
