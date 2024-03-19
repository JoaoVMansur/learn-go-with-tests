package main

import "testing"

func TestRepeat(t *testing.T){
    repeated := Repeat("a", 10)
    expect  := "aaaaaaaaaa"

    if repeated != expect{
        t.Errorf("expected %q but got %q", expect, repeated)
    }
}

func BenchmarkRepeat(b *testing.B){
    for i := 0; i < b.N; i++{
        Repeat("a", 1)
    }
}
