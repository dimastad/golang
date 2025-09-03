package main

import "fmt"

type Counter struct {
	value int
}
	
func (c *Counter) Inc(delta int) {
	d := 1
	
    if delta > 0 {
        d = delta
    }
    
    if d > 0 {
        c.value += int(d)
    }
	fmt.Println(c)
}

func (c *Counter) Dec(delta int) {
	d := 1
    if delta > 0 {
        d = delta
    }
    
    if d > 0 {
        if int(d) > c.value {
            c.value = 0
        } else {
            c.value -= int(d)
        }
    }
	fmt.Println(c)
}


func main() {
	c := &Counter{}

	c.Inc(0)
	c.Inc(10)
	c.Inc(1)
	c.Dec(2)
	c.Inc(5)
	c.Dec(10)
}