package typeinfo;

import java.util.Random;

/**
 * Created by Qiao on 14-8-23.
 */

class Initable {
    static final int staticFanal = 47;
    static final int staticFanal2 =
            ClassInintialization.rand.nextInt(1000);
    static {
        System.out.println("Initializeing Initable");
    }
}

class Initable2 {
    static int staticNonFanal = 147;
    static {
        System.out.println("Initializeing Initable2");
    }
}

class Initable3 {
    static int staticNonFanal = 74;
    static {
        System.out.println("Initializeing Initable3");
    }
}

public class ClassInintialization {
    public static Random rand = new Random(47);

    public static void main(String[] args) {
        Class initable = Initable.class;
        System.out.println("After creating Initable ref");
        System.out.println(Initable.staticFanal);
        System.out.println(Initable.staticFanal2);
        System.out.println(Initable2.staticNonFanal);
        try {
            Class initable3 = Class.forName("typeinfo.Initable3") ;
            System.out.println("After create ing initalbe3 ref");
        } catch (ClassNotFoundException e) {
            System.out.println("can not find Initable3");
        }

        System.out.println(Initable3.staticNonFanal);
    }
}

