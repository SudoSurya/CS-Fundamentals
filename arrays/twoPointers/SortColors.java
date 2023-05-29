package arrays.twoPointers;

public class SortColors {

    private static void swap(int nums[], int p1, int p2) {
        int temp = nums[p1];
        nums[p1] = nums[p2];
        nums[p2] = temp;
    }

    public static void solution(int nums[]) {
        int n = nums.length;
        int left = 0, right = n - 1, curr = 0;
        // 0 0 1 2 2 // 1 < 2 // 
        while (curr <= right) {
            if (nums[curr] == 2) {
                swap(nums, curr, right);
                right--;
            } else if (nums[curr] == 1) {
                curr++;
            } else {
                swap(nums, curr, left);
                left++;
                curr++;
            }
        }
    }

}
