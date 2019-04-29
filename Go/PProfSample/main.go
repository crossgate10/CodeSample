package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"sync"
)

const (
	output     = "output.png"
	cpuprofile = "cpu.prof"
	memprofile = "mem.mprof"
	trcprofile = "trc.trace"
	width      = 2048
	height     = 2048
)

func norm(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}

func pixel(i, j, width, height int) color.Color {
	const complexity = 1024
	xi := norm(i, width, -1.0, 2)
	yi := norm(j, height, 1.0, 1)
	const maxI = 1000
	x, y := 0., 0.
	for i := 0; (x*x+y*y < complexity) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}
	return color.Gray{uint8(x)}
}

func createDirectly(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	wg.Add(width * height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			m.Set(i, j, pixel(i, j, width, height))
		}
	}
	wg.Wait()
	return m
}

// With goroutine each pixel
func createEachPixel(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	wg.Add(width * height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			go func(i, j int) {
				m.Set(i, j, pixel(i, j, width, height))
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	return m
}

// With goroutine each row
func createEachRow(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	wg.Add(width * height)
	for i := 0; i < width; i++ {
		go func(i int) {
			for j := 0; j < height; j++ {
				m.Set(i, j, pixel(i, j, width, height))
				wg.Done()
			}
		}(i)
	}
	wg.Wait()
	return m
}

// With worker pattern
func createWithWorker(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	numWorkers := 4

	type Pixel struct{ x, y int }

	c := make(chan Pixel)

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			for px := range c {
				m.Set(px.x, px.y, pixel(px.x, px.y, width, height))
			}
			wg.Done()
		}()
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c <- Pixel{i, j}
		}
	}
	close(c)
	wg.Wait()
	return m
}

// With non-blocking chan (buffered chan)
func createWithNonBlockingChan(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	numWorkers := 4

	type Pixel struct{ x, y int }

	c := make(chan Pixel, width*height)

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			for px := range c {
				m.Set(px.y, px.y, pixel(px.x, px.y, width, height))
			}
			wg.Done()
		}()
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			c <- Pixel{i, j}
		}
	}
	close(c)
	wg.Wait()
	return m
}

// With worker each rows
func createWithWorkerAndRow(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	numWorkers := 4

	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			for i := range c {
				for j := 0; j < height; j++ {
					m.Set(i, j, pixel(i, j, width, height))
				}
			}
			wg.Done()
		}()
	}

	for i := 0; i < width; i++ {
		c <- i
	}
	close(c)
	wg.Wait()
	return m
}

func main() {
	cpuf, err := os.Create(cpuprofile)
	defer cpuf.Close()
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	trcf, err := os.Create(trcprofile)
	defer trcf.Close()
	if err != nil {
		panic(err)
	}
	trace.Start(trcf)
	defer trace.Stop()

	pngf, err := os.Create(output)
	defer pngf.Close()
	if err != nil {
		panic(err)
	}
	img := createWithWorkerAndRow(width, height)
	if err = png.Encode(pngf, img); err != nil {
		panic(err)
	}
	
	memf, _ := os.Create(memprofile)
	defer memf.Close()
	pprof.WriteHeapProfile(memf)
	memf.Close()
}
