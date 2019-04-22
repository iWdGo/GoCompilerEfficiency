package urlbuilder

import (
	"log"
	"net/url"
	"testing"
)

const (
	want   = "https://testinghello-in-the-cloud.appspot.com/"
	mypath = "mypath"
)

func TestGetAppString(t *testing.T) {
	if u := getAppString("/"); u != want {
		t.Fatalf("got %s, want %s\n", u, want)
	}
}

func TestGetAppUrl(t *testing.T) {
	if u := getAppUrl("/").String(); u != want {
		t.Fatalf("got %s, want %s\n", u, want)
	}
}

// w/o the for loop, no benchmarking occurs because of optimization
func BenchmarkGetAppString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := getAppString(mypath)
		if s != want+mypath {
			b.Logf("got %s, want %s", s, want+mypath)
		}
		u, err := url.Parse(s)
		if err != nil {
			log.Fatal(err)
		}
		if r := u.Path; r != "/"+mypath {
			log.Fatalf("raw path: got %s, want %s", r, mypath)
		}
	}
}

func BenchmarkGetAppUrl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := getAppUrl(mypath).String()
		if s != want+mypath {
			b.Logf("got %s, want %s", s, want+mypath)
		}
		u, err := url.Parse(s)
		if err != nil {
			log.Fatal(err)
		}
		if r := u.Path; r != "/"+mypath {
			log.Fatalf("raw path: got %s, want %s", r, mypath)
		}
	}
}
