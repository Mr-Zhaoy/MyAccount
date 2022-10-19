package main

import (
	"fmt"
)

func main() {
	var key int8
	loop := true
	//定义账户
	balance := 10000.0
	money := 0.0
	note := ""
	deteils := "收支\t账户金额\t收支金额\t说 明"
	//定义一个标识符
	flag := false

	for {
		fmt.Println("-----------------家庭收支记账软件-----------------")
		fmt.Println("                 1.收支明细")
		fmt.Println("                 2.登记收入")
		fmt.Println("                 3.登记支出")
		fmt.Println("                 4.退出软件")
		fmt.Print("请选择（1-4）：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("\n\n-----------------当前收支明细记录-----------------")
			if flag == false {
				fmt.Println("你没有任何明细,来一笔吧")
			} else {
				fmt.Println(deteils)
			}

		case 2:
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			deteils += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case 3:
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			deteils += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case 4:
			fmt.Println("你确定要退出吗：y or n")
			Choice := ""
			for {
				fmt.Scanln(&Choice)
				if Choice == "y" || Choice == "n" {
					break
				}
				fmt.Println("你的输入有误")
			}
			if Choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确选项...")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出了家庭账单的使用...")
}
