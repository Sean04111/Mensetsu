package util

import (
	"strconv"
	"testing"
)

func TestAtoi(t *testing.T) {
	s:="-109090"
	stdAns,_:=strconv.Atoi(s)
	TestAns:=atoi(s)
	if TestAns!=stdAns{
		t.Errorf("answer dismiss ; instance : %v",s)
	}
}
