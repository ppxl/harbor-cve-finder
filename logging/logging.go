package logging

import "github.com/sirupsen/logrus"

type Logger struct{}

func (l Logger) Errorf(s string, err error) {

}

func GetInstance() *Logger {
	return &Logger{}
}

func Init(logrus.Level) error {
	return nil
}
