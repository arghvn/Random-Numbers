package main

//Go’s math/rand package provides pseudorandom number generation.

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//For example, rand.Intn returns a random int n, 0 <= n < 100.

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64())

	//rand.Float64 returns a float64 f, 0.0 <= f < 1.0.

	//This can be used to generate random floats in other ranges, for example 5.0 <= f' < 10.0.

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	//The default number generator is deterministic, so it’ll produce the same
	// sequence of numbers each time by default. To produce varying sequences,
	// give it a seed that changes. Note that this is not safe to use for random numbers
	// we intend to be secret, use crypto/rand for those.

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	//Call the resulting rand.Rand just like the functions on the rand package.

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}

//If we seed a source with the same number, it produces the same sequence of random numbers.

//output :

//81,87
//0.6645600532184904
//7.123187485356329,8.434115364335547
//0,28
//5,87
//5,87
