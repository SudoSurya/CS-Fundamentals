package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Workflow struct {
	Conditions []Condition
	Redirect   string
}
type Condition struct {
	variable string
	operator string
	value    int
	redirect string
}

var pattern = `\{([^{}]*)\}`
var regexBraces = regexp.MustCompile(pattern)
var regexNumber = regexp.MustCompile(`\d+`)

func main() {
	input := readInput()
	workflowMap := map[string]Workflow{}
	worflows := input[0]
	for _, workflow := range strings.Split(worflows, "\n") {
		workflowName := strings.Split(workflow, "{")[0]
		workflowMap[workflowName] = Workflow{}
		conditions := regexBraces.FindStringSubmatch(workflow)
		if len(conditions) > 1 {
			redirect := strings.Split(conditions[1], ",")[len(strings.Split(conditions[1], ","))-1]
			indvidualConditions := strings.Split(conditions[1], ",")
			for i := 0; i < len(indvidualConditions)-1; i++ {
				inferRedirect := strings.Split(indvidualConditions[i], ":")[1]
				value := regexNumber.FindAllString(indvidualConditions[i], -1)[0]
				variable := strings.Split(indvidualConditions[i], "")[0]
				operator := strings.Split(indvidualConditions[i], "")[1]

				cond := Condition{
					variable: variable,
					operator: operator,
					value:    StringToInt(value),
					redirect: inferRedirect,
				}

				workflowMap[workflowName] = Workflow{
					Conditions: append(workflowMap[workflowName].Conditions, cond),
					Redirect:   redirect,
				}

			}
		}
	}
	ratings := input[1]
	sumOfRatings := 0
	for _, rating := range strings.Split(ratings, "\n") {
		ratingsNumber := regexNumber.FindAllString(rating, -1)
		x := StringToInt(ratingsNumber[0])
		m := StringToInt(ratingsNumber[1])
		a := StringToInt(ratingsNumber[2])
		s := StringToInt(ratingsNumber[3])

		Result := GetResult(workflowMap, x, m, a, s, "in")
		fmt.Println(Result)
		if Result == "A" {
			sumOfRatings += (x + m + a + s)
		}
	}
	fmt.Println("start")
	fmt.Println(sumOfRatings)
    part2()
}

func StringToInt(str string) int {
	nums, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return nums
}

func GetResult(workflowMap map[string]Workflow, x, m, a, s int, workFlowIdx string) string {

	Result := ""

	workflow := workflowMap[workFlowIdx]
	isRedirect := false
	for _, condition := range workflow.Conditions {
		if condition.variable == "x" {
			if condition.operator == "<" {
				if x < condition.value {
					workFlowIdx = condition.redirect
					Result = condition.redirect
					isRedirect = true
					break
				}
			} else if condition.operator == ">" {
				if x > condition.value {
					workFlowIdx = condition.redirect
					Result = condition.redirect
					isRedirect = true
					break
				}
			}
		} else if condition.variable == "m" {
			if condition.operator == "<" {
				if m < condition.value {
					workFlowIdx = condition.redirect
					Result = condition.redirect
					isRedirect = true
					break
				}
			} else if condition.operator == ">" {
				if m > condition.value {
					workFlowIdx = condition.redirect
					Result = condition.redirect
					isRedirect = true
					break
				}
			}
		} else if condition.variable == "a" {
			if condition.operator == "<" {
				if a < condition.value {
					workFlowIdx = condition.redirect
					Result = condition.redirect
					isRedirect = true
					break
				}
			} else if condition.operator == ">" {
				if a > condition.value {
					workFlowIdx = condition.redirect
					Result = condition.redirect
					isRedirect = true
					break
				}
			}
		} else if condition.variable == "s" {
			if condition.operator == "<" {
				if s < condition.value {
					workFlowIdx = condition.redirect
					Result = condition.redirect
					isRedirect = true
					break
				}
			} else if condition.operator == ">" {
				if s > condition.value {
					workFlowIdx = condition.redirect
					Result = condition.redirect
					isRedirect = true
					break
				}
			}
		}
	}
	if !isRedirect {
		workFlowIdx = workflow.Redirect
		Result = workflow.Redirect
	}
	if Result == "A" || Result == "R" {
		return Result
	} else {
		return GetResult(workflowMap, x, m, a, s, workFlowIdx)
	}
}

func readInput() []string {
	inputByte, err := os.ReadFile("../input/sample")
	if err != nil {
		panic(err)
	}
	inputByte = inputByte[:len(inputByte)-1]
	inputStr := string(inputByte)
	input := strings.Split(inputStr, "\n\n")
	return input
}
