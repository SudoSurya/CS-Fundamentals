export function lengthOfLongestSubstring(s: string): number {

    if (s.length === 1) return 1;
    let max = 0;

    let R = 0;
    let L = 0;

    let map = new Map<string, number>();
    // abcabcbb
    // ^
    while (L < s.length) {
        let char = s.charAt(L);
        if (map.has(char)) {
            map.delete(s.charAt(R));
            max = Math.max(max, L - R);
            R++;
        }else{
            map.set(char, L);
            L++;
        }
    }
    max = Math.max(max, L - R);
    return max;
};
