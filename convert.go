package gokit

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

// Strs2Map
func Strs2Map(s string) map[string]struct{} {
	result := make(map[string]struct{}, 0)
	if s == "" {
		return result
	}
	parts := strings.Split(s, ",")
	for _, v := range parts {
		result[v] = struct{}{}
	}
	return result
}

// e.g. user_name -> userName
func ToCamelCase(s string, seperator string) string {
	if s == "" {
		return ""
	}
	parts := strings.Split(s, seperator)
	result := []string{}
	for k, v := range parts {
		if k == 0 {
			result = append(result, v)
			continue
		}
		result = append(result, Capitalize(v))
	}
	return strings.Join(result, "")
}

// 首字母大写
func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

// int to int32
func IntToInt32(i int) (int32, bool) {
	if i < -1<<31 || i > 1<<31-1 {
		return 0, false
	}
	return int32(i), true
}

// int to int64
func IntToInt64(i int) (int64, bool) {
	if i < -1<<63 || i > 1<<63-1 {
		return 0, false
	}
	return int64(i), true
}

// int32 to int64
func Int32ToInt64(i int32) int64 {
	return int64(i)
}

// int32 to int
func Int32ToInt(i int32) int {
	return int(i)
}

// int64 to int32
func Int64ToInt32(i int64) (int32, bool) {
	if i < -1<<31 || i > 1<<31-1 {
		return 0, false
	}
	return int32(i), true
}

// int64 to int
func Int64ToInt(i int64) (int, bool) {
	if i < math.MinInt32 {
		return 0, false
	}
	return int(i), true
}

// folat32 to string
// float64 to string

// struct to string

// string to int
// string to int32
// string to int64

// any to json
// json to []

// string to time

// time to UTC-String

func ToString(i interface{}) string {
	return fmt.Sprint(i)
}
