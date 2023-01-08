package myleetcode

type Allocator struct {
	AlreadyList map[int]*AlreadyBlock6259
	FreeList    *FreeBlock6259
	Length      int
}

type AlreadyBlock6259 struct {
	Length int
	Start  int
	End    int
	ID     int
	Prev   *AlreadyBlock6259
	Next   *AlreadyBlock6259
}

type FreeBlock6259 struct {
	Length int
	Start  int
	End    int
	Prev   *FreeBlock6259
	Next   *FreeBlock6259
}

func Constructor6259(n int) Allocator {
	head := new(FreeBlock6259)
	return Allocator{Length: n, AlreadyList: make(map[int]*AlreadyBlock6259), FreeList: head}
}

func (this *Allocator) Allocate(size int, mID int) int {
	if this.FreeList == nil {
		return -1
	}
	z := this.FreeList
	for z != nil {
		if z.Length == size {
			f := z.Next
			q := z.Prev
			q.Next = f
			f.Prev = q
			al := new(AlreadyBlock6259)
			al.Start = z.Start
			al.Length = size
			al.End = al.Start + size - 1
			al.ID = mID
			this.AlreadyList[mID] = al
			return z.Start
		} else if z.Length > size {
			nn := new(FreeBlock6259)
			nn.Start = z.Start + size
			nn.Length = z.Length - size
			nn.End = nn.Start + size - 1
			z.Prev.Next = nn
			z.Next.Prev = nn
			nn.Next = z.Next
			nn.Prev = z.Prev
			al := new(AlreadyBlock6259)
			al.Start = z.Start
			al.Length = size
			al.End = al.Start + size - 1
			al.ID = mID
			this.AlreadyList[mID] = al
			return z.Start
		}
		z = z.Next
	}
	return -1
}

func (this *Allocator) Free(mID int) int {
	al := this.AlreadyList[mID]
	z := this.FreeList
	for z != nil {
		if z.Start == al.End+1 {

		} else if al.End < z.Start {
+
		}
	}
}
