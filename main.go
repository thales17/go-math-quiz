package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const promptTmpl string = `
<html>
	<head>
		<title>Go Math Quiz</title>
		<style>
body {
	color: white;
	background-color: black;
}

main {
	max-width: 38rem;
	padding: 2rem;
	margin: auto;
	text-align: center;
}

input {
max-width: 4rem;
}
		</style>
	</head>
	<body>
		<main>
			<p>%d</p>
			<p>%s</p>
			<p>%d</p>
			<form action="/answer" method="post">
				<input type="hidden" name="a" value="%d">
				<input type="hidden" name="op" value="%s">
				<input type="hidden" name="b" value="%d">
				<label>=</label>
				<input type="text" name="answer">
				<button>Submit</button>
			</form>
		</main>
	</body>
</html>
`

func calcAnswer(a, b int, op string) (float64, error) {
	answer := 0.0
	switch op {
	case "+":
		answer = float64(a) + float64(b)
	case "-":
		answer = float64(a) - float64(b)
	case "*":
		answer = float64(a) * float64(b)
	case "/":
		answer = float64(a) / float64(b)
	default:
		return -1, errors.New(fmt.Sprintf("Invalid operator, %s", op))
	}

	return answer, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ops := []string{"+", "-"}
		op := ops[rand.Intn(len(ops))]
		a := rand.Intn(10)
		b := rand.Intn(10)
		if op == "-" {
			tmp := a
			a = int(math.Max(float64(a), float64(b)))
			b = int(math.Min(float64(tmp), float64(b)))
		}
		fmt.Fprintf(w, fmt.Sprintf(promptTmpl, a, op, b, a, op, b))
	})

	http.HandleFunc("/answer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Invalid method")
			return
		}

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)

		aVal := r.FormValue("a")
		op := r.FormValue("op")
		bVal := r.FormValue("b")
		answerVal := r.FormValue("answer")

		fmt.Fprintf(w, "a = %s\n", aVal)
		fmt.Fprintf(w, "op = %s\n", op)
		fmt.Fprintf(w, "b = %s\n", bVal)
		fmt.Fprintf(w, "answer= %s\n", answerVal)

		a, err := strconv.Atoi(aVal)
		if err != nil {
			log.Println(err)
			return
		}

		b, err := strconv.Atoi(bVal)
		if err != nil {
			log.Println(err)
			return
		}

		answer, err := strconv.ParseFloat(answerVal, 64)
		if err != nil {
			log.Println(err)
			return
		}

		answerCalc, err := calcAnswer(a, b, op)
		if err != nil {
			log.Println(err)
			return
		}

		if answerCalc == answer {
			fmt.Fprintf(w, "\n\nCorrect!\n")
		} else {
			fmt.Fprintf(w, fmt.Sprintf("\n\nWrong!\nCorrect Answer:%.1f", answerCalc))
		}
	})

	port := 8666
	log.Printf("Starting server on %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
