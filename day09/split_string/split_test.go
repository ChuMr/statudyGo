package split_string

import (
	"reflect"
	"testing"
)

// func TestSplict(t *testing.T) {
// 	ret := Split("asdasda", "d")
// 	want := []string{"as", "as", "a"}
// 	if !reflect.DeepEqual(want, ret) {
// 		t.Errorf("want:%v,got:%v", want, ret)
// 	}
// }

// func Test2Splict(t *testing.T) {
// 	ret := Split("abcdefg", "de")
// 	want := []string{"abc", "fg"}
// 	if !reflect.DeepEqual(ret, want) {
// 		t.Fatalf("want:%v,got:%v", want, ret)
// 	}
// }

func TestSplict(t *testing.T) {
	//go test -run=TestSplict/case_3
	// go test -cover 覆盖率
	// 测试函数覆盖率100%，代码覆盖率60%
	// go test -bench=Splict -benchmem
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	TestGroup := map[string]testCase{
		"case_1": testCase{"asdasda", "d", []string{"as", "as", "a"}},
		"case_2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": testCase{"abcdefg", "de", []string{"abc", "fg"}},
	}

	for name, tc := range TestGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v,got:%#v", tc.want, got)
			}
		})
	}
}

func BenchmarkSplict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}
