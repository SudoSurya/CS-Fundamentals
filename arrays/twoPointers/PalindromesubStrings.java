package arrays.twoPointers;

public class PalindromesubStrings {

    public static int countSubstrings(String s) {
        // a a a a
        // ^
        // ^
        int left = 0, right = left + 1;
        int count = 0;
        char[] chars = s.toCharArray();
        while (left < chars.length) {
            while (right < chars.length) {
                String s1 = "";
                String s2 = "";
                for (int i = left; i <= right; i++) {
                    s1 += chars[i];
                }
                for (int i = right; i >= left; --i) {
                    s2 += chars[i];
                }
                if (s1.equals(s2)) {
                    count += 1;
                }
                System.out.println(s2);
                right++;
            }
            left++;
            right = left + 1;
        }
        return count;
    }

    public static void main(String[] args) {
        System.out.println(countSubstrings("aaa"));
    }

}
