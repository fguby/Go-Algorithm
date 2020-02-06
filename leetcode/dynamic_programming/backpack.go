package dynamic_programming

/**
有一个背包，背包容量是M=150kg。有7个物品，物品不可以分割成任意大小。要求尽可能让装入背包中的物品总价值最大，但不能超过总容量。
物品 A B C D E F G
重量 35kg 30kg 6kg 50kg 40kg 10kg 25kg
价值 10$ 40$ 30$ 50$ 35$ 40$ 30$
**/

var (
	//背包问题
	weights = []int{35, 30, 6, 50, 40, 10, 25}
	prices  = []int{10, 40, 30, 50, 35, 40, 30}
)

/**
	150Kg在7个物品里挑选最大价值
		Max： 150里6个物品挑选，（150-p）6个物品里 + 第7个物品价值

**/

//动态规划解法
func GetMaxPrices(weight, count int) int {

	//maxPrices := getPrices(weight, 1)

	maxPrices := make([]int, weight+1)

	for k := range maxPrices {
		if k >= weights[0] {
			maxPrices[k] = prices[0]
		}
	}

	for i := 1; i < len(weights); i++ {
		temp := make([]int, len(maxPrices))
		for k := 1; k < len(maxPrices); k++ {
			p1 := getPrices(k, i+1)
			if k > weights[i] {
				p1 += maxPrices[k-weights[i]]
			}
			p2 := maxPrices[k]
			temp[k] = comparePrices(p1, p2)
		}
		copy(maxPrices, temp)
	}

	return maxPrices[weight]
}

//递归解法
func GetMaxPrices1(weight, count int) int {
	if count == 1 {
		if weight >= weights[0] {
			return prices[0]
		} else {
			return 0
		}
	}

	maxPrices := comparePrices(
		GetMaxPrices1(weight, count-1),
		GetMaxPrices1(weight-weights[count-1], count-1)+getPrices(weight, count),
	)

	return maxPrices
}

func getPrices(weight, n int) int {
	if weight >= weights[n-1] {
		return prices[n-1]
	}
	return 0
}

func comparePrices(p1, p2 int) int {
	if p1 > p2 {
		return p1
	}
	return p2
}
