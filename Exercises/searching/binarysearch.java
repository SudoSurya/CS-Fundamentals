public class binarysearch {
    static int binarySearch(int[] haystack, int needle) {
        int low = 0;
        int high = haystack.length - 1;

        do {
            int mid = (low + high) / 2;

            if (haystack[mid] == needle) {
                return mid;
            } else if (haystack[mid] < needle) {
                low = mid + 1;
            } else {
                high = mid - 1;
            }

        } while (low <= high);

        return -1;
    }

    public static void main(String[] args) {
        int[] haystack = { 1, 2, 3, 4, 5, 6, 7, 8, 9 };
        int needle = 1;

        int result = binarySearch(haystack, needle);
        System.out.println(result);

    }
}
