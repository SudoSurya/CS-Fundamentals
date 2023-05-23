package arrays;

public class subarraySWT {
    static void SubArrays(int args[]) {
        int L = 0, R = 0;

        while (R < args.length) {
            while (L < args.length) {
                for (int i = R; i <= L; i++) {
                    System.out.print(args[i]);
                }
                System.out.println();
                L++;
            }
            R++;
            L = R;
        }
    }

    public static void main(String[] args) {
        int[] arr = { 1, 2, 3, 4 };
        SubArrays(arr);
    }
}
