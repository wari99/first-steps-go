package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Pedido struct { //Tipo pedido: a encomenda
	ID int
}

type Motorista struct { //Tipo motorista: quem vai entregar fisicamente a encomenda
	ID int
}

type logistica_transporte struct { // Tipo logistica_transporte: quem vai ser encarregado do processo de entrega
	pedidos    chan Pedido 
	entregas   chan Pedido
	motoristas []Motorista
	wg         sync.WaitGroup
}

func (e *logistica_transporte) novoPedido(id int) { // novoPedido cria e envia um novo pedido

	pedido := Pedido{ID: id}
	e.pedidos <- pedido
	fmt.Printf("Pedido %d enviado\n", pedido.ID)
}

func (e *logistica_transporte) distribuirPedidos() { // distribuirPedidos atribui/destribui os pedidos entre os motoristas da empresa
	defer e.wg.Done()
	for pedido := range e.pedidos {
		motorista := e.escolherMotorista()
		fmt.Printf("Pedido %d atribuído ao motorista %d\n", pedido.ID, motorista.ID)
		e.entregas <- pedido
	}
	close(e.entregas) //O canal é fechado quando todas as entregas são concluidas
}

func (e *logistica_transporte) escolherMotorista() Motorista { // escolherMotorista escolhe um motorista aleatoriamente

	rand.Seed(time.Now().UnixNano())
	return e.motoristas[rand.Intn(len(e.motoristas))]
}

func (e *logistica_transporte) fazerEntregas() { // fazerEntregas representa o transporte das entregas até o destino
	defer e.wg.Done()
	for entrega := range e.entregas {
		fmt.Printf("Pedido %d entregue pelo motorista\n", entrega.ID)
		time.Sleep(time.Second * 2) // simulacao o tempo de entrega
	}
}

func main() {
	const numMotoristas = 3
	const numPedidos = 10

	empresa := logistica_transporte{
		pedidos:    make(chan Pedido),
		entregas:   make(chan Pedido),
		motoristas: make([]Motorista, numMotoristas),
	}

	for i := 0; i < numMotoristas; i++ { 	//Inicia os motoristas
		empresa.motoristas[i] = Motorista{ID: i + 1}
	}

	//iniciando a distribuição de pedidos
	empresa.wg.Add(1)
	go empresa.distribuirPedidos()

	//iniciando a entrega de pedidos
	empresa.wg.Add(1)
	go empresa.fazerEntregas()

	//simulando novos pedidos sendo feitos
	for i := 1; i <= numPedidos; i++ {
		empresa.novoPedido(i)
		time.Sleep(time.Second) //intervalo entre realizacao de pedidos pelos clientes
	}

	close(empresa.pedidos) //fechando esse canal apos entregar todas as encomendas
	empresa.wg.Wait()
}
