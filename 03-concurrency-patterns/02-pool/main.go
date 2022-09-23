package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"
)

//Resource
type DBConnection struct {
	ID int
}

//io.Closer implementation
func (dbConnection *DBConnection) Close() error {
	fmt.Printf("Closing and discarding the resource (ID): %d\n", dbConnection.ID)
	return nil
}

var idCounter int

//factory

func DBConnectionFactory() (io.Closer, error) {
	idCounter++
	return &DBConnection{ID: idCounter}, nil
}

func main() {
	/*
		create a Pool instance (with the pool size, factory function)

		Resource => any object that implements "close()" method (io.Closer interface)

		When a resource is Acquire()d
			the pool will check if it has any resources
			if yes, return the resource from the pool
			else create a new resource (using the factory) and return

		When a resource is Release()d to the pool
			the pool will check if it is full
			if yes, the discard the resource (after 'closing' the resource)
			else, keep the resource to serve future requests (maintain the resource in the pool)

		When Close()ing the pool
			prevent anymore aqcuisition of the resourcer
			make sure all the resources are closed and discarded

		Notes:
			The number of requests for the resources CAN exceed the pool size
			When a resource is acquired by the client, the same resource SHOULD NOT be given to another client until the resource is released back to pool


		Functions:
			New(poolSize, factory) => pool

		Methods (pool)
			Aquire() => resource
			Release(resource)
			Close()

	*/

	p, err := pool.New(5, DBConnectionFactory)
	if err != nil {
		log.Fatalln(err)
	}

	wg := &sync.WaitGroup{}
	clientCount := 10
	wg.Add(clientCount)
	for client := 1; client <= clientCount; client++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(client)
	}
	wg.Wait()

	fmt.Println("Second batch of operation, hit ENTER to start")
	fmt.Scanln()
	wg.Add(4)
	for client := 1; client <= 4; client++ {
		go func(client int) {
			doWork(client, p)
			wg.Done()
		}(client)
	}
	wg.Wait()
	p.Close()
}

func doWork(id int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err {
		log.Fatalln(err)
	}
	fmt.Printf("Worker : %d Acquired : %d\n", id, conn.(*DBConnection).ID)
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	fmt.Printf("Worker %d Done. Releasing : %d\n", id, conn.(*DBConnection).ID)
	p.Release(conn)
}
