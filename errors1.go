package main

import "fmt"

type lame interface {
	answer()
}

type kvadrat struct {
	UDside int
	LRside int
}

func (k kvadrat) square() (int, error) {
	if k.UDside == 0 || k.LRside == 0 {
		return 0, fmt.Errorf("что то здесь равно нулю LR: %d, UD: %d ", k.UDside, k.LRside)
	}
	return k.LRside * k.UDside, nil
}

func (k kvadrat) answer() {
	fmt.Println("ответы")
	s, err := k.square()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("square:", s)
	}
}

func main() {
	kv1 := kvadrat{
		UDside: 12,
		LRside: 32,
	}
	kv1.answer()

}
