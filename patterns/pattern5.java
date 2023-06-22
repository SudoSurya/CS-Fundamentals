package patterns;

public class pattern5 {
    static void pattern(int n) {
        int temp = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j <= i; j++) {
                temp++;
                System.out.print(temp + " ");
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        pattern(5);
    }
}