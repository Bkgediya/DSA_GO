package graph

// pattern: disjoint set

// when to use
// asked to merge groups dynamically.
// if two nodes belong to the same group.
// need properties of connected components (size, leader, etc.).
// Problems like: Friend circles, merging communities, accounts merge, network connectivity.

// Union-Find
// Union: merge two groups.

type DSU struct {
	parents []int
	sizes   []int
}

func NewDSU(n int) *DSU {
	parents := make([]int, n+1)
	sizes := make([]int, n+1)

	for i := 0; i < n; i++ {
		parents[i] = i
		sizes[i] = 1
	}
	return &DSU{
		parents: parents,
		sizes:   sizes,
	}
}

func (d *DSU) FindParent(num int) int {
	// path compression
	if num != d.parents[num] {
		d.parents[num] = d.FindParent(d.parents[num])
	}

	return d.parents[num]
}

func (d *DSU) Union(num1 int, num2 int) {
	ra := d.FindParent(num1)
	rb := d.FindParent(num2)

	if ra == rb {
		return
	}

	if d.sizes[ra] < d.sizes[rb] {
		ra, rb = rb, ra
	}
	d.parents[rb] = ra
	d.sizes[ra] += d.sizes[rb]
}

func (d *DSU) GetCommunitySize(x int) int {
	r := d.FindParent(x)
	return d.sizes[r]
}
