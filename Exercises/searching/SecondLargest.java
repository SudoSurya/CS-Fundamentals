public class SecondLargest {
    public static int SecondLargeEle(int[] array) {
        int largest = Integer.MIN_VALUE;
        int Secondlargest = Integer.MIN_VALUE;

        for (int i = 0; i < array.length; i++) {
            if (array[i] > largest) {
                Secondlargest = largest;
                largest = array[i];
            } else if (array[i] > Secondlargest && array[i] != largest) {
                Secondlargest = array[i];
            }
        }

        return Secondlargest;
        // 1 > largest , secondlargest = MIN_VALUE, largest = 1
        // 1 > secondlargest true, 1 != 1 false
        // 2 > largest true, secondlargest = 1, largest = 2
        // 2 > secondlargest true, 2 != 2 false
        // 3 > largest true, secondlargest = 2, largest = 3
        // 3 > secondlargest true, 3 != 3 false
        // 4 > largest true, secondlargest = 3, largest = 4
        // 4 > secondlargest true, 4 != 4 false

    }

    public static void main(String[] args) {
        int[] array = {9,8};
        System.out.println(SecondLargeEle(array));
    }
}
