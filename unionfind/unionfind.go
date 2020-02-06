package unionfind

type UnionFind struct {
	parent []int
	sz     []int
}

func assertImplementation() {
	var _ uf = (*UnionFind)(nil)
}

func NewUnionFind(size int) *UnionFind {

	//data := make([]int, size)
	parent := make([]int, size)
	sz := make([]int, size)

	for i := 0; i < size; i++ {
		parent[i] = i
		sz[i] = 1
	}
	return &UnionFind{
		parent: parent,
		sz:     sz,
	}
}

func (u *UnionFind) GetSize() int {
	return len(u.parent)
}

func (u *UnionFind) find(p int) int {
	//find的时候进行路径压缩
	for u.parent[p] != p {
		p, u.parent[p] = u.parent[p], u.parent[u.parent[p]]
	}
	return u.parent[p]
}

func (u *UnionFind) UnionElements(p, q int) {
	//取p，q两点的根节点
	parentP := u.find(p)
	parentQ := u.find(q)

	//判断深度
	if u.sz[parentP] < u.sz[parentQ] {
		u.parent[p] = u.sz[parentQ]
	} else if u.sz[parentP] > u.sz[parentQ] {
		u.parent[parentQ] = u.sz[parentP]
	} else {
		u.parent[p] = u.sz[parentQ]
		u.sz[parentQ]++
	}
}

func (u *UnionFind) IsConnected(p, q int) bool {

	//取p，q两点的根节点
	parentP := u.find(p)
	parentQ := u.find(q)

	return u.parent[parentQ] == u.parent[parentP]
}
