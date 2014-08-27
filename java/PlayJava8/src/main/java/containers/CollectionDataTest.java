package containers;

import think.util.CollectionData;
import think.util.Generator;

import java.awt.*;
import java.util.LinkedHashMap;
import java.util.LinkedHashSet;
import java.util.Set;

/**
 * Created by Qiao on 14-8-24.
 */

class Government implements Generator<String> {
    String[] foundation = ("strange women lying in ponds " +
            "distributing swords is no basis for a system of " +
            "government") .split(" ");
    private int index;

    @Override
    public String next() {
        return foundation[index++];
    }
}
public class CollectionDataTest {
    public static void main(String[] args) {
        Set<String> set = new LinkedHashSet<String>(
                new CollectionData<String>(new Government(),15)
        );
//        set.addAll(CollectionData.list(new Government(),15));
        System.out.println(set);
    }
}
