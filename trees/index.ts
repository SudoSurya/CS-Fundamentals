import { Node, TreeNode, search } from "./src/binary-search-trees/basic-tree"

console.log("hello world")

let node1 = new TreeNode(6) as Node<number>
let node2 = new TreeNode(20) as Node<number>
let root = new TreeNode(10, node1, node2)

let node4 = new TreeNode(4)
let node5 = new TreeNode(8)
node1.left = node4 as Node<number>
node1.right = node5 as Node<number>

let node6 = new TreeNode(15)
let node7 = new TreeNode(22)
node2.left = node6 as Node<number>
node2.right = node7 as Node<number>
root.insert(9, root as Node<number>)
root.insert(150, root as Node<number>)

console.log(root)

console.log(search(9, root as Node<number>))
console.log(search(150, root as Node<number>)?.val)
