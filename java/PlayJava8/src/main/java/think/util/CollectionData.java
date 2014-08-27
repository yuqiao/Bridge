package think.util;

import java.util.ArrayList;

/**
 * Created by Qiao on 14-8-24.
 */
public class CollectionData<T> extends ArrayList<T> {
    public CollectionData(Generator<T> gen, int quantity){
        for(int i=0;i<quantity; i++){
            add(gen.next());
        }
    }

    public static <T> CollectionData<T> list(Generator<T> gen, int qualtity){
        return new CollectionData<T>(gen, qualtity);
    }
}
