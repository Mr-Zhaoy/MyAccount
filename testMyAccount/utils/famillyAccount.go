package utils

import "fmt"

type FamilyAccount struct {
	key     string
	loop    bool
	balance float64
	money   float64
	note    string
	details string
	flag    bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说 明",
	}
}

// 显示明细
func (this *FamilyAccount) showDetails() {
	fmt.Println("\n\n-----------------当前收支明细记录-----------------")
	if this.flag == false {
		fmt.Println("你没有任何明细,来一笔吧")
	} else {
		fmt.Println(this.details)
	}

}

// 显示收入
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 显示支出
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true

}

// 退出菜单
func (this *FamilyAccount) exit() {
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
		this.loop = false
	}
}

// 显示主菜单
func (this *FamilyAccount) Mainmenu() {
	for {
		fmt.Println("-----------------家庭收支记账软件-----------------")
		fmt.Println("                 1.收支明细")
		fmt.Println("                 2.登记收入")
		fmt.Println("                 3.登记支出")
		fmt.Println("                 4.退出软件")
		fmt.Print("请选择（1-4）：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
			fmt.Println("你退出了家庭账单的使用...")
		default:
			fmt.Println("请输入正确选项...")
		}
		if !this.loop {
			break
		}
	}
}
