package arrays;

public class slidingWindow {
    static int MaxSubArray(int a[], int k) {
        int kSum = 0;
        int mSum = Integer.MIN_VALUE;
        for (int i = 0; i < k; i++) {
            kSum += a[i];
            // System.out.println(kSum);
        }
        for (int i = k; i < a.length; i++) {
            mSum = Math.max(mSum, kSum);
            kSum = kSum - a[i - k] + a[i];
            mSum = Math.max(mSum, kSum);
        }
        return mSum;
    }

    public static void main(String[] args) {
        int[] arr = { 900, 100, 100, 1001, 99972, 3, 4000 };
        int k = 3;
        int result = MaxSubArray(arr, k);
        System.out.println(result);
    }
}
