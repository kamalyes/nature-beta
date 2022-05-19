package example

import (
	"nature-beta/server/statistics"
	"reflect"
	"sync"
	"testing"
)

// TestPrintMap
func TestPrintMap(t *testing.T) {

	a := &sync.Map{}
	a.Store(200, 50)
	a.Store(100, 20)
	a.Store(500, 10)

	tt := map[string]struct {
		a      *sync.Map
		result string
	}{
		"test1": {a: a, result: "100:20;200:50;500:10"},
	}

	for _, value := range tt {
		str := statistics.PrintMap(value.a)
		if !reflect.DeepEqual(value.result, str) {
			t.Errorf("数据不一致 预期:%v 实际:%v", value.result, str)
		}
	}
}

func Test_PrintTop(t *testing.T) {
	type args struct {
		requestTimeList []uint64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil",
			args: args{
				requestTimeList: nil,
			},
		},
		{
			name: "one data",
			args: args{
				requestTimeList: []uint64{1 * 1e6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statistics.PrintTop(tt.args.requestTimeList)
		})
	}
}
