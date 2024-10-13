package main

func main() {
	c, out := make(chan int), make(chan int)

	m := map[int]int{1: 2, 3: 4}
	for i, v := range m {
		go func() {
			<-c
			out <- i + v
		}()
	}

	close(c)

	println(<-out + <-out)
}
