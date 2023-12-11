import { Node, TreeNode, search } from "./src/binary-search-trees/basic-tree"
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
