package utility

import "testing"

func TestGetMin(t *testing.T) {
	ans := GetMin(10,3)
	if ans != 3 {
		t.Errorf("GetMin(10,3) = %v, want 3", ans)
	}
	ans = GetMin(3,10)
	if ans != 3 {
		t.Errorf("GetMin(3,10 = %v, want 3", ans)
	}
}

func TestGetMinTable(t *testing.T) {
	var testCases = []struct{
		a, b, expected int
	} {
		{10,3,3},
		{100, 750, 100},
		{-5, -15, -15},
	}
	for _, test := range testCases {
		testName := fmt.Sprintf("GetMin(%v, %v)", test.a, test.b)
		testFunc := func(t *testing.T) {
			ans := GetMin(test.a, test.b)
			if ans != test.expected {
				t.Errorf("got %v, want %v", ans, test.expected)
			}
		}
		t,Run(testName, testFunc)
	}
}

func BenchmarkGetMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMin(10, 3)
	}
}