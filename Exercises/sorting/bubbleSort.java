public class bubbleSort {
    static int[] BubbleSort(int[] arr) {
        boolean sorted = false;
        int unsorted = arr.length - 1;

        while (!sorted) {
            sorted = true;
            for (int i = 0; i < unsorted; i++) {
                if (arr[i] > arr[i + 1]) {
                    swap(arr, i, i + 1);
                    sorted = false;
                }
            }
            unsorted--;
        }
        return arr;

    }

    private static void swap(int[] arr, int i, int i1) {
        int temp = arr[i];
        arr[i] = arr[i1];
        arr[i1] = temp;
    }

    public static void main(String[] args) {
        int[] arr = { 5, 4, 3, 2, 1 };
        int[] sorted = BubbleSort(arr);
        for (int i = 0; i < sorted.length; i++) {
            System.out.println(sorted[i]);
        }

    }
}
