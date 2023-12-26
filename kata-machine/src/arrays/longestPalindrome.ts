export function longestPalindrome(s: string): string {
    let max = "";
    let maxLen = 0;
    for (let i = 0; i < s.length; i++) {
        for(let j = i; j < s.length; j++){
            let subStr = s.substring(i, j+1);
            if(isPalindrome(subStr) && subStr.length > maxLen){
                max = subStr;
                maxLen = subStr.length;
            }
        }
    }
    return max;
};

function isPalindrome(s: string): boolean {
    let L = 0;
    let R = s.length - 1;
    while (L < R) {
        if (s.charAt(L) !== s.charAt(R)) {
            return false;
        }
        L++;
        R--;
    }
    return true;
}
