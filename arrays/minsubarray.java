package arrays;

public class minsubarray {
    public static int minSubArrayLen(int target, int args[]) {

        int L = 0, R = 0;
        int min = Integer.MAX_VALUE, sum = 0;
        int n = args.length;

        while (R < n) {
            sum += args[R];
            while (target <= sum) {
                min = Math.min(min, R - L + 1);
                sum -= args[L];
                L++;
            }
            R++;
        }

        return min == Integer.MAX_VALUE ? 0 : min;
    }

    public static void main(String[] args) {
        int[] ARGS = { 11, 33, 44, 4, 4 };
        int target = 510;
        System.out.println(minSubArrayLen(target, ARGS));

    }

}
