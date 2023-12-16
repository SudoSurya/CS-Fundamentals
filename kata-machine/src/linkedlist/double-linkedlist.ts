type Node<T> = {
    value: T,
    next: Node<T> | null,
    prev: Node<T> | null
}
export class DoublyLinkedList<T>{
    private head: Node<T> | null;

    constructor(value: T, next: Node<T> | null = null, prev: Node<T> | null = null) {
        this.head = {
            value: value,
            next: next,
            prev: prev
        }
    }

    convertArrayToLinkedList(data: T[]): Node<T> {
        let current = this.head
        let prev: Node<T> | null = current
        for (let i = 0; i < data.length; i++) {
            const newNode: Node<T> = {
                value: data[i],
                prev: prev,
                next: null
            }
            if (prev) prev.next = newNode
            prev = newNode as Node<T>
        }
        return current as Node<T>
    }

    deleteHead(): Node<T> | null {
        if (!this.head) return null
        if (this.head.next == null) {
            this.head = null
            return null
        }


        return this.head = this.head.next
    }
    deleteTail(): Node<T> | null {
        if(!this.head || this.head.next == null) return this.head = null
        let current = this.head
        while (current.next) {
            current = current.next
        }
        if (current.prev){
            current.prev.next = null
        }
        return this.head
    }
    log(): string {
        let current = this.head
        let str = ""
        while (current) {
            str += current.value + "\n"
            current = current.next as Node<T>
        }
        return str
    }
    insertAtEnd(value: T): void {
        if (!this.head) return
        let current = this.head
        while (current.next) {
            current = current.next
        }
        const newNode: Node<T> = {
            value: value,
            next: null,
            prev: current
        }
        current.next = newNode
    }
}
