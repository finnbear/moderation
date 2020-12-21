package moderation

import (
	"testing"
)

func IsProfane(phrase string) bool {
	return Analyze(phrase).IsInappropriate()
}

// from https://github.com/TwinProduction/go-away/blob/master/goaway_bench_test.go
func BenchmarkIsProfaneWhenShortStringHasNoProfanity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("aaaaaaaaaaaaaa")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenShortStringHasProfanityAtTheStart(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("fuckaaaaaaaaaa")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenShortStringHasProfanityInTheMiddle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("aaaaafuckaaaaa")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenShortStringHasProfanityAtTheEnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("aaaaaaaaaafuck")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenMediumStringHasNoProfanity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("How are you doing today?")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenMediumStringHasProfanityAtTheStart(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("Shit, you're cute today.")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenMediumStringHasProfanityInTheMiddle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("How are you fu ck doing?")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenMediumStringHasProfanityAtTheEnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("you're cute today. Fuck.")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenLongStringHasNoProfanity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("Hello John Doe, I hope you're feeling well, as I come today bearing terrible news regarding your favorite chocolate chip cookie brand")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenLongStringHasProfanityAtTheStart(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("Fuck John Doe, I hope you're feeling well, as I come today bearing terrible news regarding your favorite chocolate chip cookie brand")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenLongStringHasProfanityInTheMiddle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("Hello John Doe, I hope you're feeling well, as I come today bearing shitty news regarding your favorite chocolate chip cookie brand")
	}
	b.ReportAllocs()
}

func BenchmarkIsProfaneWhenLongStringHasProfanityAtTheEnd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsProfane("Hello John Doe, I hope you're feeling well, as I come today bearing terrible news regarding your favorite chocolate chip cookie bitch")
	}
	b.ReportAllocs()
}
