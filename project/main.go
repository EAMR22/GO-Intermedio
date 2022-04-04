package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}

	go d.Dispatch()
}

// Esta funcion es la que se va a encargar de manejar las peticiones http del servidor.
func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay")) // FormValue nos permite acceder a todo lo que se esta enviando del request.
	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid Value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)

	dispatcher.Run()

	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
