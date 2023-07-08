package strings;

public class OccureneStr {
    public int strStr(String haystack, String needle) {
        // int L = 0;
        // int R = 0;
        // int needlelen = needle.length();
        // int haystacklen = haystack.length();
        // while (L + R < haystacklen) {
        // if (haystack.charAt(L) == needle.charAt(R)) {
        // L++;
        // R++;
        // } else {
        // L = L - R + 1;
        // R = 0;
        // }
        // if (R == needlelen) {
        // return L - R;
        // }
        // }
        // return -1;
        if (needle.length() == 0)
            return 0;

        if (needle.length() > haystack.length())
            return -1;
        // abc c
        char[] h = haystack.toCharArray();
        String res = "";
        for (int i = 0; i < needle.length(); ++i) {
            res += h[i];
        }
        if (res.equals(needle))
            return 0;
        for (int j = needle.length(); j < h.length; ++j) {
            System.out.println(h[j]);
            res = res.substring(1) + h[j];
            if (res.equals(needle))
                return j - needle.length() + 1;
        }
        return -1;
    }

    public static void main(String[] args) {
        OccureneStr os = new OccureneStr();
        System.out.println(os.strStr("sadbutsad", "ad"));
    }
}

// Compare this snippet from strings/ReverseStr.java:
