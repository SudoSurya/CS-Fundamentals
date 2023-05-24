package arrays;

import java.util.HashMap;

/**
 * maps
 */
public class maps {

    public static int LongSubStr(char args[]) {
        int n = args.length;
        if (n < 2)
            return n;

        int L = 0, R = 0;
        int max = Integer.MIN_VALUE;

        HashMap<Character, Integer> hm = new HashMap<>();
        // abcabcde
        // {a: 1,b:1,c:1} r=4 l=2

        while (R < n) {
            hm.put(args[R], hm.getOrDefault(args[R], 0) + 1);
            while (hm.size() != R - L + 1) {
                hm.put(args[L], hm.get(args[L]) - 1);
                if (hm.get(args[L]) == 0) {
                    hm.remove(args[L]);
                }
                L++;
            }
            max = Math.max(max, R - L + 1);
            R++;
        }
        return max;

    }

    public static void main(String[] args) {

        HashMap<Character, Integer> numsObj = new HashMap<>();

        char[] arrs = { 'a', 'b', 'b', 'd', 'b', 'b', 'a', 'a', 'a' };

        for (int i = 0; i < arrs.length; i++) {
            if (numsObj.containsKey(arrs[i])) {
                numsObj.put(arrs[i], numsObj.get(arrs[i]) + 1);
            } else {
                numsObj.put(arrs[i], 1);
            }


        }
        System.out.println(numsObj.size());
        System.out.println(LongSubStr(arrs));

    }
}