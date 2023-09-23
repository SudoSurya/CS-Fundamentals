public class QuickSort {
    public int[] quickSort(int[] arr , int low, int high) {
        return arr;
    }
    private int partition(int[] arr, int low, int high) {
        return 0;
    }
    public static void main(String[] args) {

        qs = new QuickSort();
        int[] arr = { 1, 2, 3, 4, 5, 6, 7 };
        int[] sortedArr = qs.quickSort(arr, 0, arr.length - 1);
        for (int i = 0; i < sortedArr.length; i++) {
            System.out.println(sortedArr[i]);
        }

    }
}
