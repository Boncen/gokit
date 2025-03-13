package gokit

import "time"

// DeduplicateSlice 接受一个字符串切片，并返回一个新的没有重复元素的切片
func DeduplicateSlice(input []string) []string {
	uniqueMap := make(map[string]struct{})
	var result []string

	for _, item := range input {
		if item == "" {
			continue
		}
		if _, exists := uniqueMap[item]; !exists {
			uniqueMap[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// 延迟执行
func DelayAction(dur time.Duration, f func(p ...interface{}), params ...interface{}) {
	time.Sleep(dur)
	f(params)
}

// 异步延迟执行
func DelayActionAsync(dur time.Duration, f func(p ...interface{}), params ...interface{}) {
	go func() {
		DelayAction(dur, f, params...)
	}()
}
