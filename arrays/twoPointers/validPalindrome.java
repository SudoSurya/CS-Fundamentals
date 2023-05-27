package arrays.twoPointers;

public class validPalindrome {
    static boolean isChar(char c) {
        if ((int) c > 96 && (int) c < 123 || (int) c < 58 && (int) c > 47 ) {
            return true;
        }
        return false;
    }
    public static boolean isPalindrome(String s) {
        // A man, a plan, a canal: Panama
        if(s.length() < 1){
            return true;
        }
        char[] arr = s.toLowerCase().toCharArray();
        int n = arr.length;
        int right = 0;
        int left = n - 1;
        String s1 = "";
        String s2 = "";
        while (right < n) {
            if (isChar(arr[right])) {
                s1 = s1 + arr[right];
            }
            if (isChar(arr[left])) {
                s2 = s2 + arr[left];
            }
            right++;
            left--;
        }
        System.out.println(s1);
        System.out.println(s2);
        return s1.equals(s2);
    }

    public static void main(String[] args) {
        boolean res = isPalindrome("race a car");
        System.out.println(res);
    }
}
