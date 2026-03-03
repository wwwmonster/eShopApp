package testing

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wwwmonster/eShopApp/go/v2/configs"
	"github.com/wwwmonster/eShopApp/go/v2/internal/helper"
)

func Testing() {
}

func testYy(i int) func(yield func(i int) bool) {
	return func(yield func(i int) bool) {
		for x := range i {
			if !yield(x) {
				return
			}
		}
	}
}

func testMap() {
	results := helper.Stream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Filter(func(x int) bool { return x%2 == 0 }).
		Map(func(x int) int { return x * 10 }).
		Collect()

	fmt.Println(results)
}

func Countdown(v int) func(func(int) bool) {
	// next, we return a callback func which is typically
	// called yield, but names like next could also be
	// applicable
	return func(y func(int) bool) {
		// we then start a for loop that iterates
		for i := v; i >= 0; i-- {
			// once we've finished looping
			if !y(i) {
				// we then return and finish our iterations
				return
			}
		}
	}
}

func mainTest() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s1 := s[:0]
	printSlice(s1)

	// Extend its length.
	s2 := s[:4]
	printSlice(s2)

	// Drop its first two values.
	s3 := s[2:]
	printSlice(s3)
}

func mainTest1() {
	fmt.Println("teste")
	fmt.Println("2222")

	configs.LoadAppSettings()
	app := fiber.New()

	myArr := []string{"tewt1", "test2"}
	fmt.Println(myArr)

	myArr = append(myArr, "teste3")

	fmt.Println(myArr)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("for testing only")
	})
	//	_ = app.Listen(":9000")

	myCourse := [][]string{
		{"mat", "eng"},
		{"teste", "jnfek"},
	}

	fmt.Println(myCourse)
	newCourse := make([]string, 2, 3)

	newCourse = append(newCourse, ("1-1"), ("1-2"), ("fds"))
	fmt.Println("newCourse:  ", newCourse)
	//:wwwmonster	newCourse = append(newCourse, "1")
	mainTest()
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Println("tes")
}

func testVim(i int) {
	fmt.Println("tet:", i)
	fmt.Println("tet:", i)
	fmt.Println("tet:", i)
	fmt.Println("tet:", i)
	fmt.Println("tet:", i)
	fmt.Println("tet:", i)
}
