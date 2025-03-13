package gokit

import (
	"fmt"
	"testing"
	"time"
)

var a chan int = make(chan int)

func printSomething(p ...interface{}) {
	s := []string{}
	for _, v := range p {
		s = append(s, fmt.Sprintf("printing: %v", v))
	}
	fmt.Println(s)
	a <- 1
}

func TestDelayAction(t *testing.T) {
	DelayAction(time.Second*time.Duration(4), printSomething, "a", "b")
	<-a
}

func TestDelayActionAsync(t *testing.T) {
	DelayActionAsync(time.Second*time.Duration(4), printSomething, "a", "b")
	<-a
}
