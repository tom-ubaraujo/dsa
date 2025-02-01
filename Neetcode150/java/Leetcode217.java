import java.util.HashMap;
import java.util.HashSet;

public class Leetcode217 {
    public static void main(String[] args) {
        int[] arr1 = {1, 5, 7, 10, 4, 1}; //true
        int[] arr2 = {2, 5, 7, 10, 4, 1}; //false
        int[] arr3 = {3, 7, 7, 10, 0, 9}; //true

        System.out.println(hasDuplicateOn(arr1));
        System.out.println(hasDuplicateOn(arr2));
        System.out.println(hasDuplicateOn(arr3));

        System.out.println(hasDuplicateOn2(arr1));
        System.out.println(hasDuplicateOn2(arr2));
        System.out.println(hasDuplicateOn2(arr3));

    }

    public static boolean hasDuplicateOn(int[] arr) {
        HashSet<Integer> visited = new HashSet<>();

        for (int num : arr) {
            if (visited.contains(num)) {
                return true;
            }
            visited.add(num);
        }
        return false;
    }

    public static boolean hasDuplicateOn2(int[] arr) {
        for (int i = 0; i < arr.length; i++) {
            for (int j = i + 1; j < arr.length; j++) {
                if (arr[i] == arr[j]) {
                    return true;
                }
            }
        }
        return false;
    }
}

