package dynamic_programming

import "math"

type obj struct {
	golds  int
	people int
}

/**
	有一个国家发现了5座金矿，每座金矿的黄金储量不同.
	需要参与挖掘的工人数也不同。参与挖矿工人的总数是10人。
	每座金矿要么全挖，要么不挖，不能派出一半人挖取一半金矿。
	要想得到尽可能多的黄金，应该选择挖取哪几座金矿？
**/
var (
	m = make(map[obj]int)
	//代表金矿
	golds = []int{500, 400, 350, 300, 200}
	//代表每座金矿需要的人数
	peoples = []int{5, 5, 3, 4, 3}
)

//普通递归方式
func GetMaxGold(n, people int) int {

	//人数少
	if people < peoples[n] {
		return 0
	}

	if n == 0 {
		return golds[n]
	}

	if max, ok := m[obj{golds: n, people: people}]; ok {
		return max
	}

	maxGolds := math.Max(float64(GetMaxGold(n-1, people)),
		float64(GetMaxGold(n-1, people-peoples[n])+golds[n]))

	m[obj{golds: n, people: people}] = int(maxGolds)

	return int(maxGolds)
}

//动态规划
func GetMaxGold1(n, people int) int {

	/**
		F(N, W) = Max(
			F(N-1, W),
			F(N-1, 10-w)+F(N, P(N))
		)
	**/

	maxGolds := make([]int, people)

	for p := 0; p < people; p++ {
		if (p + 1) < peoples[0] {
			maxGolds[p] = 0
		} else {
			maxGolds[p] = golds[0]
		}
	}

	i := 1

	for i < (n + 1) {

		tempGolds := make([]int, people)

		copy(tempGolds, maxGolds)

		p := 0

		for p < people {

			m1 := getGolds(i, p)

			m2 := maxGolds[p]

			if (p + 1) > peoples[i] {
				m1 += maxGolds[p-peoples[i]]
			} else if (p + 1) < peoples[i] {
				m1 = 0
			}

			if m1 > m2 {
				tempGolds[p] = m1
			} else {
				tempGolds[p] = m2
			}

			p++
		}

		maxGolds = tempGolds

		i++
	}

	return maxGolds[people-1]
}

func getGolds(n, people int) int {
	if people < (peoples[n] - 1) {
		return 0
	}
	return golds[n]
}
