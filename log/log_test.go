package log

import "testing"

func TestLog(t *testing.T) {
	Info("this is info log")
	Error("this is error log")
	Infof("this is log %s", "1")
	Errorf("this is error log %s", "1")
}
