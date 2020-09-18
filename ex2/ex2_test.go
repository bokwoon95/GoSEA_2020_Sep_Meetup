package ex2

import (
	"database/sql"
	"testing"
)

var testdb *sql.DB

func init() {
	var err error
	testdb, err = sql.Open("postgres", "postgres://pg:pg@localhost:5431/gosea_2020_sep_meetup?sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func BenchmarkSql1(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		VanillaExample1(testdb)
	}
}

func BenchmarkSql2(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		VanillaExample2(testdb)
	}
}

func BenchmarkSql3(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		VanillaExample3(testdb)
	}
}

func BenchmarkSqlx1(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		SqlxExample1(testdb)
	}
}

func BenchmarkSqlx2(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		SqlxExample2(testdb)
	}
}

func BenchmarkSqlx3(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		SqlxExample3(testdb)
	}
}

func BenchmarkSq1(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		SqExample1(testdb)
	}
}

func BenchmarkSq2(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		SqExample2(testdb)
	}
}

func BenchmarkSq3(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		SqExample3(testdb)
	}
}
