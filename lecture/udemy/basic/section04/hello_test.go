package hello

import "testing"

func TestHello(t *testing.T) {
	want := "안녕, 세상."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

// go test는 _test로 끝나는 .go 파일을 실행시킨다.