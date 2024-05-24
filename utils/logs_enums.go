package utils

const (
	LogTypeStdout = "stdout"
	LogTypeStderr = "stderr"
)

func GetLogTypeStdout() *string {
	logType := LogTypeStdout
	return &logType
}

func GetLogTypeStderr() *string {
	logType := LogTypeStderr
	return &logType
}
