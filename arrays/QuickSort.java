package arrays;

public class QuickSort {
    public static void quickSort(int[] arr) {
        qs(arr, 0, arr.length - 1);
    }

    public static void qs(int[] arr, int li, int hi) {
        if (li >= hi) {
            return;
        }

        int pivotIdx = partition(arr, li, hi);
        qs(arr, li, pivotIdx - 1);
        qs(arr, pivotIdx + 1, hi);
    }

    public static int partition(int[] arr, int li, int hi) {
        int pivot = arr[hi];

        int idx = li - 1;

        for (int i = li; i < hi; i++) {
            if (arr[i] < pivot) {
                idx++;
                int temp = arr[i];
                arr[i] = arr[idx];
                arr[idx] = temp;
            }
        }

        idx++;
        arr[hi] = arr[idx];
        arr[idx] = pivot;

        return idx;
    }

    public static void main(String[] args) {
        int[] arr = { 9, 3, 7, 5, 1, 6, 8, 2, 4, 10 };
        quickSort(arr);
        for (int num : arr) {
            System.out.print(num + " ");
        }
    }
}
