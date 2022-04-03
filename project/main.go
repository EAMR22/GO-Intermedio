package main

import (
	"fmt"
	"time"
)

// Job representa un trabajo a ejecutar, con un nombre y un número y un retraso.
type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

// El trabajador será nuestro trabajador amigable con la concurrencia.
type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

// Dispatcher es un despachador que enviará trabajos a los trabajadores.
type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

// NewWorker devuelve un nuevo trabajador con la identificación y el grupo de trabajadores proporcionados.
func NewWorker(id int, WorkerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: WorkerPool,
		QuitChan:   make(chan bool),
	}
}

// El método de inicio, inicia a todos los trabajadores.
func (w Worker) Start() {
	go func() {
		w.WorkerPool <- w.JobQueue

		select {
		case job := <-w.JobQueue:
			fmt.Printf("Worker with id %d Started\n", w.Id)
			fib := Fibonacci(job.Number)
			time.Sleep(job.Delay)
			fmt.Printf("Worker with id %d Finished with result %d\n", w.Id, fib)
		case <-w.QuitChan:
			fmt.Printf("Worker with id %d Stopped\n", w.Id)
		}
	}()
}

// Método de detención, detiene al trabajador.
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// NewDispatcher devuelve un nuevo Dispatcher con los maxWorkers proporcionados.
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

// El despacho enviará empleos a los trabajadores.
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}
