package arrays;

public class subarrays {
    static void subArrays(char args[]) {
        for (int i = 0; i < args.length; i++) {
            for (int j = i; j < args.length; j++) {
                for (int j2 = i; j2 <= j; j2++) {
                    System.out.print(args[j2]);
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
