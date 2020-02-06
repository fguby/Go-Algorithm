package dynamic_programming

/**
	台阶问题:

	一共有N层台阶，一次最多只能走1或2台阶。
	一共有多少种走法？
**/

var (
	maps = make(map[int]int)
)

//动态规划
func GetClimbingWays3(steps int) int {

	temp1 := 1
	temp2 := 2
	ways := 0

	for i := 3; i <= steps; i ++ {
		ways = temp1 + temp2
		temp1 = temp2
		temp2 = ways
	}
	return ways
}

//寻常递归解法
func GetClimbingWays(steps int) int {

	if steps < 1 {
		return 0
	}
	
	if steps == 1 {
		return 1
	}

	if steps == 2 {
		return 2
	}

	ways := GetClimbingWays(steps-1) + GetClimbingWays(steps-2)

	return ways
}

//备忘录解法
func GetClimbingWays1(steps int) int {

	if steps == 1 || steps == 2 {
		return steps
	}

	if m, ok := maps[steps]; ok {
		return m
	}

	ways := GetClimbingWays1(steps-1) + GetClimbingWays1(steps-2)

	maps[steps] = ways
	return ways
}
