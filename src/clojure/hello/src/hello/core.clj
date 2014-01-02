(ns hello.core
  (:gen-class))

(def vowel? (set "aeiou"))

(defn pig-latin [word]
	(let [first-letter (first word)]
		(if (vowel? first-letter)
			(str word "ay")
			(str (subs word 1) first-letter "ay"))))


(defn -main
  "I don't do a whole lot ... yet."
  [& args]
  (println "Hello, World!")
  (println (pig-latin "red"))
  (println (pig-latin "orange")))
