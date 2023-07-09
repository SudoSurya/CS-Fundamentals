package arrays.twoPointers;

public class ProductSign {
    public int arraySign(int[] nums) {

        int negSign = 0;
        int posSign = 0;
        for (int i = 0; i < nums.length; ++i) {
            if (nums[i] == 0) {
                return 0;
            }
            if (nums[i] < 0) {
                negSign++;
            } else {
                posSign++;
            }

        }
        if (negSign % 2 == 0) {
            return 1;
        }
        return -1;

    }

    public static void main(String[] args) {
        int nums[] = { 41, 65, 14, 80, 20, 10, 55, 58, 24, 56, 28, 86, 96, 10, 3, 84, 4, 41, 13, 32, 42, 43, 83, 78, 82,
                70, 15, -41 };
        ProductSign ps = new ProductSign();
        int result = ps.arraySign(nums);
        System.out.println(result);
    }

}
