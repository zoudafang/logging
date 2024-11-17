// logging/logging.go
package logging

import (
	"fmt"
	"os"
)

// InitLogger 初始化日志文件
func InitLogger(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// LogKVOperation 记录一个 KV 操作
func LogKVOperation(logger *os.File, opType, key, value, location string) error {
	logMsg := fmt.Sprintf("%s %s %s %s\n", opType, key, value, location)
	_, err := logger.WriteString(logMsg)
	return err
}
