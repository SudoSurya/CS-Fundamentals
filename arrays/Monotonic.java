package arrays;

public class Monotonic {
    public static boolean isMonotonic(int[] nums) {

        int n = nums.length;
        boolean inc = false;
        boolean dec = false;

        for (int i = 0; i < n - 1; i++) {
            if (nums[i] < nums[i + 1]) {
                inc = true;
            }
            if (nums[i] > nums[i + 1]) {
                dec = true;
            }
        }
        return inc && dec;
    }

    public static void main(String[] args) {
        int nums[] = { 4, 1, 2, 2, 3 };
        System.out.println(isMonotonic(nums));
    }
}
