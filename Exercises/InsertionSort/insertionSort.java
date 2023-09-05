class insertionSort {
    static int[] insertion(int[] haystack) {
        int n = haystack.length;

        for (int j = 1; j < n; j++) {
            int key = haystack[j];
            int i = j - 1;
            // if haystack[i] < key , then arrays sorted in descending order
            // if haystack[i] > key , then arrays sorted in ascending order
            while (i >= 0 && haystack[i] < key) {
                haystack[i + 1] = haystack[i];
                i--;
            }
            haystack[i + 1] = key;
        }
        return haystack;
    }

    public static void main(String[] args) {
        int[] haystack = { 5, 2, 4, 6, 1, 3, 71 };
        int[] sorted = insertion(haystack);
        for (int i = 0; i < sorted.length; i++) {
            System.out.println(sorted[i]);
        }
    }
}
