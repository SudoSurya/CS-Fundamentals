package arrays;

public class PermutateStrings {

    static boolean solution(String s1, String s2) {
        
        // define table
        int[] arr = new int[128];
        
        // define pointers
        int L = 0, R = 0;

        char[] s1_arr = s1.toCharArray();
        char[] s2_arr = s2.toCharArray();

        for (char c : s1_arr) {
            arr[c]++;
        }
        int minLen = Integer.MAX_VALUE;
        int counter = 0;

        

        while (R < s2_arr.length) {
            char cur = s2_arr[R];
            if (--arr[cur] >= 0) {
                counter++;
            }

            while (s1.length() == counter) {
                int curLen = R - L + 1;
                minLen = Math.min(minLen, curLen);

                char LeftCur = s2_arr[L];
                if (++arr[LeftCur] > 0) {
                    counter--;
                }
                L++;
            }
            R++;
        }
        return minLen == s1.length();

    }
    public static void main(String[] args) {
        boolean res = solution("ab", "ababa");
        System.out.println(res);
    }

}
