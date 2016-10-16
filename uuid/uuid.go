package uuid

import (
	"strings"

	"github.com/satori/go.uuid"
)

// CreateUUID 创建一个uuid已经过滤了-的
func CreateUUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
