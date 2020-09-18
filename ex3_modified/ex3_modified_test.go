package ex3_modified

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

func BenchmarkSql(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		VanillaExample(testdb)
	}
}

func BenchmarkSqlx(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		SqlxExample(testdb)
	}
}

func BenchmarkSq(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		SqExample(testdb)
	}
}
