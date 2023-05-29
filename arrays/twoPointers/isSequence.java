package arrays.twoPointers;

public class isSequence {
    public static boolean isSubsequence(String s, String t) {
        // ahbgdc **acc**
        if (t.length() < s.length()) {
            return false;
        }
        if (s.length() < 1) {
            return true;
        }

        char[] s1 = s.toCharArray();
        char[] t1 = t.toCharArray();

        int R = 0;
        int L = 0;
        while (R < t1.length) {
            if (t1[R] == s1[L]) {
                L++;
            }
            R++;
            if (L == s.length()) {
                return true;
            }
        }
        return false;

    }

    public static void main(String[] args) {
        System.out.println(isSubsequence("abc", "ahnbdhc"));
    }
}
