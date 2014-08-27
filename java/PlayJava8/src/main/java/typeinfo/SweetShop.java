package typeinfo;

/**
 * Created by Qiao on 14-8-23.
 */

class Candy{
    static {
        System.out.println("Loading Candy");}
}

class Gum {
    static {
        System.out.println("loading Gum");
    }
}

class Cookie {
    static {
        System.out.println("Loading Cookie");
    }
}

public class SweetShop {
    public static void main(String[] args) {
        System.out.println("inside main");
        new Candy();
        System.out.println("After creating Candy");
        try {
            Class.forName("typeinfo.Gum");
        } catch (ClassNotFoundException e) {
            System.out.println("Couldn't find Gum");
        }

        System.out.println("After Class.forNam(\"Gum\")");
        new Cookie();
        System.out.println("After creating Cookie");
    }
}
