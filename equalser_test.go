package equalser_test

import (
	"testing"

	. "github.com/mattn/equalser"
)

type Foo struct {
	name string
}

func (f *Foo) Equals(v interface{}) bool {
	if vf, ok := v.(*Foo); ok {
		return f.name == vf.name
	}
	return false
}

type TB struct {
	testing.T
	failed bool
}

func (tb *TB) Error(args ...interface{}) {
	tb.failed = true
	panic(1)
}

func (tb *TB) Errorf(format string, args ...interface{}) {
	tb.failed = true
	panic(1)
}

func (tb *TB) Fail() {
	tb.failed = true
	panic(1)
}

func (tb *TB) FailNow() {
	tb.failed = true
}

func (tb *TB) Failed() bool {
	return tb.failed
}

func (tb *TB) Fatal(args ...interface{}) {
	tb.failed = true
	panic(1)
}

func (tb *TB) Fatalf(format string, args ...interface{}) {
	tb.failed = true
	panic(1)
}

func (tb *TB) Log(args ...interface{})                  {}
func (tb *TB) Logf(format string, args ...interface{})  {}
func (tb *TB) Skip(args ...interface{})                 {}
func (tb *TB) SkipNow()                                 {}
func (tb *TB) Skipf(format string, args ...interface{}) {}
func (tb *TB) Skipped() bool                            { return false }

func TestEqualser(t *testing.T) {
	ore1 := &Foo{"俺"}
	ore2 := &Foo{"俺"}

	AssertEquals(t, ore1, ore2, "俺ってお前ちゃうやん")

	ore3 := &Foo{"お前"}

	// dummy tester
	tb := &TB{}
	defer func() {
		recover()
	}()
	AssertEquals(tb, ore2, ore3, "俺ってお前ちゃうやん")
	if !tb.failed {
		t.Fatal("should be failed")
	}
}
