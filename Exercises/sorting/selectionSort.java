/**
 * selectionSort
 */
public class selectionSort {
    static void SelectionSort(int[] arr) {
        for (int i = 0; i < arr.length; i++) {
            int min_val_index = i;
            for (int j = i + 1; j < arr.length; j++) {
                if (arr[j] < arr[min_val_index]) {
                    min_val_index = j;
                }
            }
            if (min_val_index != i) {
                swap(arr, i, min_val_index);
            }
        }
    }

    static void swap(int[] arr, int i, int j) {
        int temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }

    public static void main(String[] args) {
        int arr[] = { 5, 4, 3, 2, 1 };
        SelectionSort(arr);
        for (int i : arr) {
            System.out.print(i + " ");
        }

    }
}
