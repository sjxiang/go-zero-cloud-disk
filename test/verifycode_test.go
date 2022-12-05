package test

import (
	"testing"

	"github.com/sjxiang/go-zero-cloud-disk/core/pkg/util"
)
func TestRandVerifyCode(t *testing.T) {
	t.Log(util.RandVerifyCode())
}