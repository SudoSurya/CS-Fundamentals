
export class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.next = (next === undefined ? null : next)
    }
}


export function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {
    let length = 0;
    let currNode = head;
    while (currNode) {
        length++;
        currNode = currNode.next;
    }
    if (length === 1) {
        return null;
    }
    let deletePos = 0

    let curr = head;
    let prev = null;
    while (deletePos !== length - n) {
        console.log("deletePos", deletePos);
        console.log("curr", curr?.val);
        prev = curr;
        if (curr) {
            curr = curr.next
        }
        deletePos++;
    }
    if (prev !== null) {
        prev.next = curr?.next || null;
    } else {
        head = curr?.next || null;
    }
    return head;
};
