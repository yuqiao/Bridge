(ns hello.core
  (:gen-class))
(def vowel? (set "aeiou"))

(defn pig-latin [word]
	(let [first-letter (first word)]
		(if (vowel? first-letter) 
			(str word "ay")
			(str (subs word 1) first-letter "ay"))))
(defn parting
  "returns a String parting in a giving language"
  ([](parting "World"))
  ([name] (parting name "en"))
  ([name language]
    (condp = language
      "en" (str "Goodbye, " name)
      "es" (str "Adios, " name)
      (throw (IllegalArgumentException. 
        (str "unsupported language " language))) ) ))


(defmulti what-am-i class)
(defmethod what-am-i Number [arg] (println arg "is a Number"))
(defmethod what-am-i String [arg] (println arg "is a String"))
(defmethod what-am-i :default [arg] (println arg "is a something else."))

(defn callback1 [n1 n2 n3] (+ n1 n2 n3))
(defn callback2 [n1 _ n3] (+ n1 n3))
(defn caller [callback value]
  (callback (+ value 1) (+ value 2) (+ value 3)))

(defn- polynomial
  "computes the value of a polynomial"
  [coefs x] )

(defn -main
  "I don't do a whole lot ... yet."
  [& args]
  (println "Hello, World!")
  (println (pig-latin "red"))
  (println (pig-latin "orange"))
  (println (parting))
  (println (parting "Mark"))
  (println (parting "Mark" "es"))
  (what-am-i 19)
  (what-am-i "Hello")
  (what-am-i true)
  (println (caller callback1 10) )
  (println (caller callback2 10) )
  )