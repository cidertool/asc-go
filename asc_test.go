package asc

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if got, want := c.common.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.common.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}

	c2 := NewClient(nil)
	if c.common.Client == c2.common.Client {
		t.Error("NewClient returned same http.Clients, but they should differ")
	}
}

func TestBool(t *testing.T) {
	got := true
	want := Bool(got)
	if &got == want {
		t.Error("Bool returned same *bool, but they should differ")
	}
}

func TestInt(t *testing.T) {
	got := 100
	want := Int(got)
	if &got == want {
		t.Error("Int returned same *int, but they should differ")
	}
}

func TestFloat(t *testing.T) {
	got := 100.5
	want := Float(got)
	if &got == want {
		t.Error("Float returned same *float64, but they should differ")
	}
}

func TestString(t *testing.T) {
	got := "App Store Connect"
	want := String(got)
	if &got == want {
		t.Error("String returned same *string, but they should differ")
	}
}
