(defn blank? [str]
	(every? #(Character/isWhitespace %) str))

(println (blank? ""))
(println (blank? "    "))
(println (blank? "a "))

(defrecord Person [first-name last-name])
(def foo ( ->Person "Aeron" "Bedra"))

(println foo)

(blank? "")


