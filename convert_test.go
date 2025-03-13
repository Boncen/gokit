package gokit

import (
	"fmt"
	"testing"
)

func TestCapitalUpper(t *testing.T) {
	s := "nice"
	if s1 := Capitalize(s); s1 != "Nice" {
		t.Errorf("nice Capitalize expect Nice,got:%v", s1)
	}

}

func TestToCamelCase(t *testing.T) {
	s := "user_name"
	if s1 := ToCamelCase(s, "_"); s1 != "userName" {
		t.Errorf("user_name TestToCamelCase expect userName,got:%v", s1)
	}
	s2 := "user-name"
	if s1 := ToCamelCase(s2, "-"); s1 != "userName" {
		t.Errorf("user-name TestToCamelCase expect userName,got:%v", s1)
	}
}

// func TestIntToInt32(t *testing.T) {
// 	fmt.Println(1 << 32)
// 	fmt.Println(-1 << 31)  // -2147483648
// 	fmt.Println(1<<31 - 1) // 2147483647

// 	fmt.Println(-1 << 63)  // -9223372036854775808
// 	fmt.Println(1<<63 - 1) // 9223372036854775807
// }

func TestIntToInt32(t *testing.T) {
	a := 100
	a32, ok := IntToInt32(a)
	if !ok {
		t.Errorf("expect a32 100,got:%v", a32)
	}

	a1 := 2147483648
	a32_2, ok := IntToInt32(a1)
	if ok {
		t.Errorf("expect a32_2 0,got:%v", a32_2)
	}

	a2 := -2147483649
	a32_3, ok := IntToInt32(a2)
	if ok {
		t.Errorf("expect a32_3 0,got:%v", a32_3)
	}
}

func TestInt32ToInt64(t *testing.T) {
	var a int32 = 100
	a64 := Int32ToInt64(a)
	fmt.Println(a64)
}

func TestInt32ToInt(t *testing.T) {
	var a int32 = 214748364
	aInt := Int32ToInt(a)
	fmt.Println(aInt)
}

func TestToString(t *testing.T) {
	var i int = 100
	var i32 int32 = 200
	var i64 int64 = 300
	var f32 float32 = 200.11
	var f64 float64 = 20022.11

	fmt.Println(ToString(i))
	fmt.Println(ToString(i32))
	fmt.Println(ToString(i64))
	fmt.Println(ToString(f32))
	fmt.Println(ToString(f64))
}
