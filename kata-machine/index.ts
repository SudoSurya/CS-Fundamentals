import { zerosOnesTwos } from "./src/linkedlist/problems/oddEvenList";
import { ListNode } from "./src/linkedlist/problems/removeNthNode";

let head = new ListNode(1)
head.next = new ListNode(2)
head.next.next = new ListNode(0)
head.next.next.next = new ListNode(1)
head.next.next.next.next = new ListNode(2)
head.next.next.next.next.next = new ListNode(0)

let res = zerosOnesTwos(head)

let curr = res
while (curr) {
    let val = curr.val
    console.log(val)
    curr = curr.next
}





