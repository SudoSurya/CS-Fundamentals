package arrays.twoPointers;

public class twosum {

    static int[] solution(int[] arr, int target) {

        int right = 0, n = arr.length, left = n - 1;

        while (right < n) {

            while (left != right) {
                if (arr[right] + arr[left] == target) {
                    return new int[] { right, left };
                }
                left--;
            }
            right++;
            left = n - 1;
        }
        return new int[] {};
    }

    public static void main(String[] args) {
        int[] arr = { 3, 2, 4 };
        int[] res = solution(arr, 6);
        for (int i : res) {
            System.out.println(i);
        }
    }
}
