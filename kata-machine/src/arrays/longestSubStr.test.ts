import { describe, expect, test } from "@jest/globals";
import { lengthOfLongestSubstring } from "./longestSubStr";

describe("Longest Substring Without Repeating Characters", () => {
    test("Example 1", () => {
        expect(lengthOfLongestSubstring("abcabcbb")).toBe(3);
    });
    test("Example 2", () => {
        expect(lengthOfLongestSubstring("bbbbb")).toBe(1);
    });
    test("Example 3", () => {
        expect(lengthOfLongestSubstring("pwwkew")).toBe(3);
    });
    test("Example 4", () => {
        expect(lengthOfLongestSubstring("")).toBe(0);
    });
    test("Example 5", () => {
        expect(lengthOfLongestSubstring(" ")).toBe(1);
    });
    test("Example 6", () => {
        expect(lengthOfLongestSubstring("au")).toBe(2);
    });
    test("Example 7", () => {
        expect(lengthOfLongestSubstring("dvdf")).toBe(3);
    });
})
