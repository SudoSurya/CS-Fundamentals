public class GreatestProduct {
    public int partition(int arr[],left,right){
        int pivot = arr[right];
        int i = left - 1;
        for(int j = left; j < right; j++){
            if(arr[j] <= pivot){
                i++;
                int temp = arr[i];
                arr[i] = arr[j];
                arr[j] = temp;
            }
        }
        int temp = arr[i+1];
        arr[i+1] = arr[right];
        arr[right] = temp;
        return i+1;
    }

    public int Product(int arr[], int left, int right) {
        int sum = 1;
        if (arr.length < 3) {
            for (int i = 0; i < arr.length; i++) {
                sum *= arr[i];
            }
            return sum;
        }

        int pivot = partition(arr, left, right);

        if()
    }

    public static void main(String[] args) {

    }
}
