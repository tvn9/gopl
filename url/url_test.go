package url

// Tidying Up

// Writing tesable example
// Producing executable documentation
// Measuring test coverage and benchmarking
// Refacturing the URL parser
// Differences between external and internal tests

import (
	"fmt"
	"testing"
)

// Testing the string method
func TestURLString(t *testing.T) {
	u := &URL{
		Scheme: "http",
		Host:   "foo.com",
		Path:   "go",
	}

	got, want := u.String(), "https://foo.com/go"
	if got != want {
		t.Errorf("%#v.String()\ngot %q\nwant %q", u, got, want)
	}
}

// Share test TestURLPort test cases to test TestURLHostname

var hostTests = map[string]struct {
	in       string
	hostname string
	port     string
}{
	"with port":       {in: "foo.com:80", hostname: "foo.com", port: "80"},
	"with empty port": {in: "foo.com:", hostname: "foo.com", port: ""},
	"without port":    {in: "foo.com", hostname: "foo.com", port: ""},
	"ip with port":    {in: "127.0.0.1:90", hostname: "127.0.0.1", port: "90"},
	"ip without port": {in: "127.0.0.1", hostname: "127.0.0.1", port: ""},
}

// TestURLHostname
func TestURLHost(t *testing.T) {
	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.hostname), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.GetHost(), tt.hostname; got != want {
				t.Errorf("for host %q; got %q; want %q", tt.in, got, want)
			}
		})
	}
}

// TestURLPort
func TestURLPort(t *testing.T) {
	for name, tt := range hostTests {
		t.Run(fmt.Sprintf("%s/%s", name, tt.in), func(t *testing.T) {
			u := &URL{Host: tt.in}
			if got, want := u.GetPort(), tt.port; got != want {
				t.Errorf("for host %q; got %q; want %q", tt.in, got, want)
			}
		})
		// ofcause subtest can be addted to with new t.Run to test other cases
		// t.Run(""), func(t *testing.T) { ... }
	}
}
