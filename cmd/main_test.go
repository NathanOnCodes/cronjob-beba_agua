package main

import (
	"bytes"
	"os/exec"
	"regexp"
	"syscall"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	cmd := exec.Command("somecommand", "parameter") 
	cmd.Stdout = &buf
	go cmd.Run()

	time.Sleep(2 * time.Hour)

	want := regexp.MustCompile(`Beba agua, ja são: \d{2}:\d{2}\nse voce bebeu um copo de agua que geralmente tem 200ml, no total voce bebeu durante a execução: \d+ml\n`)
	got := buf.String()
	if !want.MatchString(got) {
		t.Errorf("main() printed %q; want match for %q", got, want)
	}

	wantCode := 0
	if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
		gotCode := status.ExitStatus()
		if gotCode != wantCode {
			t.Errorf("main() exited with code %d; want %d", gotCode, wantCode)
		}
	}
}


func TestCountFinal(t *testing.T) {
	tests := []struct {
		value int
		want  string
	}{
		{0, "se voce bebeu um copo de agua que geralmente tem 200ml, no total voce bebeu durante a execução: 0ml"},
		{1, "se voce bebeu um copo de agua que geralmente tem 200ml, no total voce bebeu durante a execução: 200ml"},
		{2, "se voce bebeu um copo de agua que geralmente tem 200ml, no total voce bebeu durante a execução: 400ml"},
	}

	for _, test := range tests {
		got := countFinal(test.value)
		if got != test.want {
			t.Errorf("countFinal(%d) = %q; want %q", test.value, got, test.want)
		}
	}

}
