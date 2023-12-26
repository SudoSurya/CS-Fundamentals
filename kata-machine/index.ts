import { ListNode, removeNthFromEnd } from "./src/linkedlist/problems/removeNthNode";


// Path: index.ts
let head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
head.next.next.next = new ListNode(4);
head.next.next.next.next= new ListNode(5);

printList(head);
let result = removeNthFromEnd(head, 2);
console.log(result);
printList(result);

function printList(head: ListNode | null) {
    let curr = head;
    while (curr) {
        console.log(curr.val);
        curr = curr.next;
    }
}
