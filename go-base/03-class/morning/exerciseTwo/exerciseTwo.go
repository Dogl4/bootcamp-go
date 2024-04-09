package exercise

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func ExerciseTwo() {
	fmt.Println("\n============ Exercicio 02 ============\n")

	readCsv()
}

func readCsv() {
	file, err := os.Open("./exerciseOne/productClear.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := tabwriter.NewWriter(os.Stdout, 20, 30, 1, '\t', tabwriter.AlignRight)

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	head := strings.Split(scanner.Text(), ",")

	for _, h := range head {
		fmt.Fprintf(reader, "%s\t", h)
	}

	fmt.Fprintln(reader)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		for i, v := range values {
			if i == 2 {
				v1, _ := strconv.ParseFloat(values[i-1], 64)
				v2, _ := strconv.ParseFloat(v, 64)
				fmt.Fprintf(reader, "%v\t", (v1 * v2))
			} else {
				fmt.Fprintf(reader, "%s\t", v)
			}
		}
		fmt.Fprintln(reader)
	}
	reader.Flush()
}
