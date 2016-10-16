package uuid

import "testing"
import "fmt"

func TestCreateUUID(t *testing.T) {
	data := CreateUUID()
	fmt.Println("create uuid is:", data)
	if data == "" {
		t.Error("create uuid error")
	}
}
