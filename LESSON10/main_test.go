package main_test

import (
	. "github.com/smart48ru/Golang-Level1/LESSON10"
	"math"
	"testing"
)

func TestValue_Addition(t *testing.T) {
	var (
		testTable = []struct {
			name string
			args *Value
			want float64
		}{
			// Первый сценарий
			{
				name: "2+2",
				args: &Value{2, 2},
				want: 4,
			},
			// Второй сценарий
			{
				name: "10.22+10.87",
				args: &Value{10.22, 10.87},
				want: 21.09,
			},
			// Третий сценарий
			{
				name: "250.22234+340.88844",
				args: &Value{250.22234, 340.88844},
				want: 591.11078,
			},
			//Четвертый сценарий
			{
				name: "465830+874639",
				args: &Value{465830, 874639},
				want: 1340469,
			},
		}
	)
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if AreaCalculated.Addition(testCase.args) != testCase.want {
				t.Errorf("Addition() = %v, want %v", AreaCalculated.Addition(testCase.args), testCase.want)
			}
		})
	}
}

func TestValue_Multiplication(t *testing.T) {
	var (
		testTable = []struct {
			name string
			args *Value
			want float64
		}{
			// Первый сценарий
			{
				name: "2.984*2",
				args: &Value{2.984, 2},
				want: 5.968,
			},
			// Второй сценарий
			{
				name: "10*10",
				args: &Value{10, 10},
				want: 100,
			},
			// Третий сценарий
			{
				name: "44.233780567384*66",
				args: &Value{44.233780567384, 66},
				want: 2919.429517447344,
			},
			//Четвертый сценарий
			{
				name: "465830*874639",
				args: &Value{465830, 874639},
				want: 407433085370,
			},
		}
	)
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if AreaCalculated.Multiplication(testCase.args) != testCase.want {
				t.Errorf("Multiplication() = %v, want %v", AreaCalculated.Addition(testCase.args), testCase.want)
			}
		})
	}
}

func TestValue_Subtraction(t *testing.T) {
	var (
		testTable = []struct {
			name string
			args *Value
			want float64
		}{
			// Первый сценарий
			{
				name: "2.984-5",
				args: &Value{2.984, 5},
				want: -2.016,
			},
			// Второй сценарий
			{
				name: "-10-220",
				args: &Value{-10, 220},
				want: -230,
			},
			// Третий сценарий
			{
				name: "44.123457775-66",
				args: &Value{44.123457775, 66},
				want: -21.876542225,
			},
			//Четвертый сценарий
			{
				name: "465830-874639",
				args: &Value{465830, 874639},
				want: -408809,
			},
		}
	)
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if AreaCalculated.Subtraction(testCase.args) != testCase.want {
				t.Errorf("Subtraction() = %v, want %v", AreaCalculated.Addition(testCase.args), testCase.want)
			}
		})
	}
}

func TestValue_Division(t *testing.T) {
	var (
		testTable = []struct {
			name string
			args *Value
			want float64
		}{
			// Первый сценарий
			{
				name: "55.6687/5",
				args: &Value{55.6687, 5},
				want: 11.13374,
			},
			// Второй сценарий
			{
				name: "-107853.22/220",
				args: &Value{-107853.22, 220},
				want: -490.241909090909091,
			},
			// Третий сценарий
			{
				name: "22.22/2",
				args: &Value{22.22, 2},
				want: 11.11,
			},
			//Четвертый сценарий
			{
				name: "88476387348635/666326228364283",
				args: &Value{88476387348635, 66632},
				want: 1327836285.097775843438588,
			},
		}
	)
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if AreaCalculated.Division(testCase.args) != testCase.want {
				t.Errorf("Division() = %v, want %v", AreaCalculated.Addition(testCase.args), testCase.want)
			}
		})
	}
}

func TestValue_Square(t *testing.T) {
	var (
		testTable = []struct {
			name string
			args *Value
			want float64
		}{
			// Первый сценарий
			{
				name: "55.6687/5",
				args: &Value{55.6687, 5},
				want: 0,
			},
			// Второй сценарий
			{
				name: "-107853.22/220",
				args: &Value{37642, 220},
				want: 0,
			},
			// Третий сценарий
			{
				name: "22.22/2",
				args: &Value{22.22, 2},
				want: 0,
			},
			//Четвертый сценарий
			{
				name: "88476387348635/666326228364283",
				args: &Value{88476387348635, 66632},
				want: 0,
			},
		}
	)
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if AreaCalculated.Square(testCase.args) != math.Sqrt(testCase.args.A) {
				t.Errorf("Square() = %v, want %v", AreaCalculated.Addition(testCase.args), math.Sqrt(testCase.args.A))
			}
		})
	}
}