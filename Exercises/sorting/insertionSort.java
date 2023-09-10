
public class insertionSort {

    static int insertion(int[] arr){

        int steps = 0;
        for(int i = 1; i < arr.length; i++){
            int key = arr[i];
            int j = i - 1;
            steps++;
            while(j >=0 && arr[j] > key){
                steps++;
                arr[j+1] = arr[j];
                j--;
            }
            arr[j+1] = key;
        }
        return steps;

    }

    public static void main(String[] args) {

        int[] arr = {5, 4, 3, 2, 1,7,6};
        int steps = insertion(arr);
        for(int i = 0; i < arr.length; i++){
            System.out.print(arr[i] + " ");
        }
        System.out.println("\nSteps: " + steps);

    }

}
