package unionfind

type uf interface {
	IsConnected(p, q int) bool

	UnionElements(p, q int)

	GetSize() int
}
