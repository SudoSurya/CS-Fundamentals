package arrays.twoPointers;

import java.util.HashMap;

public class ContainsDuplicate {
    public static boolean containsDuplicate(int[] nums) {
        HashMap<Integer, Integer> res = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            if (res.containsKey(nums[i])) {
                return true;
            }
            res.put(nums[i], 1);
        }
        return false;
    }

}
