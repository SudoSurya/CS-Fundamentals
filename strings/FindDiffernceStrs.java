package strings;

import java.util.Arrays;

public class FindDiffernceStrs {
    public char findTheDifference(String s, String t) {
        char s1[] = s.toCharArray();
        char t1[] = t.toCharArray();
        Arrays.sort(s1);
        Arrays.sort(t1);

        for (int i = 0; i < t1.length; i++) {
            if (i == s1.length) {
                return t1[i];
            }
            if (s1[i] != t1[i]) {
                return t1[i];
            }
        }
        return ' ';
    }

    public static void main(String[] args) {
        FindDiffernceStrs fds = new FindDiffernceStrs();
        System.out.println(fds.findTheDifference("abcde", "abcdef"));
    }
}
