package arrays.twoPointers;

public class containerwithmostwater {

    static int waterWall(int[] nums) {
        int n = nums.length;
        int R = 0;
        int L = n - 1;
        int max = Integer.MIN_VALUE;
        while (R < L) {
            int width = L - R;
            max = Math.max(max, (Math.min(nums[R], nums[L]) * width));
            if (nums[R] <= nums[L]) {
                R++;
            } else {
                L--;
            }
        }
        return max;
    }

    static int solution(int[] nums) {
        // [1,8,6,2,5,4,8,3,7]
        // 1 1

        int n = nums.length;
        if (n < 3) {
            return 1;
        }
        int R = 0;
        int L = n - 1;
        int max = Integer.MIN_VALUE;
        int count = 1;
        while (R < n) {
            while (R != L) {
                int min;
                min = Math.min(nums[R], nums[L]);
                count *= min * (L - 1);
                max = Math.max(max, count);
                L--;
                count = 1;
            }
            R++;
            L = n - 1;
        }
        return max;
    }

    public static void main(String[] args) {
        // int[] nums = { 1, 8, 6, 2, 5, 4, 8, 3, 7 };
        int[] nums = { 3, 3 };
        int res = waterWall(nums);
        System.out.println(res);

    }
}
