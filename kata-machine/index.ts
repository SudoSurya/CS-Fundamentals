import { Node, TreeNode} from "./src/binary-search-trees/basic-tree"
import { bfs_traversal } from "./src/graphs/BFS-Traversal"
import { dfs_traversal } from "./src/graphs/DFS-Traversal"
import { constructAdjList } from "./src/graphs/adjacencyList"
import { constructAdjMatrix } from "./src/graphs/adjacencyMatrix"
import { Vertex, dfsTraversal } from "./src/graphs/graph"

console.log("hello world")

let root = new TreeNode(10)

root.insert(5, root as Node<number>)
root.insert(15, root as Node<number>)
root.insert(25, root as Node<number>)
root.insert(6, root as Node<number>)
root.delete(10, root as Node<number>)
console.log("before",root)
root.traverse(root as Node<number>)

let alice = new Vertex("alice")
let bob = new Vertex("bob")
let charlie = new Vertex("charlie")

alice.addEdge(bob)
alice.addEdge(charlie)
bob.addEdge(charlie)
charlie.addEdge(alice)
console.log(dfsTraversal(alice))
// console.log(alice)


let adjMatrix = constructAdjMatrix(3, [[0, 1], [1, 2], [2, 0]])
console.log(adjMatrix)
let adjList = constructAdjList(5, [[0, 1], [0, 2], [0, 3],[1,3],[3,2]])
console.log(adjList)


// let result = bfs_traversal(adjList)

let result = dfs_traversal(6, adjList)
console.log("dfs",result)

