package utils

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"strconv"
)

type Conf struct {
	Port                     string `yaml:"port"`
	LogDebugInfo             bool   `yaml:"log-debug-info"`
	MaxConcurrentConnections int    `yaml:"max-concurrent-connections"`
	MaxNumber                int    `yaml:"max-number"`
	UseNumberHandler         bool   `yaml:"use-number-handler"`
}

var Config Conf

func GetOptions() bool {
	var configFile string

	flag.StringVar(&configFile, "config-file", "", "configuration file")
	flag.StringVar(&Config.Port, "port", "8080", "port that the API listens on")
	flag.BoolVar(&Config.LogDebugInfo, "log-debug-info", false, "whether to log debug output to the log (set to true for debug purpose)")
	flag.IntVar(&Config.MaxConcurrentConnections, "max-concurrent-connections", 6, "maximum number of concurrent connections (not currently used")
	flag.IntVar(&Config.MaxNumber, "max-number", 10, "maximum number that user is allowed to enter")
	flag.BoolVar(&Config.UseNumberHandler, "use-number-handler", false, "use number handler (to display number, optionally with FormatNumber applied) else display files in project root")
	flag.Parse()

	if configFile != "" {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			HandlePanic(errors.Wrap(err, "unable to read file"))
		}

		if err := yaml.Unmarshal(data, &Config); err != nil {
			HandlePanic(errors.Wrap(err, "unable to unmarshal from file"))
		}
	}
	return true
}

type Datastore interface{}

func UpdateConfig(d Datastore, key, val string) (oldValue string) {
	Debug.Printf("key (%s), val (%v)\n", key, val)
	value := reflect.ValueOf(d)
	if value.Kind() != reflect.Ptr {
		panic("not a pointer")
	}
	valElem := value.Elem()
	for i := 0; i < valElem.NumField(); i++ {
		tag := valElem.Type().Field(i).Tag
		field := valElem.Field(i)
		switch tag.Get("yaml") {
		case key:
			if fmt.Sprintf("%v", field.Kind()) == "int" {
				oldValue = strconv.FormatInt(field.Int(), 10)
				intVal, err := strconv.Atoi(val)
				if err != nil {
					fmt.Printf("could not parse int, key(%) val(%s)", key, val)
				} else {
					field.SetInt(int64(intVal))
				}
			} else if fmt.Sprintf("%v", field.Kind()) == "bool" {
				oldValue = strconv.FormatBool(field.Bool())
				b, err := strconv.ParseBool(val)
				if err != nil {
					fmt.Printf("could not parse bool, key(%s) val(%s)", key, val)
				} else {
					field.SetBool(b)
				}
			} else {
				oldValue = field.String()
				field.SetString(val)
			}
		}
	}
	return
}
