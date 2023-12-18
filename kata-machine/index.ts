import { DoublyLinkedList } from "./src/linkedlist/double-linkedlist";
import { ListNode, addTwoNumbers } from "./src/linkedlist/problems/addTwoNumbers";

const head = new DoublyLinkedList<number>(0)
head.insertAtEnd(1)
head.insertAtEnd(2)
head.insertAtEnd(3)
head.insertAtEnd(4)
console.log(head.log())
head.reverse()
console.log(head.log())
head.reverseSwapping()
console.log(head.log())

console.log("-------------------");
let l1: ListNode = {
    val: 2,
    next: {
        val: 4,
        next: {
            val: 3,
            next: null
        }
    }
}

let l2: ListNode = {
    val: 5,
    next: {
        val: 6,
        next: {
            val: 4,
            next: null
        }
    }
}

console.log(addTwoNumbers(l1, l2))

