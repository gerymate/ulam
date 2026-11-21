package spiral

import "testing"

func TestSpiral(t *testing.T) {
	cases := []struct {
		in int
		want string
	} {
		{1, "1"},
	}
	for _, c := range cases {
		got := Spiral(c.in)
		if got != c.want {
			t.Errorf("Spiral(%d)==%q, want: %q", c.in, got, c.want)
		}
	}
}