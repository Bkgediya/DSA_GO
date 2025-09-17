package graph

import "DSA_Go/linked_list/models"

// time complexity: O(V + E)
// space complexity: O(V)
func CloneGraph(node *models.GraphNode) *models.GraphNode {
	if node == nil {
		return nil
	}

	cloneMap := map[*models.GraphNode]*models.GraphNode{}

	return dfs(node, cloneMap)
}

func dfs(node *models.GraphNode, cloneMap map[*models.GraphNode]*models.GraphNode) *models.GraphNode {
	// base case
	if cloned, ok := cloneMap[node]; ok {
		return cloned
	}

	// # clone the current nodecloneMap
	cloneNode := &models.GraphNode{
		Val: node.Val,
	}

	// # add to cloneMap
	cloneMap[node] = cloneNode

	for _, neighbor := range node.Neighbors {
		cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(neighbor, cloneMap))
	}

	return cloneNode
}
