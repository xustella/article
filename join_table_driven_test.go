package article

import (
  "testing"
"fmt"
)

type testData struct {
  list []string
  want string
}

func TestJoinWithCommas(t *testing.T) {
  testCases:=[]testData{
    testData{[]string{}, ""},
    testData{[]string{"an apple"}, "an apple"},
    testData{[]string{"an apple","an orange"},"an apple and an orange"},
    testData{[]string{"an apple","an orange","a pear","a dog","a cat"},"an apple, an orange, a pear, a dog, and a cat"},
  }
  for _,test:=range testCases{
    got:=JoinWithCommas(test.list)
    if got!=test.want {
        t.Error(errString(test.list,got,test.want))
    }
  }
}
func errString(list []string, got string, want string) string{
  return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",list, got, want )
  }
