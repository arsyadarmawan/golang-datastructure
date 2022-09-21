package graph

import "fmt"

// graph represent ajency list
type Graph struct {
	Vertices []*Vertex
}

// vertex represents a graph vertex
type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(k int) {
	if isDuplicate(k, g.Vertices) {
		return
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Key: k})
	}
}

// Print will print the adjacent list for each vertex to the graph
func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("\nVertex %v :", v.Key)
		for _, v := range v.Adjacent {
			fmt.Printf(" %v", v.Key)
		}
	}
	fmt.Println()
}

// validate vertex is available or not
func isDuplicate(k int, s []*Vertex) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

// AddEdgge adds an edge to the graph
func (g *Graph) AddEdgge(from, to int) {
	// get vertex
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge ( %v --> %v )", from, to)
		fmt.Println(err.Error())
	} else if isDuplicate(to, fromVertex.Adjacent) {
		err := fmt.Errorf("existing edge (%v -> %v )", from, to)
		fmt.Println(err)
	} else {
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
	}
	// add edge

}

func (g *Graph) GetVertex(k int) *Vertex {
	for i, v := range g.Vertices {
		if v.Key == k {
			return g.Vertices[i]
		}
	}
	return nil
}
