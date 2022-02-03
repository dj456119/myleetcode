/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2022-01-27 00:50:16
 * @LastEditors: cm.d
 * @LastEditTime: 2022-02-02 11:25:18
 */
package myleetcode

type ListNode146 struct {
	Prev  *ListNode146
	Next  *ListNode146
	Key   int
	Value int
}

type LRUCache struct {
	Capacity  int
	Cache     map[int]*ListNode146
	ValueList *ListNode146
	LastNode  *ListNode146
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Capacity: capacity, Cache: make(map[int]*ListNode146)}
}

func (this *LRUCache) Get(key int) int {
	valueList, ok := this.Cache[key]
	if !ok {
		return -1
	}
	if valueList == this.ValueList {
		return valueList.Value
	}
	if valueList.Prev != nil {
		valueList.Prev.Next = valueList.Next
		if this.LastNode.Key == valueList.Key {
			this.LastNode = valueList.Prev
		}
	}
	if valueList.Next != nil {
		valueList.Next.Prev = valueList.Prev
	}
	if this.ValueList != nil {
		this.ValueList.Prev = valueList
	}
	valueList.Next = this.ValueList
	this.ValueList = valueList
	return valueList.Value
}

func (this *LRUCache) Put(key int, value int) {
	if valueList, ok := this.Cache[key]; ok {
		valueList.Value = value
		this.Get(key)
	} else {
		valueList := new(ListNode146)
		valueList.Value = value
		valueList.Key = key
		valueList.Next = this.ValueList
		if valueList.Next != nil {
			valueList.Next.Prev = valueList
		}
		this.ValueList = valueList
		this.Cache[key] = valueList

		if this.LastNode == nil {
			this.LastNode = valueList
		}
		if len(this.Cache) > this.Capacity {
			prev := this.LastNode.Prev
			delete(this.Cache, this.LastNode.Key)
			prev.Next = nil
			this.LastNode = prev

		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
