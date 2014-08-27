package typeinfo.toys;

/**
 * Created by Qiao on 14-8-23.
 */

interface HasBatteries {}
interface Waterproof {}
interface Shoots {}

class Toy {
    Toy() {}
    Toy(int i) {}
}

class FancyToy extends  Toy implements  HasBatteries, Waterproof, Shoots{
    FancyToy() {super(1);}
}

public class ToyTest {
    static void printInfo(Class cc){
        System.out.println("Class name: " + cc.getName() +
            "is insterface?[" + cc.isInterface()+"]" );
        System.out.println("Simple name:" + cc.getSimpleName());
        System.out.println("Canonical name:" + cc.getCanonicalName());
        System.out.println();
    }

    public static void main(String[] args) {
        Class c = null;
        try {
            c = Class.forName("typeinfo.toys.FancyToy");
        } catch (ClassNotFoundException e) {
            System.out.println("Can not find FancyToy");
            System.exit(1);
        }

        printInfo(c);
        for(Class face:c.getInterfaces())
            printInfo(face);
        Class up = c.getSuperclass();
        Object obj = null;
        try {
            obj = up.newInstance();
        } catch (InstantiationException e) {
            System.out.println("can not instantiate");
            System.exit(1);
        } catch (IllegalAccessException e) {
            System.out.println("can not access");
            System.exit(1);
        }

        printInfo(obj.getClass());
    }
}
