package article

import "testing"

//1. name of this file has to end with _test.go
//2. name of each tesing function has to begin with Test

func TestNoPhrase(t *testing.T) {
    list:=[]string{}
    want:=""
    got:=JoinWithCommas(list)
    if got!=want {
        t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",list, got, want )
    }
}

func TestOnePhrase(t *testing.T) {
    list:=[]string{"an apple"}
    want:="an apple"
    got:=JoinWithCommas(list)
    if got!=want {
        t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",list, got, want )
    }
}


func TestTwoPhrases(t *testing.T) {
    list:=[]string{"an apple","an orange"}
    want:="an apple and an orange"
    got:=JoinWithCommas(list)
    if got!=want {
        t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",list, got, want )
    }
}

func TestThreeOrMorePhrases(t *testing.T) {
    list:=[]string{"an apple","an orange","a pear","a dog","a cat"}
    want:="an apple, an orange, a pear, a dog, and a cat"
    got:=JoinWithCommas(list)
    if got!=want {
        t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"",list, got, want )
    }
}
