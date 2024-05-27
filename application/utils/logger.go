package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"strings"
)

func SetLogger()  {
	fileLogger := GetFileLogger()
	multi := io.MultiWriter(os.Stdout, fileLogger)
	log.SetOutput(multi)

	log.SetReportCaller(true)
	formatter := &log.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// this function is required when you want to introduce your custom format.
			// In my case I wanted file and line to look like this `file="engine.go:141`
			// but f.File provides a full path along with the file name.
			// So in `formatFilePath()` function I just trimmet everything before the file name
			// and added a line number in the end
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	log.SetFormatter(formatter)
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func GetFileLogger() *os.File {
	nameFile := "log.txt"
	path := GetStoragePrivatePath()
	pathFolder := path+"/log"

	// Create log directory if not exist
	err := os.MkdirAll(pathFolder, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create log directory")
	}

	// Create log file if not exist
	filePathLogger := pathFolder + "/" + nameFile
	file, err := os.OpenFile(filePathLogger, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("log.txt not found, creating new log.txt")
		var err2 error
		file, err2 = os.Create(filePathLogger)
		if err2 != nil {
			fmt.Println("Failed to create log.txt")
		}
	}

	return file
}
