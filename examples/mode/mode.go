package main

import (
	"bufio"
	"bytes"
	"compress/flate"
	"encoding/json"
	"fmt"
	"math"
	. "misc/partridge/vec"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Comment struct {
	ID        string `json:"id"`
	Author    string `json:"author"`
	Score     int64  `json:"score"`
	Subreddit string `json:"subreddit"`
	Body      string `json:"body"`
}

func main() {
	var redditDataFile, errOpenRedditData = os.Open("/home/paul/Downloads/torrents/reddit_comments_100k.json")
	if errOpenRedditData != nil {
		panic(errOpenRedditData)
	}
	defer redditDataFile.Close()
	var comments = make([]Comment, 0, 10000)
	var lines = bufio.NewScanner(redditDataFile)
	for lines.Scan() {
		if lines.Text() == "" {
			continue
		}
		var comment = Comment{}
		if err := json.Unmarshal([]byte(lines.Text()), &comment); err != nil {
			fmt.Printf("skipping value %q\n"+
				"error: %v\n", lines.Text(), err)
			continue
		}
		comments = append(comments, comment)
	}
	var reddit = MakeDataFrame(len(comments), func(index int) Frame {
		var comment = comments[index]
		var buf = &bytes.Buffer{}
		var compressor, _ = flate.NewWriter(buf, flate.BestSpeed)
		compressor.Write([]byte(comment.Body))
		compressor.Flush()
		var count = float64(buf.Len()) / float64(len(comment.Body))
		return FrameFromValues(comment.Author, comment.ID, float64(comment.Score), count)
	}, "author", "id", "score", "fuck_score").Filter(func(frame Frame) bool {
		return frame.Value(2).AsFloat64() != 1.0
	}).GroupByFloat(
		func(frame Frame) float64 {
			var fuck_score = frame.Value(3).AsFloat64()
			return fuck_score
		},
		func(dataframe DataFrame) DataFrame {
			return MakeDataFrame(1, func(int) Frame {
				var mean = dataframe.ColumnFloats("score").Map(math.Abs).Mean()
				return FrameFromValues(dataframe.Len(), dataframe.Frame(0).Value(3).AsFloat64(), mean)
			}, "n", "fuck_score", "score")
		},
	).Filter(func(frame Frame) bool {
		var count = frame.Value(1).AsFloat64()
		var score = frame.Value(2).AsFloat64()
		return count < 2 && score < 200
	})
	var word_score = reddit.ColumnFloats("fuck_score")
	var scores = reddit.ColumnFloats("score")
	fmt.Printf("%v\n%v\n%v\n", reddit.Scheme(), word_score, scores)
	var scatter, errScatter = plotter.NewScatter(NewPoints(
		word_score, scores))
	if errScatter != nil {
		panic(errScatter)
	}

	var plt, errPlt = plot.New()
	if errPlt != nil {
		panic(errPlt)
	}
	//	var stats = reddit.ColumnFloats("score").Stats()
	plt.Y.Min = 0
	plt.Y.Max = 200
	plt.X.Min = 0
	plt.X.Max = 2
	plt.Add(scatter)
	if err := plt.Save(15*vg.Centimeter, 10*vg.Centimeter, "fuck_score.png"); err != nil {
		panic(err)
	}
}
