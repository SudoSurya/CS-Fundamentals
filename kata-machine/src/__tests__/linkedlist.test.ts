import { describe, expect, test } from "@jest/globals";
import { LinkedList } from "../linkedlist/linkedList";

const linkedList = new LinkedList<number>()
linkedList.add(1)
linkedList.add(2)

describe("Linked List", () => {
    test("Add", () => {
        expect(linkedList).toBeInstanceOf(LinkedList)
        expect(linkedList.log()).toMatch("1\n2")
    })
    test("Length", () => {
        expect(linkedList.length).toBe(2)
    })
    test("Search", () => {
        expect(linkedList.search(1)).toBe(true)
        expect(linkedList.search(2)).toBe(true)
        expect(linkedList.search(3)).toBe(false)
    })
    test("Remove Tail", () => {
        linkedList.removeTail()
        expect(linkedList.log()).toMatch("1")
    })
    test("Delete", () => {
        linkedList.delete(1)
        expect(linkedList.log()).toMatch("")
    })
})
