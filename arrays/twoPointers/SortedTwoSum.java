package arrays.twoPointers;

public class SortedTwoSum {
    public static int[] twoSum(int[] arr, int target) {
        int right = 0, n = arr.length, left = 1;
        if (n < 2) {
            return new int[] {};
        }
        // [5,25,75]
        // 3,24,50,79,88,150,345

        while (left < n) {

            if (arr[right] + arr[left] == target) {
                return new int[] { right + 1, left + 1 };
            } else if (arr[right] + arr[left] < target && left < n) {
                left++;
            } else {
                right++;
                left = right + 1;
            }
        }
        return new int[] {};
    }

    public static void main(String[] args) {
        int[] arr = twoSum(new int[] { 3, 24, 50, 79, 88, 150, 345 }, 200);
        for (int i = 0; i < arr.length; i++) {
            System.out.println(arr[i]);
        }

    }
}
