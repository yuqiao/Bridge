package think.util;

import java.util.Iterator;
import java.util.LinkedHashMap;

/**
 * Created by Qiao on 14-8-24.
 */
public class MapData<K,V> extends LinkedHashMap<K,V> {
    public MapData(Generator<Pair<K,V>> gen, int quantity){
        for(int i=0;i<quantity;i++){
            Pair<K,V> p = gen.next();
            put(p.key, p.val);
        }
    }

    public MapData(Generator<K> kgen, Generator<V> vgen, int quantity){
        for(int i=0;i<quantity;i++){
            put(kgen.next(), vgen.next());
        }
    }

    public MapData(Generator<K> kgen, V val, int quantity){
        for(int i=0;i<quantity;i++){
            put(kgen.next(), val);
        }
    }

    public MapData(Iterable<K> kgen, Generator<V> vgen){
        for(K key:kgen){
           put(key, vgen.next());
        }
    }

    public MapData(Iterable<K> kgen, V val){
        for(K key:kgen){
            put(key, val);
        }
    }



}
