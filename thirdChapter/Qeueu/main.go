package main

import "fmt"

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

type Queue []*Order

func (order *Order) New(priority, quantity int, product, customerName string) {
	order.customerName = customerName
	order.priority = priority
	order.product = product
	order.quantity = quantity
}

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended bool
		for i, addedOrder := range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func main() {
	var queue Queue
	queue = make(Queue, 0)
	var order1 *Order = &Order{}
	var priority1 int = 2
	var quantity1 int = 20
	var product1 string = "Computer"
	var customerName1 string = "Greg White"
	order1.New(priority1, quantity1, product1, customerName1)
	var order2 *Order = &Order{}
	var priority2 int = 1
	var quantity2 int = 10
	var product2 string = "Monitor"
	var customerName2 string = "John Smith"
	order2.New(priority2, quantity2, product2, customerName2)
	var order3 *Order = &Order{}
	var priority3 int = 3
	var quantity3 int = 34
	var product3 string = "mouse"
	var customerName3 string = "John Rock"
	order3.New(priority3, quantity3, product3, customerName3)
	queue.Add(order1)

	queue.Add(order2)
	queue.Add(order3)
	var i int
	for i = 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
