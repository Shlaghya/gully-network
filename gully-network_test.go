package main

import (
	"testing"
)

func TestDecodeComment(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{
			input: "create_parking_lot 2",
			want:  "Created a parking lot with 2 slots",
		},
		{
			input: "park KA-00-HH-1234 White",
			want:  "Allocated slot number: 1",
		},
		{
			input: "park KA-00-HH-5678 Red",
			want:  "Allocated slot number: 2",
		},
		{
			input: "park KA-00-HH-5678 Red",
			want:  "Sorry, the parking lot is full",
		},
		{
			input: "leave 2",
			want:  "Slot number 2 is free",
		},
		{
			input: "leave 4",
			want:  "Please specify lot number within the lot size limit",
		},
		{
			input: "leave 2",
			want:  "There is no car parked in that slot",
		},
		{
			input: "registration_numbers_for_cars_with_colour White",
			want:  "KA-00-HH-1234",
		},
		{
			input: "slot_numbers_for_cars_with_colour white",
			want:  "1",
		},
		{
			input: "slot_number_for_registration_number KA-00-HH-1234",
			want:  "1",
		},
	}
	for _, test := range tests {
		got := DecodeComment(test.input)
		if got != test.want {
			t.Errorf("read() mismatch (want : %s \ngot : %s)\n", test.want, got)
		}
	}
}
