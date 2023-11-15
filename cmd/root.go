package cmd

import (
	"Counter/Packages"
	"fmt"
	"os"
	"time"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Counter",
	Short: "File Content Counter",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(readFile)
	readFile.Flags().IntP("routines", "r", 2, "Num of go routines")
	readFile.Flags().String("path", "file.txt", "File to  Path")
}

var readFile = &cobra.Command{
	Use:   "ReadFile",
	Short: "File Content Counter",
	Long:  `This Command will read the given file and count the num of words, vowels, punctuations and lines in the file `,
	Run: func(cmd *cobra.Command, args []string) {
		goRoutine, _ := cmd.Flags().GetInt("routines")
		filePath, _ := cmd.Flags().GetString("path")
		fmt.Print("\n")
		fmt.Printf("Routines are : %v\n", goRoutine)
		fmt.Printf("File Path is : %v\n", filePath)
		fileReader(goRoutine, filePath)
	},
}

func fileReader(routine int, filepath string) {
	var lines, words, vowels, punctuations int
	fileContent, err := read(filepath)
	if err != nil {
		fmt.Print("Error : ", err)
	}
	startTime := time.Now()

	channel := make(chan Packages.Counter)

	chunk := len(fileContent) / routine

	fmt.Printf("File Chunks = %v ", routine)
	fmt.Printf("\n")

	for i := 0; i < routine; i++ {
		start := i * chunk
		end := (i + 1) * chunk
		go Packages.Count(fileContent[start:end], channel)
	}

	for i := 0; i < routine; i++ {
		counts := <-channel
		fmt.Printf("No of Words of Chunk %d: %d \n", i+1, counts.WordCount)
		fmt.Printf("No of Lines of Chunk %d: %d \n", i+1, counts.LineCount)
		fmt.Printf("No of Vowels of Chunk %d: %d \n", i+1, counts.VowelCount)
		fmt.Printf("No of Punctuation of Chunk %d: %d \n", i+1, counts.PunctuationCount)
		fmt.Printf("\n\n")
		lines = lines + counts.LineCount
		words = words + counts.WordCount
		vowels = vowels + counts.VowelCount
		punctuations = punctuations + counts.PunctuationCount
	}

	fmt.Printf("Execution time: %v\n", time.Since(startTime))
	fmt.Printf("Total Num of Lines: %v\n", lines)
	fmt.Printf("Total Num of Words: %v\n", words)
	fmt.Printf("Total Num if Vowels: %v\n", vowels)
	fmt.Printf("total Num of Punctuations: %v\n", punctuations)

}

func read(path string) (string, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	Content := string(fileContent)
	return Content, nil
}
