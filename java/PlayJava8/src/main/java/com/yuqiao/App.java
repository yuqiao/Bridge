package com.yuqiao;

import java.math.BigDecimal;
import java.util.Arrays;
import java.util.List;
import java.util.function.Predicate;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
        System.out.println( "Hello World!" );
        final List<BigDecimal> prices = Arrays.asList(
                new BigDecimal("10"), new BigDecimal("30"), new BigDecimal("17"), new BigDecimal("20"),
                new BigDecimal("15"), new BigDecimal("18"), new BigDecimal("45"), new BigDecimal("12"));

        final BigDecimal totalDicountedPrices = prices.stream().filter(price -> price.compareTo(BigDecimal.valueOf(20))>0)
                .map(price -> price.multiply(BigDecimal.valueOf(0.9)))
                .reduce(BigDecimal.ZERO, BigDecimal::add);
        System.out.println("total:" + totalDicountedPrices);


        final List<String> friends =
                Arrays.asList("Brian", "Nate", "Neal", "Raju", "Sara", "Scott");
//        friends.forEach(fr -> System.out.println(fr));
        friends.forEach(System.out::println);
        friends.stream().map(String::toUpperCase).forEach(System.out::println);

        final Predicate<String> startWithN = name -> name.startsWith("N");

        final long countFriendsStartN = friends.stream()
                .filter(startWithN).count();
        System.out.println("countFriendsStartN:" + countFriendsStartN);
    }
}
