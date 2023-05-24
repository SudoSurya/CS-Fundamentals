package arrays;

import java.util.HashMap;

public class substrswt {

    static int lenOfLongSubstrTwoDistict(char[] arr,int k) {
        int n = arr.length;
        if (n == k)
            return n;
        int max = 0;
        int L = 0, R = 0;
        HashMap<Character, Integer> hm = new HashMap<>();
        while (R < n) {
            hm.put(arr[R], hm.getOrDefault(arr[R], 0) + 1);
            while (hm.size() > k) {
                hm.put(arr[L], hm.get(arr[L]) - 1);
                if (hm.get(arr[L]) == 0) {
                    hm.remove(arr[L]);
                }
                L++;
            }
            max = Math.max(max, R - L + 1);
            R++;
        }
        return max;
    }

    public static void main(String[] args) {
        String str = "eceba";
        int k = 2;
        int res = lenOfLongSubstrTwoDistict(str.toCharArray(),k);
        System.out.println(res);
    }
}
