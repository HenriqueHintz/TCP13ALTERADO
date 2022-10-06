package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	//"strings"
	//"time"
)

var variosClientes map[*Cliente]int

type Cliente struct {
	entrada chan string
	saida chan string
	reader   *bufio.Reader
	writer   *bufio.Writer
	conn     net.Conn
	conexao  *Cliente
}

func (cliente *Cliente) Read() {
	for {
		line, err := cliente.reader.ReadString('\n')
		if err == nil {
			if cliente.conexao != nil {
				cliente.conexao.saida <- line
			}
			fmt.Println(line)
		} else {
			break
		}

	}

	cliente.conn.Close()
	delete(variosClientes, cliente)
	if cliente.conexao != nil {
		cliente.conexao.conexao = nil
	}
	cliente = nil
}

func (cliente *Cliente) Write() {
	for data := range cliente.saida {
		cliente.writer.WriteString(data)
		cliente.writer.Flush()
	}
}

func (cliente *Cliente) Listen() {
	go cliente.Read()
	go cliente.Write()
}

func NovoCliente(conexao net.Conn) *Cliente {
	writer := bufio.NewWriter(conexao)
	reader := bufio.NewReader(conexao)

	cliente := &Cliente{
		entrada: make(chan string),
		saida: make(chan string),
		conn:     conexao,
		reader:   reader,
		writer:   writer,
	}
	cliente.Listen()

	return cliente
}

func teste(mensagem string) string {
	// eliminando a quebra de parágrafo e de linha
	msg := mensagem[0:(len(mensagem) - 2)]
	switch msg {
	case "1":
		fmt.Println("R$ 10,00")
		return "R$ 10,00"
	//	time.Sleep(1 * time.Second)
	case "2":
		fmt.Println("R$ 15,00")
		return "R$ 15,00"
	//	time.Sleep(1 * time.Second)
	case "3":
		fmt.Println("R$ 18,00")
		return "R$ 18,00"
	//	time.Sleep(1 * time.Second)
	case "4":
		fmt.Println("R$ 30,00")
		return "R$ 30,00"
	//	time.Sleep(1 * time.Second)
	case "5":
		fmt.Println("R$ 35,00")
		return "R$ 35,00"
	//	time.Sleep(1 * time.Second)
	case "6":
		fmt.Println("R$ 37,00")
		return "R$ 37,00"
	//	time.Sleep(1 * time.Second)
	case "7":
		fmt.Println("R$ 55,00")
		return "R$ 55,00"
	//	time.Sleep(1 * time.Second)
	case "8":
		fmt.Println("R$ 60,00")
		return "R$ 60,00"
	//	time.Sleep(1 * time.Second)
	case "9":
		fmt.Println("62,00")
		return "R$ 62,00"
	//	time.Sleep(1 * time.Second)
	case "10":
		fmt.Println("65,00")
		return "R$ 65,00"
	//	time.Sleep(1 * time.Second)
	case "0":
		fmt.Println("saindo.....")
		os.Exit(0)
		return "\n"
	default:
		return "Não identificado"

	}

}

func main() {

	variosClientes = make(map[*Cliente]int)
	listener, _ := net.Listen("tcp", ":8081")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		cliente := NovoCliente(conn)
		for clienteList, _ := range variosClientes {
			if clienteList.conexao == nil {
				cliente.conexao = clienteList
				clienteList.conexao = cliente
				fmt.Println("Conectado...")
			}
		}
		variosClientes[cliente] = 1
		fmt.Println(len(variosClientes))
	}

}
