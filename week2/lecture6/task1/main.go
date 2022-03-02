package main

import "fmt"

type MagicList struct {
	LastItem *Item
}

type Item struct {
	Value    int
	PrevItem *Item
}

func add(l *MagicList, value int) {
	v := Item{value, nil}
	if l.LastItem == nil {
		v.PrevItem = &v
		l.LastItem = &v
	} else {
		prev := &l.LastItem
		v.PrevItem = *prev
		l.LastItem = &v

	}
}

func toSlice(l *MagicList) []int {
	magicSlice := []int{l.LastItem.Value}

	currentItem := l.LastItem
	for currentItem != currentItem.PrevItem {
		magicSlice = append([]int{currentItem.PrevItem.Value}, magicSlice...)
		currentItem = currentItem.PrevItem
	}
	return magicSlice
}

func main() {
	magic := MagicList{}
	add(&magic, 10)
	add(&magic, 20)
	add(&magic, 30)
	add(&magic, 40)
	add(&magic, 50)
	add(&magic, 60)
	add(&magic, 70)
	add(&magic, 80)
	add(&magic, 90)
	add(&magic, 100)

	magicSlice := toSlice(&magic)

	fmt.Println(magicSlice)
}
