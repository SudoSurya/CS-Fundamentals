export type ListNode = {
    val: number
    next: ListNode | null
}


export function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {

    let carry = 0

    let curr1 = l1
    let curr2 = l2
    let dummyHead: ListNode = {
        val: 0,
        next: null
    }
    let curr3 = dummyHead
    while (curr1 || curr2) {
        let sum = carry
        if (curr1) sum += curr1.val
        if (curr2) sum += curr2.val
        curr3.next = {
            val: sum % 10,
            next: null
        }
        carry = Math.floor(sum / 10)
        curr3 = curr3.next
        if (curr1) curr1 = curr1.next
        if (curr2) curr2 = curr2.next
    }
    if (carry > 0) {
        curr3.next = {
            val: carry,
            next: null
        }
    }
    return dummyHead.next
};
