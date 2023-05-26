package arrays.twoPointers;

public class RemoveDups {

    static int solution(int[] nums) {
        int n = nums.length;
        if (n < 2)
            return n;

        int R = 1, L = 0;
        // 1 2 3 4 3 4
        //              ^ 
        //          ^
        while (R < n) {
            if (nums[L] != nums[R]) {
                L++;
                nums[L] = nums[R];
            }
            R++;
        }
        return L + 1;

    }

    public static void main(String[] args) {
        int nums[] = {1,1,2,3,4,5,6};
        int result = solution(nums);
        System.out.println(result);
    }

}
