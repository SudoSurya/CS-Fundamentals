type Node<T> = {
    value: T;
    next: Node<T> | null;
}

export class LinkedList<T>{
    private head: Node<T> | null;

    constructor() {
        this.head = null;
    }

    add(value: T) {
        const newnode: Node<T> = {
            value: value,
            next: null
        }
        if (!this.head) {
            this.head = newnode;
            return;
        }
        let current = this.head;
        while (current.next) {
            current = current.next;
        }
        current.next = newnode;
    }
}
