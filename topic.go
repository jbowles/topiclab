package topiclab

// defaultProb is the tiny smoothing non-zero probability that a word
// we have not Seen before appears in the class.
const defaultProb = 0.00000000001

type Class string
type Distribution []float64
type Histogram []int
type FreqCount int
type Corpus []*Document

type Sampler struct {
	topic_prior float64
	word_prior  float64
	model       *Model
	accum_model *Model
}

type Model struct {
	topic_histograms map[string]Histogram
	global_histogram Histogram
	zero_histogram   Histogram
}

type Document struct {
	unique_words       []string
	wordtopics_indices []int
	wordtopics         []int
	topic_histogram    Histogram
}

// Classifier implements the Naive Bayesian Classifier.
type Classifier struct {
	Classes []Class
	Learned int // docs Learned
	Seen    int // docs Seen
	Datas   map[Class]*BayesData
	//datas   map[Class]*classData
}

// serializableClassifier represents a container for
// Classifier objects whose fields are modifiable by
// reflection and are therefore writeable by gob.
type serializableClassifier struct {
	Classes []Class
	Learned int
	Seen    int
	Datas   map[Class]*BayesData
}

// topicHistogram holds the frequency data for words in a
// particular class. In the future, we may replace this
// structure with a trie-like structure for more
// efficient storage.
//classData
type BayesData struct {
	Freqs map[string]FreqCount
	Total int
}

type Category struct {
	Label string
	Text  string
}

type Topic interface {
	//single category
	Train(class, text string) (err error)

	//multiple categories
	MultiTrain(cat []Category) (err error)

	//Drop removes the category label and its text
	Drop(label, text string) (err error)

	// socres for each normalized class
	Score(text string) (scores map[string]float64, err error)

	// Classify returns category label with most likely score
	Classify(text string) (label string, err error)

	// Reset the model
	Reset() error
}
