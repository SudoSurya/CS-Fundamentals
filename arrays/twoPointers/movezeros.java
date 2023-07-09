package arrays.twoPointers;

public class movezeros {

    static void solution(int[] nums) {
        int n = nums.length;
        int R = 1, L = 0;

        while (R < n) {

            if (nums[L] != 0) {
                L++;
                R++;
            } else if (nums[R] == 0) {
                R++;
            } else {
                int temp = nums[R];
                nums[R] = nums[L];
                nums[L] = temp;
            }
        }

    }

    static void printer(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            System.out.print(nums[i] + " ");
        }
        System.out.println();
    }

    public static void main(String[] args) {
        int nums[] = { 0,1,3,0,12 };
        printer(nums);
        solution(nums);
        printer(nums);
    }

}
