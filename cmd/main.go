package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

const copo = 200 

var contador int 

func main() {
	// Criar um Scheduler com o fuso horário local
	t := cron.New(cron.WithLocation(time.Local))

	// Adicionar uma função ao Scheduler com a expressão cron
	t.AddFunc("*/45 * * * *", func() {
		message := fmt.Sprintf("Beba agua, ja são: %s", time.Now().Format("15:04"))
		fmt.Println(message)
		contador++ // incrementar o contador
		fmt.Println(countFinal(contador))
	})

	t.Start()

	// Criar uma goroutine que verifica se o horário é 17:15 e para o Scheduler se for
	// esperar um minuto parar o Scheduler enviar um sinal para o canal done sair do loop
	done := make(chan struct{})
	defer close(done)

	go func() {
		for {
			time.Sleep(time.Minute) 
			if time.Now().Hour() == 17 && time.Now().Minute() == 15 {
				t.Stop()           
				done <- struct{}{}
				break              
			}
		}
	}()
	select {
	case <-done:
		return
	}
	
}

func countFinal(value int) string {
	format := fmt.Sprintf("se voce bebeu um copo de agua que geralmente tem %dml, no total voce bebeu durante a execução: %dml", copo, value*copo)
	return format
}
