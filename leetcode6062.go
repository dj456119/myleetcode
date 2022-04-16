package myleetcode

type ATM struct {
	Data []int
}

func Constructor6062() ATM {
	return ATM{Data: make([]int, 5)}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := range banknotesCount {
		this.Data[i] = banknotesCount[i] + this.Data[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	c := make([]int, 5)
	copy(c, this.Data)
	temp := make([]int, 5)
	for i := len(this.Data) - 1; i >= 0; i-- {
		if this.Data[i] == 0 {
			continue
		}
		z := amount / mm(i)
		if z == 0 {
			continue
		}
		if z <= this.Data[i] && amount%mm(i) == 0 {
			temp[i] = z
			this.Data[i] = this.Data[i] - z
			return temp
		}
		if z <= this.Data[i] && amount%mm(i) != 0 {
			temp[i] = z
			amount = amount - z*mm(i)
			this.Data[i] = this.Data[i] - z
			continue
		}
		temp[i] = this.Data[i]
		amount = amount - this.Data[i]*mm(i)
		this.Data[i] = 0
	}
	if amount == 0 {
		return temp
	}
	this.Data = c
	return []int{-1}
}

func mm(index int) int {
	switch index {
	case 0:
		return 20
	case 1:
		return 50
	case 2:
		return 100
	case 3:
		return 200
	case 4:
		return 500
	}
	return -1
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
