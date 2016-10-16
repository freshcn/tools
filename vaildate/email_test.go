package vaildate

import "testing"

func TestIsEmail(t *testing.T) {
	if !IsEmail("fdaffasdfa@yahoo.com.cn") {
		t.Error("测试没有通过")
	}

	if !IsEmail("dafasdf@qq.com") {
		t.Error("测试没有通过")
	}

	if !IsEmail("dasfad.-2123dsfjkljsadfkdjsf@sina.com") {
		t.Error("测试没有通过")
	}

	if IsEmail("afasdfasdfasdfs@dsf") {
		t.Error("测试没有通过")
	}
}
