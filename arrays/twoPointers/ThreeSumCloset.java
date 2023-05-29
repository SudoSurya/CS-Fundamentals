package arrays.twoPointers;

public class ThreeSumCloset {
    static int Solution(int nums[], int target) {
        int gap = Integer.MAX_VALUE, ans = 0, n = nums.length;

        for (int i = 0; i < nums.length; i++) {
            int left = i, right = n - 1;
            while (left < right) {
                int curSum = nums[i] + nums[left] + nums[right];
                if (curSum == target) {
                    return target;
                } else if (curSum < target) {
                    left++;
                } else {
                    right--;
                }
                int curgap = Math.abs(curSum - target);
                if (curgap < gap) {
                    gap = curgap;
                    ans = curSum;
                }
            }
        }
        return ans;
    }

    public static void main(String[] args) {
        int nums[] = { 1, 2, 3, 4 };
        int result = Solution(nums, 3);
        System.out.println(result);
    }
}
