# Cron Job Beba água

Este é um programa simples em Go que lembra você de beber água a cada 45 minutos até 17:15h. Ele também calcula quanto de água você bebeu durante a execução do programa.

## O que foi usado

Este programa usa o pacote [github.com/robfig/cron/v3](https://github.com/robfig/cron) para criar um agendador de tarefas com base em uma expressão cron. Ele também usa goroutines e channels para verificar o horário atual e parar o agendador quando for 17:15.

## Como rodar

Para rodar este programa, você precisa ter o Go instalado na sua máquina. Você pode baixar o Go [aqui](https://golang.org/dl/).

Depois de instalar o Go, você pode clonar este repositório usando o comando:

```bash
git clone https://github.com/seu-usuario/cronjob-beba_agua.git
```
###
em seguida 
```bash
cd cronjob-beba_agua
go build -o lembrete-beber-agua cmd/main.go
```
ps: voce precisa ter o go instalado preferivelmente na versão 1.20.
