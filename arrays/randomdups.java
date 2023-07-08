package arrays;

import java.util.HashMap;
import java.util.Map;

public class randomdups {
    public static int removeDuplicates(int[] nums) {
        HashMap<Integer, Integer> hm = new HashMap<>();

        // 1,2,3,4,4
        // ^ ^
        for (int i = 0; i < nums.length; i++) {
            hm.put(nums[i], hm.getOrDefault(nums[i], 0) + 1);

        }
        // iterate over hashmap and return key with max value store max key value in
        // variable
        // iterate over hashmap and remove all keys except max key
        int max = 0;
        int maxkey = 0;
        for (Map.Entry<Integer, Integer> entry : hm.entrySet()) {
            if (entry.getValue() > max) {
                max = entry.getValue();
                maxkey = entry.getKey();
            }
        }
        return maxkey;
    }

    public static void main(String[] args) {
        int nums[] = { 0, 0, 1, 1, 1, 2, 2, 3, 3, 4 ,4,4,4,4,4,4};
        System.out.println(removeDuplicates(nums));
    }

}
