package _struct

type UnionFind struct {
	Parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		Parent: parent,
	}
}

func (u *UnionFind) Find(i int) int {
	root := i
	for u.Parent[root] != root {
		root = u.Parent[root]
	}
	for u.Parent[i] != root {
		i, u.Parent[i] = u.Parent[i], root
	}
	return root
}

func (u *UnionFind) Union(i, j int) {
	pi := u.Find(i)
	pj := u.Find(j)
	if pi != pj {
		u.Parent[pi] = pj
	}
}

func (u *UnionFind) IsConnected(i, j int) bool {
	return u.Find(i) == u.Find(j)
}
