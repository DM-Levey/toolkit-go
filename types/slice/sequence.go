package slice

import "fmt"

// Sequence 顺序表实现
type Sequence[T any] struct{}

// Add 添加元素
func Add[T any](source []T, elem T, index int) ([]T, error) {

	//判断下标越界
	if index < 0 || index > len(source) {
		return nil, fmt.Errorf("下标越界")
	}

	// 不用显示判断是否扩容，交给append

	temp := source[len(source)-1]
	for i := len(source) - 1; i > index; i-- {
		source[i] = source[i-1]
	}
	source = append(source, temp)
	source[index] = elem

	return source, nil
}

func Delete[T any](source []T, index int) ([]T, error) {

	//判断下标越界
	if index < 0 || index > len(source)-1 {
		return nil, fmt.Errorf("下标越界")
	}
	newSlice := source
	// 如果当前长度大于64，且小于容量的1/4，缩容一半
	if len(source) > 64 && len(source) < cap(source)>>2 {
		newSlice = make([]T, len(source)-1, cap(source)>>1)
	}
	j := 0
	for i := range source {
		if i == index {
			continue
		}
		newSlice[j] = source[i]
		j++
	}

	return newSlice[:len(source)-1], nil
}
