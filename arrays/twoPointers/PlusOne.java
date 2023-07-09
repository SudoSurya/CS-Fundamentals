package arrays.twoPointers;

public class PlusOne {
    public int[] plusOne(int[] digits) {
        int n = digits.length;
        int carry = 1;
        int sum = 0;
        // 1996
        for (int i = n - 1; i >= 0; --i) {
            sum = digits[i] + carry; // 7 // 9 // 9 // 2
            digits[i] = sum % 10; // 7 % 10 = 7 // 9 % 10 = 9 // 9 % 10 = 9 // 2 % 10 = 2
            carry = sum / 10; // 7 / 10 = 0
        }
        if (carry == 0) {
            return digits;
        }
        int[] result = new int[n + 1];
        result[0] = carry;
        for (int i = 1; i < n + 1; ++i) {
            result[i] = digits[i - 1];
        }
        return result;
    }
    public static void main(String[] args) {
        
        int [] digits = {1,2,3};
        PlusOne po = new PlusOne();
        int [] result = po.plusOne(digits);
        for (int i = 0; i < result.length; ++i) {
            System.out.print(result[i] + " ");
        }
    }
}
