import { DoublyLinkedList } from "./src/linkedlist/double-linkedlist";

const head = new DoublyLinkedList<number>(0)
head.insertAtEnd(1)
head.insertAtEnd(2)
head.insertAtEnd(3)
head.insertAtEnd(4)
console.log(head.log())
head.deleteHead()
head.deleteTail()

