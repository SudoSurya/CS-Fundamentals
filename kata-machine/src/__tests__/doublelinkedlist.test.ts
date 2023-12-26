// write tests for the double linked list here using jest

// Path: doublelinkedlist.ts

import { describe, expect, test } from "@jest/globals";
import { DoublyLinkedList } from "../linkedlist/double-linkedlist";

// TODO
describe("Doubly Linked List", () => {
    test("should create empty linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        expect(linkedList.log()).toBe("0\n");
    })

    test("should append node to linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        expect(linkedList.log()).toBe("0\n1\n2\n");
    });

    test("should prepend node to linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.prepend(1);
        linkedList.prepend(2);
        expect(linkedList.log()).toBe("2\n1\n0\n");
    })

    test("should delete head node from linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        linkedList.deleteHead();
        expect(linkedList.log()).toBe("1\n2\n");
    })

    test("should delete tail node from linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        linkedList.deleteTail();
        expect(linkedList.log()).toBe("0\n1\n");
    })

    test("should convert array to linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        const array = [1, 2, 3, 4, 5];
        linkedList.convertArrayToLinkedList(array);
        expect(linkedList.log()).toBe("0\n1\n2\n3\n4\n5\n");
    });

    test("should convert linked list to array", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        let res = linkedList.convertLinkedListToArray();
        expect(res).toEqual([0, 1, 2]);
    })

    test("should delete node by value from linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        linkedList.deleteNodeByValue(1);
        expect(linkedList.log()).toBe("0\n2\n");
    })

    test("should delete node by value from linked list", () => {
        const linkedList = new DoublyLinkedList(0);
        linkedList.append(1);
        linkedList.append(2);
        linkedList.deleteNodeByValue(1);
        expect(linkedList.log()).toBe("0\n2\n");
    })
});

