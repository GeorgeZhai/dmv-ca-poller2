package getcookies

import (
	"fmt"
	"testing"
)

//GetJSESSIONID will get the latest JSESSIONID
func TestGetJSESSIONID(t *testing.T) {
	s, e := GetJSESSIONID()
	if e != nil {
		t.Error(s, e)
	}
	fmt.Println(s)
}
