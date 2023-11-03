package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"sort"
	"time"

	"github.com/20pa5a1210/cs-fundamentals/github"
	"github.com/20pa5a1210/cs-fundamentals/structs"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var arr [3]int
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	fmt.Println(len(arr))

	for i, v := range arr {
		fmt.Printf("%d %d\n", i, v)
	}
	println("--------------------------------")
	var newArr [3]int = [3]int{1, 2, 4}
	for i, v := range newArr {
		fmt.Printf("%d %d\n", i, v)
	}
	println("--------------------------------")

	symbols := [...]string{
		USD: "$",
		EUR: "€",
		GBP: "£",
		RMB: "¥",
	}
	fmt.Println(USD, symbols[USD])

	for i, v := range symbols {
		fmt.Printf("%d %s\n", i, v)
	}
	println("--------------------------------")
	r := [...]int{99: -1}
	fmt.Println(r[99])
	for i, v := range r {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Printf("%x", newArr)
	fmt.Println("------------------------------")
	exSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	exSlice1 := []int{1, 2, 3, 5, 5, 6, 7, 8, 9}
	fmt.Println(equal(exSlice, exSlice1))
	nilSlice := []int{}
	if nilSlice == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	// NOTE: slice is a reference type
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"]++
	ages["charlie"]++
	ages["charlie"]++
	ages["charlie"]++
	ages["charlie"]++
	ages["charlie"]++
	ages["charlie"]++
	fmt.Println(ages["charlie"])
	fmt.Println("------------------------------")
	newAges := map[string]int{
		"alice":   31,
		"charlie": 34,
		"bob":     20,
	}
	for name, age := range newAges {
		fmt.Printf("%s\t%d\n", name, age)
	}
	fmt.Println("------------------------------")
	var names []string
	for name := range newAges {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, newAges[name])
	}
	fmt.Println("------------------------------")

	var ages1 = make([]string, 0, len(newAges))
	fmt.Println(cap(ages1))
	fmt.Println(len(ages1))

	fmt.Println("----------- Structs -------------------")
	structs.Swapable()
	fmt.Println("------------------------------")
	fmt.Println("----------- Github Issue JSON------------------")
	input := []string{"repo:golang/go", "is:open", "json decoder"}
	result, err := github.SearchIssues(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:
{{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age:
{{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
func missingNumber(nums []int) int {

	/* sort.Ints(nums)
	   for i := 0; i < len(nums); i++ {
	       if nums[i] != i {
	           return i
	       }
	   }

	   return len(nums) */
	// sum of n consecutive numbers = n * (n + 1) / 2
	var sum int = len(nums) * (len(nums) + 1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}

func reverseArrayq(nums []int, start int, end int) {
	for start < end {
		var temp int = nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}

func singleNumber(nums []int) int {
	distMap := map[int]int{}

	for _, v := range nums {
		distMap[v]++
	}

	for k, v := range distMap {
		if v == 1 {
			return k
		}
	}
	return 0
}
func twoSum(nums []int, target int) []int {
	var SumMap = map[int]int{}

    for i:=0; i<len(nums);i++{
        var complement int = target - nums[i]
        if j, ok := SumMap[complement]; ok {
            return []int{j, i}
        }
        SumMap[nums[i]] = i
    }
    return []int{}
}

func sortcolors(nums []int){

    var red, white, blue int = 0, 0, len(nums)-1

    for white <= blue {
        if nums[white] == 0{
            swap(&nums[white], &nums[red])
            white++
            red++
        } else if nums[white] == 1{
            white++
        } else {
            swap(&nums[white], &nums[blue])
            blue--
        }
    }

}

func swap(a *int, b *int){
    *a, *b = *b, *a
}

