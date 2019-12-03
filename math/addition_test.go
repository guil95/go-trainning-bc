package math

import (
	"reflect"
	"testing"
)

func TestAddition_Calc(t *testing.T) {
	type fields struct {
		number1 float64
		number2 float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			"Test one plus one",
			fields{
				number1: 1,
				number2: 1,
			},
			2,
		},
		{
			"Test one plus zero",
			fields{
				number1: 1,
				number2: 0,
			},
			1,
		},
		{
			"Test negative",
			fields{
				number1: -1,
				number2: 2,
			},
			1,
		},
		{
			"Test negative addition",
			fields{
				number1: -1,
				number2: -1,
			},
			-2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Addition{
				number1: tt.fields.number1,
				number2: tt.fields.number2,
			}
			if got := a.Calc(); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAddition(t *testing.T) {
	type args struct {
		number1 float64
		number2 float64
	}
	tests := []struct {
		name string
		args args
		want *Addition
	}{
		{
			"Test addition instance",
			args{
				number1: 1,
				number2: 2,
			},
			NewAddition(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAddition(tt.args.number1, tt.args.number2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAddtion() = %v, want %v", got, tt.want)
			}
		})
	}
}
