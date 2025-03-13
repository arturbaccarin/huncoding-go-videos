package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processando job %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println("Resultado: ", result)
	}
}

/*
func leakExample() chan int {
	ch := make(chan int)
	// não tem nada para parar a goroutine
	// não termina nunca de executar
	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()
	return ch
}

func main() {
	ch := leakExample()
	fmt.Println("Goroutine criada, mas nunca finalizada")
	time.Sleep(3 * time.Second)
	close(ch)
}
*/

/*
var mu sync.Mutex

func increment(mu *sync.Mutex, wg *sync.WaitGroup, counter *int) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mu.Lock()
		*counter++
		mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	counter := 0

	// Não é recomendado passar um ponteiro
	// para uma goroutine para alterar
	// o valor de uma variável compartilhada
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&mu, &wg, &counter) // tem várias goroutines que não terminaram
	}

	time.Sleep(2 * time.Second)
	fmt.Println(counter)
}
*/

/*
func increment(wg *sync.WaitGroup, counter *int) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		*counter++
	}
}

func main() {
	var wg sync.WaitGroup
	counter := 0

	// Não é recomendado passar um ponteiro
	// para uma goroutine para alterar
	// o valor de uma variável compartilhada
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg, &counter) // tem várias goroutines que não terminaram
	}

	time.Sleep(2 * time.Second)
	fmt.Println(counter)
}
*/

/*
func fetchURL(wg *sync.WaitGroup, url string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %s\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("URL: %s, Status: %d\n", url, resp.StatusCode)
}

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}

	for _, url := range urls {
		wg.Add(1)
		fetchURL(&wg, url)
	}

	fmt.Println("Done!")
}
*/
