package main

import (
	"math"
	"testing"
)

func TestCalcMean(t *testing.T) {
	tests := []struct {
		name   string
		values []float64
		want   float64
	}{
		{"empty slice", []float64{}, 0},
		{"single element", []float64{5}, 5},
		{"multiple elements", []float64{1, 2, 3, 4, 5}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcMean(tt.values); got != tt.want {
				t.Errorf("calcMean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcMedian(t *testing.T) {
	tests := []struct {
		name   string
		values []float64
		want   float64
	}{
		{"empty slice", []float64{}, 0},
		{"single element", []float64{5}, 5},
		{"even count", []float64{1, 2, 3, 4, 5, 6}, 3.5},
		{"odd count", []float64{1, 2, 3, 4, 5}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcMedian(tt.values); got != tt.want {
				t.Errorf("calcMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcMode(t *testing.T) {
	tests := []struct {
		name   string
		values []float64
		want   float64
	}{
		{"empty slice", []float64{}, 0},
		{"single element", []float64{5}, 5},
		{"multiple elements", []float64{1, 2, 2, 3}, 2},
		{"no mode", []float64{1, 2, 3}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcMode(tt.values); got != tt.want {
				t.Errorf("calcMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcStandardDeviation(t *testing.T) {
	tests := []struct {
		name   string
		values []float64
		mean   float64
		want   float64
	}{
		{"empty slice", []float64{}, 0, 0},
		{"single element", []float64{5}, 5, 0},
		{"multiple elements", []float64{1, 2, 3, 4, 5}, 3, math.Sqrt(2)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcStandartDeviation(tt.values, tt.mean); math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("calcStandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
