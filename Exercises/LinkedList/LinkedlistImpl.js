class Node {
    constructor(data) {
        this.data = data;
        this.next = null;
    }
}

class LinkedList {
    constructor(head = null) {
        this.head = head;
    }

    read(index) {
        let current = this.head;
        let CurrentIndex = 0;

        while (CurrentIndex < index) {
            current = current.next;
            CurrentIndex++;
            if (current === null) {
                return "Error: Index out of bounds";
            }
        }
        return current.data;
    }
    indexOf(value) {
        let current = this.head;
        let currentIndex = 0;

        while (current !== null) {
            if (current.data === value) {
                return currentIndex + " Search Data";
            }
            current = current.next;
            currentIndex++;
        }
        return -1;
    }
    insertAtIndex(index, value) {
        let newNode = new Node(value)

        if (index === 0) {
            newNode.next = this.head;
            this.head = newNode;
            return "Inserted" + index;
        }
        let current = this.head;
        let currentIndex = 0;

        while (currentIndex < index - 1) {
            current = current.next;
            currentIndex++;
        }
        newNode.next = current.next;
        current.next = newNode;
        return "Inserted" + index;
    }
    deleteAtIndex(index) {
        if (index == 0) {
            this.head = this.head.next;
            return "Deleted" + index;
        }
        let current = this.head;
        let currentIndex = 0;
        while (currentIndex < index - 1) {
            current = current.next;
            currentIndex++;
        }
        let NextNode = current.next?.next;
        current.next = NextNode;
        return "Deleted" + index;
    }
}
// Path: LinkedlistImpl.js

let node1 = new Node(1);
let node2 = new Node(2);
let node3 = new Node(3);
node1.next = node2;
node2.next = node3;

let list = new LinkedList(node1);
console.log(list);
console.log(list.read(2));
console.log(list.indexOf(3));
console.log(list.insertAtIndex(2, 5));
console.log(list.deleteAtIndex(2));
console.log(list);
