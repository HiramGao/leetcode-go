package main

//type NestedInteger struct {
//}
//
//func (i NestedInteger) GetInteger() int {
//	return 0
//}
//
//func (i NestedInteger) IsInteger() bool {
//	return bool()
//}
//
//func (i NestedInteger) GetList() []*NestedInteger {
//	return &[]NestedInteger{}
//}
//
//type NestedIterator struct {
//	stack [][]*NestedInteger
//}
//
//func Constructor(nestedList []*NestedInteger) *NestedIterator {
//	return &NestedIterator{[][]*NestedInteger{nestedList}}
//}
//
//func (this *NestedIterator) Next() int {
//	queue := this.stack[len(this.stack)-1]
//	val := queue[0].GetInteger()
//	this.stack[len(this.stack)-1] = queue[1:]
//	return val
//
//}
//
//func (it *NestedIterator) HasNext() bool {
//	for len(it.stack) > 0 {
//		queue := it.stack[len(it.stack)-1]
//		if len(queue) == 0 { // 当前队列为空，出栈
//			it.stack = it.stack[:len(it.stack)-1]
//			continue
//		}
//		nest := queue[0]
//		if nest.IsInteger() {
//			return true
//		}
//		// 若队首元素为列表，则将其弹出队列并入栈
//		it.stack[len(it.stack)-1] = queue[1:]
//		it.stack = append(it.stack, nest.GetList())
//	}
//	return false
//
//}
func main() {

}
