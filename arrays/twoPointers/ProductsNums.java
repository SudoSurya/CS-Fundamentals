package arrays.twoPointers;

public class ProductsNums {
    public static int[] productExceptSelf(int[] nums) {

        int n = nums.length;
        int right = 0;
        int left = 0;
        int[] ans = new int[n];
        while (right < n) {
            ans[right] = 1;
            while (left < n) {
                if (left != right) {
                    ans[right] *= nums[left];
                }
                left++;
            }
            right++;
            left = 0;
        }
        return ans;

    }

    public static void main(String[] args) {
        int nums[] = { 1, 2, 3, 4 };
        int result[] = 
        productExceptSelf(nums);
        for (int i = 0; i < result.length; i++) {
            System.out.println(result[i]);
        }
    }
}
