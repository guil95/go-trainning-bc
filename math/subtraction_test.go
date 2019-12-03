package math

import (
	"reflect"
	"testing"
)

func TestNewSubtraction(t *testing.T) {
	type args struct {
		number1 float64
		number2 float64
	}
	tests := []struct {
		name string
		args args
		want *Subtraction
	}{
		{
			"Subtraction instance",
			args{
				1,
				2,
			},
			NewSubtraction(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSubtraction(tt.args.number1, tt.args.number2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubtraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtraction_Calc(t *testing.T) {
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
			"Test one minus one",
			fields{
				number1: 1,
				number2: 1,
			},
			0,
		},
		{
			"Test one minus zero",
			fields{
				number1: 1,
				number2: 0,
			},
			1,
		},
		{
			"Test negative",
			fields{
				number1: 1,
				number2: 2,
			},
			-1,
		},
		{
			"Test negative subtraction",
			fields{
				number1: -1,
				number2: 1,
			},
			-2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Subtraction{
				number1: tt.fields.number1,
				number2: tt.fields.number2,
			}
			if got := s.Calc(); got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
