package main

import (
	"fmt"
	"os"
)

type production struct {
	rightPart string
	leftPart string
	number int
}

type params struct {
	inputString string
	L1 []string
	L2 []string
	N []string
	prodList map[string]production
}

func newProduction (leftPart string, rightPart string, number int) production{
	return production{
		leftPart: leftPart,
		rightPart: rightPart,
		number: number,
	}
}

func (p *params) NormalCondition(){
	for true {
		var L2TopStr = p.L2[len(p.L2)-1]
		if p.NTerminalCheck(L2TopStr) == true { // 1st step
			L2TopStr = p.FirstStep(L2TopStr)
		}

		if L2TopStr[0] == p.inputString[0] { //2nd step
			L2TopStr = p.SecondStep()
		}

		if (p.inputString == "") && (p.L2[len(p.L2) - 1] == "") { // 3rd step
			p.resCreate()
			return
		}

		if (p.L2[len(p.L2) - 1] == "" && len(p.inputString) != 0) ||	// 3rd' step
			(len(p.inputString) == 0 && p.L2[len(p.L2) - 1] != "") ||
			((p.NTerminalCheck(L2TopStr) == false) && (L2TopStr[0] != p.inputString[0])){	// 4th step
			fmt.Println("step: 3' or 4\n")
			p.ReturnStatus()
		}
	}
}

func (p *params) ReturnStatus(){
	fmt.Println("ReturnStatus")
	for true {
		if p.NTerminalCheck(p.L1[len(p.L1) - 1]) == true{ // 6th step
			nextAlterNumb := string(rune(int(p.L1[len(p.L1)-1][1]) + 1)) // 6a step
			newIndex := p.L1[len(p.L1)-1][:1] + nextAlterNumb
			if p.prodList[newIndex].number != 0 {
				p.SixthAStep(newIndex)
				return
			}
			if newIndex == "A2" { // 6b step
				fmt.Println("ERROR")
				os.Exit(0)
			} else {
				p.SixCStep()
			}
		}else{
			p.FifthStep()
		}
	}
}

func (p *params) FirstStep(topString string) string{
	fmt.Println("step: 1")
	key := topString[0:1] + "1"
	p.L1 = append(p.L1, key)
	topString = p.prodList[key].rightPart + topString[1:]
	p.L2 = append(p.L2, topString)
	fmt.Println(p.L1, "\t",p.L2, "\n")
	return topString
}

func (p *params) SecondStep() string{
	fmt.Println("step: 2")
	p.L1 = append(p.L1, p.inputString[:1])
	p.L2[len(p.L2) - 1] = p.L2[len(p.L2) - 1][1:]
	p.inputString = p.inputString[1:]
	fmt.Println(p.L1, "\t", p.L2, "\t", p.inputString, "\n")
	return p.L2[len(p.L2) - 1]
}

func (p *params) resCreate(){
	fmt.Println("step: 3")
	for _, v := range p.L1{
		if p.prodList[v].number != 0 {
			fmt.Print(p.prodList[v].number, " ")
		}
	}
	return
}

func (p *params) FifthStep(){
	fmt.Println("step: 5")
	p.L2 = append(p.L2, p.L1[len(p.L1) - 1] + p.L2[len(p.L2) - 1]) //reform
	p.inputString = p.L1[len(p.L1) - 1] + p.inputString
	p.L1 = p.L1[:len(p.L1) - 1]
	fmt.Println(p.L1, "\t", p.L2, "\t", p.inputString, "\n")
	return
}

func (p *params) SixthAStep(index string){
	fmt.Println("step: 6a")
	currentAlter := p.prodList[p.L1[len(p.L1) - 1]]
	newAlter := p.prodList[index]
	p.L2[len(p.L2) - 1] = newAlter.rightPart + p.L2[len(p.L2) - 1][len(currentAlter.leftPart):]
	p.L1[len(p.L1) - 1] = index
	fmt.Println(p.L1, "\t", p.L2, "\n")
	return
}

func (p *params) SixCStep(){
	fmt.Println("step: 6c")
	alternative := p.prodList[p.L1[len(p.L1) - 1]]
	p.L2[len(p.L2) - 1] = alternative.leftPart + p.L2[len(p.L2) - 1][len(alternative.rightPart):]
	p.L1 = p.L1[:len(p.L1) - 1]
	fmt.Println(p.L1, "\t", p.L2, "\n")
	return
}

func (p *params) NTerminalCheck(topString string) bool{
	for _, v := range p.N{
		if v == topString[:1]{
			return true
		}
	}
	return false
}


func main() {
	var p params
	p.L2 = append(p.L2, "A")

	p.prodList = make(map[string] production)
	p.prodList["A1"] = newProduction("A", "!B!", 1)
	p.prodList["B1"] = newProduction("B", "T", 2)
	p.prodList["B2"] = newProduction("B", "T+B", 3)
	p.prodList["T1"] = newProduction("T", "M", 4)
	p.prodList["T2"] = newProduction("T", "M*T", 5)
	p.prodList["M1"] = newProduction("M", "a", 6)
	p.prodList["M2"] = newProduction("M", "b", 7)
	p.prodList["M3"] = newProduction("M", "(B)", 8)

	p.N = append(p.N, "A", "B", "T", "M")

	fmt.Scanf("%s.n", &p.inputString)
	p.NormalCondition()
}