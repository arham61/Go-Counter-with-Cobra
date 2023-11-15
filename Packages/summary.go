package Packages

type Counter struct {
	LineCount,VowelCount,PunctuationCount,WordCount int 
}

func Count(fileContent string, channel chan Counter) {
	count := Counter{}
	for _, char := range fileContent {
		switch {
		case char == ' ' || char == '\t' || char == '\r' || char == '.' || char == ',' || char == ';' || char == ':' || char == '!' || char == '?' || char == '(' || char == ')' || char == '[' || char == ']' || char == '{' || char == '}':
			count.WordCount++
		case char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' || char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u':
			count.VowelCount++
		case char == '.' || char == '!' || char == '?' || char == ',' || char == ':' || char == ';' || char == '(' || char == ')' || char == '[' || char == ']' || char == '{' || char == '}':
			count.PunctuationCount++
		case char == '\n':
			count.LineCount++
		}
	}
	channel <- count
}