package main

import "testing"

//Table testing
var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-flaction", -1.0, -777.0, 0.0012870013, false},
}

func TestDevision(t *testing.T) {
	for _, tt := range tests {
	got, err := divide(tt.dividend, tt.divisor)
	if tt.isErr {
		if err == nil {
			t.Error("expected an error but did not get one")
		}
	} else {
		if err != nil {
			t.Error("did not expect an error but got one", err.Error())
		}
	}	
	
	if got != tt.expected {
		t.Errorf("expected %f but got %f", tt.expected, got)
	}
}

}
	

	
// Manual testing

func TestDivide(t *testing.T){
	_, err := divide(10.0, 1.0)
	if err != nil {
	t.Error("Got an error when we should not have")
	}
}
// in Terminal: go test
// or go test -v (for verbose)

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}


// in Terminal: go test -cover (shows the % coverage of the test)
// in Git Bash: go test -coverprofile=coverage.out && go tool cover -html=coverage.out