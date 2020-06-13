package Bitfinex

import (
	"testing"
)

func TestGenOrders(t *testing.T) {
	type args struct {
		availableBalance     float64
		maxSignleOrderAmount float64
		minLoan				 float64
		left                 float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"test1",
			args:args{availableBalance:50,maxSignleOrderAmount:1000,minLoan:50, left:0},
			want:1,
		},
		{
			name:"test2",
			args:args{availableBalance:2000,maxSignleOrderAmount:1000,minLoan:50,left:0},
			want:2,
		},
		{
			name:"test3",
			args:args{availableBalance:2030,maxSignleOrderAmount:1000,minLoan:50,left:0},
			want:2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenOrders(tt.args.availableBalance, tt.args.maxSignleOrderAmount,tt.args.minLoan, tt.args.left); len(*got) != tt.want {
				t.Errorf("GenOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubmitOrders(t *testing.T) {
	type args struct {
		orders *[]MussinaOrder
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name:"test1",
			args: args{&[]MussinaOrder{ MussinaOrder{
				Amount: 0.5,
				Rate:   0.05,
				Period: 2,
			} }},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}