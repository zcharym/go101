package list

import (
	"fmt"
	"testing"
)

func TestListNew(t *testing.T) {
	list1 := New()

	if actualValue := list1.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	list2 := New(1, "b")

	if actualValue := list2.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	if actualValue, ok := list2.Get(0); actualValue != 1 || !ok {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	if actualValue, ok := list2.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}

	if actualValue, ok := list2.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
}

func TestListAdd(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
}

func TestListIndexOf(t *testing.T) {
	list := New()

	expectedIndex := -1
	if index := list.IndexOf("a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	list.Add("a")
	list.Add("b", "c")

	expectedIndex = 0
	if index := list.IndexOf("a"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	expectedIndex = 1
	if index := list.IndexOf("b"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}

	expectedIndex = 2
	if index := list.IndexOf("c"); index != expectedIndex {
		t.Errorf("Got %v expected %v", index, expectedIndex)
	}
}

func TestListRemove(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	list.Remove(2)
	if actualValue, ok := list.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.Remove(1)
	list.Remove(0)
	list.Remove(0) // no effect
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListGet(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "c")
	}
	if actualValue, ok := list.Get(3); actualValue != nil || ok {
		t.Errorf("Got %v expected %v", actualValue, nil)
	}
	list.Remove(0)
	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v expected %v", actualValue, "b")
	}
}

func TestListClear(t *testing.T) {
	list := New()
	list.Add("e", "f", "g", "a", "b", "c", "d")
	list.Clear()
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
}

func TestListValues(t *testing.T) {
	list := New()
	list.Add("a")
	list.Add("b", "c")
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s", list.Values()...), "abc"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestListInsert(t *testing.T) {
	list := New()
	list.Insert(0, "b", "c")
	list.Insert(0, "a")
	list.Insert(10, "x") // ignore
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Insert(3, "d") // append
	if actualValue := list.Size(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s", list.Values()...), "abcd"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestListSet(t *testing.T) {
	list := New()
	list.Set(0, "a")
	list.Set(1, "b")
	if actualValue := list.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	list.Set(2, "c") // append
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	list.Set(4, "d")  // ignore
	list.Set(1, "bb") // update
	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue, expectedValue := fmt.Sprintf("%s%s%s", list.Values()...), "abbc"; actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func benchmarkGet(b *testing.B, list *SqList, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Get(n)
		}
	}
}

func benchmarkAdd(b *testing.B, list *SqList, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Add(n)
		}
	}
}

func benchmarkRemove(b *testing.B, list *SqList, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Remove(n)
		}
	}
}

func BenchmarkArrayListGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListGet1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New()
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkArrayListAdd1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkArrayListAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkArrayListAdd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkArrayListRemove100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkArrayListRemove1000(b *testing.B) {
	b.StopTimer()
	size := 1000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkArrayListRemove10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}

func BenchmarkArrayListRemove100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New()
	for n := 0; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkRemove(b, list, size)
}
