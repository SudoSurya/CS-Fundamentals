// write tests for the double linked list here using jest

// Path: doublelinkedlist.ts

import { describe, expect, test } from "@jest/globals";
import { DoublyLinkedList } from "../linkedlist/double-linkedlist";

// TODO
describe("Doubly Linked List", () => {
    test.skip("should create empty linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        expect(linkedList.toString()).toBe("");
    })

    test.skip("should append node to linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        expect(linkedList.toString()).toBe("0,1,2");
    });

    test.skip("should prepend node to linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.prepend(1);
        linkedList.prepend(2);
        expect(linkedList.toString()).toBe("2,1,0");
    })

    test.skip("should delete head node from linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        linkedList.deleteHead();
        expect(linkedList.toString()).toBe("1,2");
    })

    test.skip("should delete tail node from linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        linkedList.deleteTail();
        expect(linkedList.toString()).toBe("0,1");
    })

    test.skip("should convert array to linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        const array = [1, 2, 3, 4, 5];
        linkedList.convertArrayToLinkedList(array);
        expect(linkedList.toString()).toBe("0,1,2,3,4,5");
    });

    test.skip("should convert linked list to array", () => {
        const linkedList = new DoublyLinkedList(0);
        const array = [1, 2, 3, 4, 5];
        linkedList.convertArrayToLinkedList(array);
        expect(linkedList.convertLinkedListToArray()).toEqual(array);
    })

    test.skip("should delete node by value from linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        linkedList.deleteNodeByValue(1);
        expect(linkedList.toString()).toBe("0,2");
    })

    test.skip("should delete node by value from linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        linkedList.deleteNodeByValue(1);
        expect(linkedList.toString()).toBe("0,2");
    })
});

