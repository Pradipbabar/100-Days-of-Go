package main
import (
   "fmt"
)

const n = 4

// Graph represents a graph using an adjacency matrix
type Graph struct {
   AdMatrix [n][n]bool
}

// Node represents a node in the graph.
type Node struct {
	value int
	edges []*Node
 }
 
 // AddEdge adds a new edge to the node.
 func (n *Node) Add_Edge(node *Node) {
	n.edges = append(n.edges, node)
 }
 

func main() {
   // Create graph
   graph := Graph{}

   // Connect nodes
   graph.AdMatrix[0][1] = true
   graph.AdMatrix[0][2] = true
   graph.AdMatrix[1][2] = true

   // Print graph
   fmt.Println("The graph is printed as follows using adjacency matrix:")
   fmt.Println("Adjacency Matrix:")
   for i := 0; i < n; i++ {
      for j := 0; j < n; j++ {
         fmt.Printf("%t ", graph.AdMatrix[i][j])
      }
      fmt.Println()
   }

    // Create nodes.
	n1 := &Node{value: 10}
	n2 := &Node{value: 20}
	n3 := &Node{value: 30}
	n4 := &Node{value: 40}
	n5 := &Node{value: 50}
 
	// Add edges.
	n1.Add_Edge(n2)
	n1.Add_Edge(n3)
	n2.Add_Edge(n4)
	n2.Add_Edge(n5)
	n3.Add_Edge(n5)
 
	// Display the graph.
	fmt.Println("The Graph is represented as:")
	fmt.Printf("Node %d -> %v\n", n1.value, getNodeValues(n1.edges))
	fmt.Printf("Node %d -> %v\n", n2.value, getNodeValues(n2.edges))
	fmt.Printf("Node %d -> %v\n", n3.value, getNodeValues(n3.edges))
	fmt.Printf("Node %d -> %v\n", n4.value, getNodeValues(n4.edges))
	fmt.Printf("Node %d -> %v\n", n5.value, getNodeValues(n5.edges))
 }
 
 // returns a slice of node values.
 func getNodeValues(nodes []*Node) []int {
	var values []int
	for _, node := range nodes {
	   values = append(values, node.value)
	}
	return values
}