public class QuickSort {
    public void quickSort(int[] arr, int low, int high) {
        if (low >= high) {
            return;
        }
        int pivotIdx = partition(arr, low, high);
        quickSort(arr, low, pivotIdx - 1);
        quickSort(arr, pivotIdx + 1, high);
    }

    private int partition(int[] arr, int low, int high) {
        int pivot = arr[high];
        int idx = low - 1;
        for (int i = low; i < high; i++) {
            if (arr[i] < pivot) {
                idx++;
                swap(arr, idx, i);
            }
        }

        idx++;
        arr[high] = arr[idx];
        arr[idx] = pivot;
        return idx;
    }

    private void swap(int[] arr, int i, int j) {
        if (i == j) {
            return;
        }
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }

    public static void main(String[] args) {

        System.out.println("Hello World!");
        QuickSort qs = new QuickSort();
        int[] arr = { 1, 2, 3, 4, 5, 6, 7, 0 };
        qs.quickSort(arr, 0, arr.length - 1);
        for (int i = 0; i < arr.length; i++) {
            System.out.println(arr[i]);
        }

    }
}
