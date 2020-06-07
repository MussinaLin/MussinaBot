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
		{
			name :"test3",
			args:args{availableBalance:1050, maxSignleOrderAmount:1000},
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

func TestGenOrders(t *testing.T) {
	type args struct {
		availableBalance     float64
		maxSignleOrderAmount float64
		left                 float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"test1",
			args:args{availableBalance:50,maxSignleOrderAmount:1000,left:0},
			want:1,
		},
		{
			name:"test2",
			args:args{availableBalance:2000,maxSignleOrderAmount:1000,left:0},
			want:2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenOrders(tt.args.availableBalance, tt.args.maxSignleOrderAmount, tt.args.left); got != tt.want {
				t.Errorf("GenOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}