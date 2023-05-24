package arrays;

public class maxone2 {

    public static int solution(int arr[]) {
        int n = arr.length;
        if (n < 2)
            return n;

        int L = 0, R = 0;
        int counter = 0;
        int maxLen = 0;

        while (R < n) {
            if (arr[R] == 0) {
                counter++;
            }
            while (counter > 1) {
                if (arr[L] == 0) {
                    counter--;
                }
                L++;
            }
            maxLen = Math.max(maxLen, R - L + 1);
            R++;
        }
        return maxLen;
    }

    public static void main(String[] args) {
        int arr[] = { 1, 0, 1, 0, 0, 1, 1, 1, 1};
        int result = solution(arr);
        System.out.println(result);
    }

}
