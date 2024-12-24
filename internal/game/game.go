package game

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func Play(ctx context.Context,file *os.File) error {
	fileScanner := bufio.NewScanner(file)
	var score uint8



	go func(){
		<-ctx.Done()

		log.Println("your score is ",score)
	}()

	for fileScanner.Scan() {

		if errors.Is(fileScanner.Err(), io.EOF) {
			log.Println("your score is ", score)
			return nil
		}

		content := fileScanner.Text()
		quizAns := extract(content)

		fmt.Print(string(content[0]) + "+" + string(content[2]) + " = ")
		var ans string
		fmt.Fscan(os.Stdin, &ans)

		if ans != quizAns {
			log.Println("your score is ", score)
			os.Exit(-1)
		}
		score += 1
	}

	return nil
}

func extract(s string) string {
	leng := len(s) - 1
	ans := make([]byte, leng) // and this is correct

	for leng >= 1 {
		if s[leng] == ',' {
			break
		}
		ans[leng-1] = s[leng]
		leng -= 1
	}

	a := string(ans[len(ans)-1:])

	return a
}
