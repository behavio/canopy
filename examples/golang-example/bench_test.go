package golangexample_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"jsongoparser"
	"lispgoparser"
	"peggoparser"
)

var (
	jsonInput   string
	lispProgram = `(lambda (x y) (display "Hi.") (+ (* x y) 2))`
	pegGrammar  string
)

func init() {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to determine caller information")
	}

	baseDir := filepath.Dir(file)
	repoRoot := filepath.Dir(filepath.Dir(baseDir))

	jsonInput = mustRead(filepath.Join(repoRoot, "package.json"))
	pegGrammar = mustRead(filepath.Join(repoRoot, "examples", "canopy", "peg.peg"))
}

func mustRead(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("failed to read %s: %v", path, err))
	}
	return string(data)
}

func BenchmarkCanopyParseJSON(b *testing.B) {
	benchmarkParse(b, jsonInput, func(input string) error {
		_, err := jsongoparser.Parse(input, nil, nil)
		return err
	})
}

func BenchmarkCanopyParseLisp(b *testing.B) {
	benchmarkParse(b, lispProgram, func(input string) error {
		_, err := lispgoparser.Parse(input, nil, nil)
		return err
	})
}

func BenchmarkCanopyParsePEG(b *testing.B) {
	benchmarkParse(b, pegGrammar, func(input string) error {
		_, err := peggoparser.Parse(input, nil, nil)
		return err
	})
}

func benchmarkParse(b *testing.B, input string, parse func(string) error) {
	b.Helper()
	if err := parse(input); err != nil {
		b.Fatalf("initial parse failed: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := parse(input); err != nil {
			b.Fatalf("parse failed: %v", err)
		}
	}
}
