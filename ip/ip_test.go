package ip

import (
	"fmt"
	"testing"
)

// TestIP 测试ToLong和LongToIP是否可以正常的工作
func TestIP(t *testing.T) {
	ipList := []string{
		"202.118.1.81",
		"213.136.0.252",
		"129.70.132.37",
	}

	var (
		IPInt uint32
		ok    bool
		IPStr string
	)
	for _, IP := range ipList {
		IPInt, ok = ToLong(IP)
		if !ok {
			t.Errorf("ip:%s conversion fail", IP)
		}
		IPStr = LongToIP(IPInt)
		t.Logf("IP:%s \t Uint32:%d \t BackToIP:%s \t Status:%t\n", IP, IPInt, IPStr, IP == IPStr)
	}
}

func ExampleToLong() {
	ipList := []string{
		"202.118.1.81",
		"213.136.0.252",
		"129.70.132.37",
	}

	var (
		IPInt uint32
		ok    bool
		IPStr string
	)
	for _, IP := range ipList {
		IPInt, ok = ToLong(IP)
		if !ok {
			fmt.Printf("ip:%s conversion fail", IP)
			return
		}
		IPStr = LongToIP(IPInt)
		fmt.Printf("IP:%s \t Uint32:%d \t BackToIP:%s \t Status:%t\n", IP, IPInt, IPStr, IP == IPStr)
	}
}
