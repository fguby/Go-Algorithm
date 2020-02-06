package utils

import "errors"

type Comparator func(e1, e2 interface{}) (int, error)

type Accumulator func(e1, e2 interface{}) (interface{}, error)

//int类型比较器
func IntComparator(e1, e2 interface{}) (int, error) {
	aAsserted, ok1 := e1.(int)
	if !ok1 {
		return 0, errors.New("can not compare")
	}
	bAsserted, ok2 := e2.(int)
	if !ok2 {
		return 0, errors.New("can not compare")
	}
	switch {
	case aAsserted > bAsserted:
		return 1, nil
	case aAsserted < bAsserted:
		return -1, nil
	default:
		return 0, nil
	}
}

//string类型比较器
func StringComparator(e1, e2 interface{}) (int, error) {
	aAsserted, ok1 := e1.(string)
	if !ok1 {
		return 0, errors.New("can not compare")
	}
	bAsserted, ok2 := e2.(string)
	if !ok2 {
		return 0, errors.New("can not compare")
	}

	//a1 := aAsserted[0]
	//b1 := bAsserted[0]

	//	fmt.Println(a1, b1)

	switch {
	case aAsserted[0] > bAsserted[0]:
		return 1, nil
	case aAsserted[0] < bAsserted[0]:
		return -1, nil
	default:
		return 0, nil
	}
}

//int类型累加器
func IntAccumulator(e1, e2 interface{}) (interface{}, error) {
	aAsserted, ok1 := e1.(int)
	if !ok1 {
		return nil, errors.New("can not compare")
	}
	bAsserted, ok2 := e2.(int)
	if !ok2 {
		return nil, errors.New("can not compare")
	}

	sum := aAsserted + bAsserted
	return sum, nil
}

//
//func Int8Comparator(e1, e2 interface{}) (int, error) {
//
//}
