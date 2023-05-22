package arrays;

// import java.util.Scanner;

class arrays {

    public static void passByRef(String args[]) {
        args[0] = "changed";
    }

    public static void main(String[] args) {
        // Scanner sc = new Scanner(System.in);
        int arr[] = new int[10];
        arr[1] = 10;
        System.out.println(arr[0]);

        String[] Fruits = { "apple", "Banana", "Custord" };
        System.out.println("Before Function call " + Fruits[0]);
        arrays.passByRef(Fruits);
        System.out.println("After Function call " + Fruits[0]);
        // mutlidimensional arrays
        int[][] twoD = new int[4][4];
        for (int i = 0; i < twoD.length; ++i) {
            for (int j = 0; j < twoD[i].length; ++j) {
                twoD[i][j] = i * j;
                System.out.print(twoD[i][j]);
            }
            System.out.println();
        }
    }
}