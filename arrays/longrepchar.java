package arrays;

import java.util.HashMap;

public class longrepchar {

    static int solution(char[] arr, int k) {
        int left = 0, right = 0, n = arr.length;

        HashMap<Character, Integer> hm = new HashMap<>();

        int maxLen = 0, mostFreq = 0;

        while (right < n) {

            hm.put(arr[right], hm.getOrDefault(arr[right], 0) + 1);
            mostFreq = Math.max(mostFreq, hm.get(arr[right]));

            while ((right - left + 1) - mostFreq > k) {
                hm.put(arr[left], hm.get(arr[left]) - 1);
                left++;
            }
            maxLen = Math.max(maxLen, right - left + 1);
            right++;
        }
        return maxLen;
    }

    public static void main(String[] args) {
        String str = "AABABBA";
        int result = solution(str.toCharArray(), 1);
        System.out.println(result);
        System.out.println(result);
    }
}
