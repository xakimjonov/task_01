
func main() {
	var a, b int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Println(a, "is odd: ", isOdd(a))
    fmt.Println(b, "is even: ", isEven(b))
}


func isEven(num int) bool {
  return(num %2 == 0)
	
}
func isOdd (num int) bool {
	return (num %2 != 0)
}

}

func area(r float32) (area float32) {
  const PI = 3.14159265

	var l = r * 2
	var area_square = l * l
	var c = PI * r * r
	area = area_square - c

	return area

}
