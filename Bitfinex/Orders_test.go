package Bitfinex

import (
	"testing"
)


func Test_calNumOfOrders(t *testing.T) {
	type args struct {
		availableBalance     float64
		maxSignleOrderAmount float64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name :"test1" ,
			args:args{availableBalance:51, maxSignleOrderAmount:1000},
			want:1,
		},
		{
			name :"test2",
			args:args{availableBalance:2000, maxSignleOrderAmount:1000},
			want:2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calNumOfOrders(tt.args.availableBalance, tt.args.maxSignleOrderAmount); got != tt.want {
				t.Errorf("calNumOfOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}