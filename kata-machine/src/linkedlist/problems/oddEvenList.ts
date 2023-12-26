class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val === undefined ? 0 : val)
        this.next = (next === undefined ? null : next)
    }
}

export function oddEvenList(head: ListNode | null): ListNode | null {

    if (!head || !head.next) return head
    let odd = head
    let even = head.next
    let evenHead = even
    while (even && even.next) {
        odd.next = even.next
        odd = odd.next
        even.next = odd.next
        even = even.next as ListNode
    }
    odd.next = evenHead
    return head
};

export function zerosOnesTwos(head: ListNode | null): ListNode | null {

    let zeroHead = new ListNode(0)
    let oneHead = new ListNode(0)
    let twoHead = new ListNode(0)
    let zero = zeroHead
    let one = oneHead
    let two = twoHead
    let curr = head
    while (curr) {
        if (curr.val == 0) {
            zero.next = curr
            zero = zero.next
        } else if (curr.val == 1) {
            one.next = curr
            one = one.next
        } else {
            two.next = curr
            two = two.next
        }
        curr = curr.next
    }
    zero.next = oneHead.next ? oneHead.next : twoHead.next
    one.next = twoHead.next
    two.next = null
    return zeroHead.next
}


