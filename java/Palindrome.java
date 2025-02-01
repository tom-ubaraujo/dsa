public class Palindrome {
    public static void main(String[] args) {
        String str = "aaabbcbbaaa";
        StringBuilder reversed = new StringBuilder(str).reverse();
        System.out.println(reversed.toString().equals(str));

        String str2 = "Ugleiston";
        reversed = new StringBuilder(str2).reverse();
        System.out.println(reversed.toString().equals(str2));

        // outra forma de fazer com 'two pointers', mais performática
        int left = 0; int right = str.length()-1;
        boolean ans = true;
        while (left < right){
            if (str.charAt(left) != str.charAt(right)){
                ans = false; break;
            }
            left++; right--;
        }
        System.out.println("ans = " + ans);
    }
}
