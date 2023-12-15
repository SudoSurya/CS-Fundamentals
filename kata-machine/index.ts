import { LinkedList } from "./src/linkedlist/linkedList"

console.log("Linked list")
let temp = [1, 2, 3,4, 4, 5]
const linkedList = new LinkedList<number>()

for (let i = 0; i < temp.length; i++) {
    linkedList.add(temp[i])
}

linkedList.removeKthNode(3)
linkedList.addNodeAtFirst(10)
linkedList.insertAtSpecificPosition(20, 6)
console.log(linkedList.log())
console.log("-------------------------")
linkedList.removeElements(4)
linkedList.removeElements(2)
console.log(linkedList.log())
