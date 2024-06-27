/* package main

import (
	"fmt"
	"io"
	"math"
	"strings"

	"golang.org/x/tour/tree"
)

func Pic(dx, dy int) [][]uint8 {

	img := make([][]uint8, dy)

	for y := range img {

		img[y] = make([]uint8, dx)

		for x := range img[y] {
			img[y][x] = uint8((x+y)/2*(y-x) - 220)
		}
	}

	return img
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(s) {
		m[word] += 1
	}
	return m
}

func fibonacci() func() int {
	f2, f1 := 0, 1
	return func() int {
		f := f2
		f2, f1 = f1, f+f1

		return f
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, &ErrNegativeSqrt{x}
	}
	z := x / 2
	for i := 0; i < 10; {
		if diff := (z*z - x) / (2 * z); math.Abs(diff) < 0.0005 {
			i++
			fmt.Println("diff: ", diff)
			fmt.Println("icount: ", i)
		} else {
			z -= (z*z - x) / (2 * z)
			fmt.Println("diff: ", diff)
			fmt.Println("Z: ", z)
		}
	}
	/* return z */
	return 0, nil
}

type ErrNegativeSqrt struct {
	value float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number %v",
		e.value)
}

type MyReader struct{}

func (read MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, e := rot.r.Read(b)

	for i := 0; i < n; i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return n, e
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)
	for n := range ch1 {
		if n != <-ch2 {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	/* 	fmt.Println(Sqrt(9))
		pic.Show(Pic)

	   	wc.Test(WordCount)

	   	f := fibonacci()
	   	for i := 0; i < 10; i++ {
	   		fmt.Println(f())
	   	}

		hosts := map[string]IPAddr{
			"loopback":  {127, 0, 0, 1},
			"googleDNS": {8, 8, 8, 8},
		}
		for name, ip := range hosts {
			fmt.Printf("%v: %v\n", name, ip)
		}

		fmt.Println(Sqrt(2))
		fmt.Println(Sqrt(-2))

		reader.Validate(MyReader{})

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)*/
}
 */