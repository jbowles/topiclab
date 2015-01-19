package main

import (
	"fmt"
	topic "github.com/jbowles/topiclab"
)

func PrintClf(c *topic.Classifier) {
	fmt.Printf("%+v\n", c)
	for k, v := range c.Datas {
		fmt.Printf(" --- %v --- \n", k)
		for word, counts := range v.Freqs {
			fmt.Printf("%v: %v\n", word, counts)
		}
		fmt.Printf("\n")
	}
}

func Train(c *topic.Classifier, text []string, which topic.Class) {
	data := c.Datas[which]
	for _, word := range text {
		data.Freqs[word]++
		data.Total++
	}
	c.Learned++
}

func main() {
	c := topic.NewClassifier("Good", "Bad", "Neutral")
	//c.Learn([]string{"that", "when", "who", "where"}, "THIS")
	c.Learn([]string{"wonderful", "fantastic", "awesome", "good"}, "Good")
	c.Learn([]string{"horrible", "hate", "worst", "bad"}, "Bad")
	c.Learn([]string{"alright", "decent", "ok", "even"}, "Neutral")
	PrintClf(c)
}
