package strings;

import java.util.HashMap;

public class LongRep {
    public int lengthOfLongestSubstring(String s) {

        char arr[] = s.toCharArray();

        int n = arr.length;
        int max = 0;

        // define hashmap key as string value as index
        HashMap<Character, Integer> map = new HashMap<>();

        int R = 0;
        // dvddf
        // ^
        int count = 0;
        while (R < n) {
            char ch = arr[R];

            if (map.containsKey(ch)) {
                int idx = map.get(ch);
                count = R - idx;
                max = Math.max(max, count);
                map.clear();
                count = 0;
            }
            map.put(arr[R], R);
            R++;
            count++;
        }

        return Math.max(max, count);

    }

    public static void main(String[] args) {
        LongRep lr = new LongRep();
        System.out.println(lr.lengthOfLongestSubstring(" "));
    }

}
