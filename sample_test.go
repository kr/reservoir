package reservoir

import (
	"math/rand"
	"reflect"
	"testing"
)

var s0 = "s0"

func TestNew(t *testing.T) {
	got := New[string](1)
	want := &Sample[string]{data: make([]string, 1)}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("New(1) = %v, want %v", got, want)
	}
}

func TestReset(t *testing.T) {
	rs := &Sample[string]{n: 1}
	rs.Reset()
	want := 0
	if rs.n != want {
		t.Errorf("after Reset, rs.n = %d, want %d", rs.n, want)
	}
}

func TestAdd0(t *testing.T) {
	rs := &Sample[string]{data: make([]string, 2)}
	rs.Add(s0)
	if rs.n != 1 {
		t.Errorf("after Add, n = %d, want 1", rs.n)
	}
	if rs.data[0] != s0 {
		t.Errorf("after Add, data[0] = %v, want s0", rs.data[0])
	}
	if rs.data[1] != "" {
		t.Errorf(`after Add, data[1] = %v, want ""`, rs.data[1])
	}
}

func TestAddN(t *testing.T) {
	rand.Seed(0)
	rs := &Sample[string]{data: make([]string, 2), n: 2}
	rs.Add(s0)
	if rs.n != 3 {
		t.Fatalf("after Add, rs.n = %d, want 3", rs.n)
	}
	if (rs.data[0] == "") == (rs.data[1] == "") {
		t.Errorf(`after Add, exactly one buffer cell must be ""`)
		t.Logf("%#v", rs)
	}
	if (rs.data[0] == s0) == (rs.data[1] == s0) {
		t.Errorf("after Add, exactly one buffer cell must be s0")
		t.Logf("%#v", rs)
	}
}

func TestRead(t *testing.T) {
	cases := []struct{ cap, param, added int }{
		{1, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}
	for _, test := range cases {
		rs := &Sample[string]{data: make([]string, test.cap), n: test.added}
		got := make([]string, test.param)
		ngot := rs.Read(got)
		if ngot != 1 {
			t.Errorf("case %+v ngot = %d, want 1", test, ngot)
		}
	}
}

func TestCap(t *testing.T) {
	rs := &Sample[string]{data: make([]string, 2)}
	got := rs.Cap()
	if got != 2 {
		t.Errorf("rs.Cap() = %d, want 2", got)
	}
}

func TestAdded(t *testing.T) {
	rs := &Sample[string]{n: 2}
	got := rs.Added()
	if got != 2 {
		t.Errorf("rs.Added() = %d, want 2", got)
	}
}

func TestZero(t *testing.T) {
	var rs Sample[string]
	gotCap := rs.Cap()
	if gotCap != 0 {
		t.Errorf("Sample[string]{}.Cap() = %d, want 0", gotCap)
	}
	gotAdded := rs.Added()
	if gotAdded != 0 {
		t.Errorf("Sample[string]{}.Added() = %d, want 0", gotAdded)
	}
	gotSample := make([]string, 5)
	gotSample = gotSample[:rs.Read(gotSample)]
	wantSample := []string{}
	if !reflect.DeepEqual(gotSample, wantSample) {
		t.Errorf("Sample[string]{}.Read() = %v, want %v", gotSample, wantSample)
	}

	rs.Add("a")
	rs.Add("b")
	rs.Add("c")

	gotAdded = rs.Added()
	if gotAdded != 3 {
		t.Errorf("Sample[string]{}.Added() = %d, want 3", gotAdded)
	}
	gotSample = make([]string, 5)
	gotSample = gotSample[:rs.Read(gotSample)]
	wantSample = []string{}
	if !reflect.DeepEqual(gotSample, wantSample) {
		t.Errorf("Sample[string]{}.Read() = %v, want %v", gotSample, wantSample)
	}
}
