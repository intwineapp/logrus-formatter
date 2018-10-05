package logformatter

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const runtimeSkip = 2

// WithLineNumberFormatter struct for use with logrus
type WithLineNumberFormatter struct {
	logrus.TextFormatter
}

// Format entry
func (f *WithLineNumberFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	te, err := f.TextFormatter.Format(entry)
	if err != nil {
		return nil, err
	}
	fe := append(fileInfo(), te...)
	return fe, nil
}

// filInfo extract the fileInfo
func fileInfo() []byte {
	_, file, line, ok := runtime.Caller(runtimeSkip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	s := fmt.Sprintf("[%s:%d] ", file, line)
	return []byte(s)
}
