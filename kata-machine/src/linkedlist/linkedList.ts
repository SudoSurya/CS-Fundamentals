type Node<T> = {
    value: T;
    next: Node<T> | null;
}

export class LinkedList<T>{
    private head: Node<T> | null;
    public length: number = 0;

    constructor() {
        this.head = null;
    }

    addNodeAtFirst(value: T) {
        const newNode: Node<T> = {
            value: value,
            next: this.head
        }
        this.head = newNode;
    }

    insertAtSpecificPosition(value: T, position: number) {
        if (position < 1 || position > this.length) {
            return;
        }
        if (position === 1 || this.head === null) {
            this.addNodeAtFirst(value);
            return;
        }
        const newNode: Node<T> = {
            value: value,
            next: null
        }

        let current = this.head;
        let prev: Node<T> | null = null;
        let count = 0;

        while (current) {
            count++;
            if (count === position) {
                if (prev) {
                    prev.next = newNode
                    newNode.next = current;
                    return;
                }
            }
            prev = current;
            current = current.next as Node<T>;
        }
    }

    add(value: T) {
        const newnode: Node<T> = {
            value: value,
            next: null
        }
        if (!this.head) {
            this.head = newnode;
            this.length++;
            return;
        }
        let current = this.head;
        while (current.next) {
            current = current.next;
        }
        current.next = newnode;
        this.length++;
    }
    delete(value: T): Node<T> | null {
        if (!this.head) {
            return null;
        }

        if (this.head.value === value) {
            this.head = this.head.next;
            this.length--;
            return this.head;
        }

        let current = this.head;
        while (current.next) {
            if (current.value == value) {
                current.next = current.next.next;
                this.length--;
                return current;
            }
            current = current.next;
        }
        return null;
    }
    removeTail() {
        if (!this.head || !this.head.next) {
            return null
        }
        let current = this.head;
        while (current.next?.next) {
            current = current.next;
        }
        current.next = null;
    }
    search(value: T): Boolean {
        let current = this.head;
        while (current) {
            if (current.value === value) {
                return true;
            }
            current = current.next;
        }
        return false;
    }
    removeKthNode(k: number) {
        if (this.head === null) {
            return;
        }
        if (k === 1) {
            this.head = this.head.next;
            return;
        }
        let current = this.head;
        let prev: Node<T> | null = null
        let count = 0;
        while (current) {
            count++;
            if (count == k) {
                if (prev) {
                    prev.next = prev.next?.next as Node<T>;
                }
                return;
            }
            prev = current;
            current = current.next as Node<T>;
        }
    }
    log() {
        let current = this.head;
        let result = [];
        while (current) {
            if (current.value !== undefined) {
                result.push(current.value);
            }
            current = current.next;
        }
        return result.join("\n");
    }
    removeElements(val: number): Node<T> | null {
        if (!this.head) {
            return null
        }

        let curr = this.head
        let prev: Node<T> | null = null
        while (curr) {
            console.log(curr.value)
            if (curr.value === val) {
                if (prev) {
                    prev.next = curr.next
                    prev = curr.next
                    curr = curr.next as Node<T>
                    break
                }
            }
            prev = curr
            curr = curr.next as Node<T>
        }
        return this.head
    }
}

