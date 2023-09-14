
public class reverseStr {
    static char[] reverse(String needle) {
        char[] arr = needle.toCharArray();

        char reverse[] = new char[arr.length];

        for (int i = 0; i < arr.length; i++) {
            reverse[i] = arr[arr.length - i - 1];
        }
        return reverse;
    }

    public static void main(String[] args) {
        String needle = "Hello World";
        char[] arr = reverse(needle);
        System.out.println(arr);
    }
}
