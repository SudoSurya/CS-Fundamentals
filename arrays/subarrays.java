package arrays;

public class subarrays {
    static void subArrays(char args[]) {
        for (int i = 0; i < args.length; i++) {
            for (int j = i; j < args.length; j++) {
                for (int k = i; k <= j; k++) {
                    System.out.print(args[k]);
                }
                System.out.println();
            }
        }
    }

    public static void main(String[] args) {
        char[] arr = { 'a', 'b', 'c', 'd', 'e' };
        subArrays(arr);
    }
}
