package arrays;

public class maxOnes {
    public static void main(String[] args) {
        int arr[] = { 1, 1, 1, 1, 0, 0, 1, 1, 1,1,1,1,0,1 };
        int counter = 0;
        int max = 0;

        for (int i = 0; i < arr.length; i++) {

            if (arr[i] == 1) {
                counter += 1;
            } else if (arr[i] == 0) {
                max = Math.max(max, counter);
                counter = 0;
            }
            max = Math.max(max, counter);
        }
        System.out.println(max);

    }

}
