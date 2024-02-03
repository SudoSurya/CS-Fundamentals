import { describe, expect, test } from '@jest/globals';
import { longestPalindrome } from './longestPalindrome';

describe("Longest Palindromic Substring", () => {
    test("Example 1", () => {
        expect(longestPalindrome("babad")).toBe("bab");
    });
    test("Example 2", () => {
        expect(longestPalindrome("cbbd")).toBe("bb");
    });
    test("Example 3", () => {
        expect(longestPalindrome("a")).toBe("a");
    })
    test("Example 4", () => {
        expect(longestPalindrome("ac")).toBe("a");
    })
    test("Example 5", () => {
        expect(longestPalindrome("bb")).toBe("bb");
    })
    test("Example 6", () => {
        expect(longestPalindrome("ccc")).toBe("ccc");
    })
    test("Example 7", () => {
        expect(longestPalindrome("aaaa")).toBe("aaaa");
    })
});
