package calc

import (
	"testing"
)

func TestComputeValues_Success(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		n     int
		check func(t *testing.T, got []int, err error)
	}{
		"n=5 standard case": {
			n: 5,
			check: func(t *testing.T, got []int, err error) {
				t.Helper()
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				want := []int{100, 50, 33, 25, 20}
				if len(got) != len(want) {
					t.Fatalf("length mismatch: got %d, want %d", len(got), len(want))
				}
				for i := range got {
					if got[i] != want[i] {
						t.Errorf("index %d: got %d, want %d", i, got[i], want[i])
					}
				}
			},
		},
		"n=1": {
			n: 1,
			check: func(t *testing.T, got []int, err error) {
				t.Helper()
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if len(got) != 1 {
					t.Fatalf("expected 1 element, got %d", len(got))
				}
				if got[0] != 100 {
					t.Errorf("got %d, want 100", got[0])
				}
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := ComputeValues(tt.n)
			tt.check(t, got, err)
		})
	}
}

func TestSum(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input []int
		want  int
	}{
		"several elements": {input: []int{1, 2, 3, 4}, want: 10},
		"empty slice":      {input: []int{}, want: 0},
		"single element":   {input: []int{42}, want: 42},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := Sum(tt.input); got != tt.want {
				t.Errorf("Sum(%v) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestComputeValues_Error(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		n     int
		check func(t *testing.T, got []int, err error)
	}{
		"n=0 returns error": {
			n: 0,
			check: func(t *testing.T, got []int, err error) {
				t.Helper()
				if err != ErrNonPositive {
					t.Fatalf("expected ErrNonPositive, got %v", err)
				}
				if got != nil {
					t.Errorf("expected nil result, got %v", got)
				}
			},
		},
		"negative n returns error": {
			n: -3,
			check: func(t *testing.T, got []int, err error) {
				t.Helper()
				if err != ErrNonPositive {
					t.Fatalf("expected ErrNonPositive, got %v", err)
				}
				if got != nil {
					t.Errorf("expected nil result, got %v", got)
				}
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := ComputeValues(tt.n)
			tt.check(t, got, err)
		})
	}
}
