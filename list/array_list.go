package list

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	Data []int
	Size int
}

func (arrayList *ArrayList) GetLength() int {
	return arrayList.Size
}

func (arrayList *ArrayList) GetCapacity() int {
	return cap(arrayList.Data)
}

func (arrayList *ArrayList) IsEmpty() bool {
	return arrayList.Size == 0
}

func (arrayList *ArrayList) Add(index int, value int) error {
	if arrayList == nil || arrayList.Size == 0 {
		if index != 0 {
			return errors.New("add faild. arrayList is empty. index must be 0")
		} else {
			if arrayList.Data == nil {
				arrayList.Data = []int{value}
				arrayList.Size++
			} else {
				arrayList.Data = append(arrayList.Data, value)
				arrayList.Size++

			}
		}
		return nil
	}
	if index < 0 {
		return errors.New("add faild. index out of range")
	}

	//判断数组是否有剩余空间，先把值放到最后
	if arrayList.Size < len(arrayList.Data) {
		arrayList.Data[arrayList.Size] = value
	} else {
		arrayList.Data = append(arrayList.Data, value)
	}
	arrayList.Size++

	//放到最后的值挪动的相对应的index的位置
	var tempValue int = 0
	for i := arrayList.Size - 1; i > index; i-- {
		tempValue = arrayList.Data[i]
		arrayList.Data[i] = arrayList.Data[i-1]
		arrayList.Data[i-1] = tempValue
	}

	return nil
}

func (arrayList *ArrayList) AddFirst(value int) error {
	return arrayList.Add(0, value)
}

func (arrayList *ArrayList) AddLast(value int) error {
	return arrayList.Add(arrayList.Size, value)
}

func (arrayList *ArrayList) Find(value int) int {
	if arrayList == nil || arrayList.Data == nil {
		return -1
	}
	for i, v := range arrayList.Data {
		if v == value {
			return i
		}
	}
	return -1
}

func (arrayList *ArrayList) Contains(value int) bool {
	if arrayList == nil || arrayList.Data == nil {
		return false
	}
	for i := range arrayList.Data {
		if value == arrayList.Data[i] {
			return true
		}
	}
	return false
}

func (arrayList *ArrayList) Remove(index int) (bool, error) {
	if index < 0 || index > arrayList.Size-1 {
		return false, errors.New("remove faild. index out of range")
	}

	//因为是删除，所以直接把对应index后的值都向前挪一个位置
	for i := index; i < arrayList.Size-1; i++ {
		arrayList.Data[i] = arrayList.Data[i+1]
	}
	arrayList.Data[arrayList.Size-1] = 0
	arrayList.Size--

	//为了防止占用太多的空间，需及时回收剩余空间
	if arrayList.Size < len(arrayList.Data)/2 {
		arrayList.ReSetSize()
	}

	return true, nil
}

func (arrayList *ArrayList) ReSetSize() {
	//根据需求这里可以分配更多，或不分配剩余空间
	values := make([]int, arrayList.Size*3/2)
	for i := range arrayList.Data {
		if arrayList.Data[i] != 0 {
			values[i] = arrayList.Data[i]
		}
	}
	arrayList.Data = values
}

func (arrayList *ArrayList) RemoveFirst() (bool, error) {
	return arrayList.Remove(0)
}

func (arrayList *ArrayList) RemoveLast() (bool, error) {
	return arrayList.Remove(arrayList.Size - 1)
}

func (arrayList *ArrayList) RemoveValue(value int) (bool, error) {
	index := arrayList.Find(value)
	if index != -1 {
		return arrayList.Remove(index)
	} else {
		return false, errors.New("remove faild. no values to remove")
	}
}

func (arrayList *ArrayList) ToString() string {
	return fmt.Sprint(arrayList)
}

func MainArrayList() {
	var arrayList ArrayList
	for i := 0; i < 10; i++ {
		arrayList.Add(i, i+1)
	}
	fmt.Println("--------------Add---------------")
	fmt.Println(arrayList)

	fmt.Println("--------------Remove(4)---------------")
	arrayList.Remove(4)
	fmt.Println(arrayList)

	fmt.Println("--------------Add(4)---------------")
	arrayList.Add(4, 5)
	fmt.Println(arrayList)

	fmt.Println("--------------ReSetSize---------------")
	for i := 0; i < 8; i++ {
		arrayList.RemoveLast()
		fmt.Printf("len: %d, cap:%d, Values:%+v, Size: %d \n", len(arrayList.Data), cap(arrayList.Data), arrayList.Data, arrayList.Size)
	}

}

func AA() {
	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)

}
