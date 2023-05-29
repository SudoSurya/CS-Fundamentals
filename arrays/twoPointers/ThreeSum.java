package arrays.twoPointers;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

public class ThreeSum {
    static List<List<Integer>> ThreeSumSolution(int nums[]) {
        List<List<Integer>> res = new LinkedList<>();
        int n = nums.length;
        if (n < 3)
            return res;

        Arrays.sort(nums);

        for (int i = 0; i < nums.length; i++) {
            if (i != 0 && nums[i] == nums[i - 1])
                continue;
            int L = i + 1, R = n - 1;
            while (L < R) {
                int curSum = nums[i] + nums[L] + nums[R];
                if (curSum == 0) {
                    List<Integer> Sub = new LinkedList<>();
                    Sub.add(nums[i]);
                    Sub.add(nums[L]);
                    Sub.add(nums[R]);
                    res.add(Sub);
                    L++;
                    R--;
                    while (L < R && nums[L] == nums[L - 1]) {
                        L++;
                    }
                    while (L < R && nums[R] == nums[R - 1]) {
                        R--;
                    }
                } else if (curSum < 0) {
                    L++;
                } else {
                    R--;
                }
            }
        }
        return res;
    }

    public static void main(String[] args) {
        int nums[]  = {-1,0,1,2,-1,-4};
        List<List<Integer>> res = ThreeSumSolution(nums);
        System.out.println(res);
    }
}
