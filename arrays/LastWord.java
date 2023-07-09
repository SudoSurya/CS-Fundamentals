package arrays;

public class LastWord {
    public int lengthOfLastWord(String s) {
        int n = s.length();
        int i = n - 1;
        int count = 0;

        while (i >= 0 && s.charAt(i) == ' ') {
            i--;
        }

        while (i >= 0 && s.charAt(i) != ' ') {
            count++;
            i--;
        }
        return count;

        // String strs[] = s.split(" ");
        /// iterrate strs array from last to first

        // for (int j = strs.length - 1; j >= 0; j--) {
        // if (strs[j].length() > 0) {
        // return strs[j].length();
        // }
        // }
    }

    public static void main(String[] args) {
        String s = "Hello World";
        LastWord lw = new LastWord();
        int result = lw.lengthOfLastWord(s);
        System.out.println(result);
    }
}
