import java.nio.charset.StandardCharsets;

public class ReverseString {
    public static void main(String[] args) {
        String str = "Hello World!!!", reversed = "";

        for (int i = str.length()-1; i >= 0; i--) {
            reversed += str.charAt(i);
        }

        // revertendo com StringBuilder
        // desta forma lidamos com um objeto da classe StringBuilder, que permite manipulações
        // mas precisa ser convertido para String depois.
        StringBuilder reversed2 = new StringBuilder(str).reverse();

        System.out.println("reversed = " + reversed);
        System.out.println("reversed2 = " + reversed2.toString());

        // revertendo um número
        int str3 = 12345;
        StringBuilder reversed3 = new StringBuilder(String.valueOf(str3)).reverse();
        System.out.println("Num Reversed: " + Integer.valueOf(reversed3.toString()));

    }
}
