package strings;

public class mergestr {
    static String solution(String word1, String word2) {
        int word1len = word1.length();
        int word2len = word2.length();
        // abc pqr
        // ^ ^
        int L = 0;
        int R = 0;
        StringBuilder sb = new StringBuilder();

        while (L + R < word1len + word2len) {
            if (L < word1len) {
                sb.append(word1.charAt(L));
                L++;
            }
            if (R < word2len) {
                sb.append(word2.charAt(R));
                R++;
            }
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        System.out.println(solution("abc", "pqr"));
    }

}
