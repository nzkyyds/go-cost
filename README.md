# go-cost

## install

go get github.com/nzkyyds/go-cost/cost

## usage

```go
package main

func main(){
    cost.Begin("camera-001")
    cost.Printf("hello world")

    time.Sleep(time.Millisecond * 56)
    cost.Printf("welcome to Beijing")
    cost.End()
}

```
## print

2022/08/04 09:56:09 camera-001 duration=0s, cost=0s, cost Begin

2022/08/04 09:56:09 camera-001 duration=31.2355ms, cost=0s, hello world

2022/08/04 09:56:09 camera-001 duration=89.1032ms, cost=57.8677ms, welcome to Beijing

2022/08/04 09:56:09 camera-001 duration=89.1032ms, cost=0s, cost End