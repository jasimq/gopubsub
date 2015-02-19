package main

import (
  "fmt"
  "log"
  "time"
  "strconv"
)

type Controller struct {
  chans map[string][]chan string
  capacity int 
}



var controller Controller

func main() {
  log.Println("Hello World!")
  
  controller.init()
  //controller.c = make(chan string)
  
  go func() {
    log.Println("Starting listener 1")
    
    myChan := make(chan string)
    controller.Sub("topic", myChan)
    
    for{ 
      msg := <-myChan
      log.Println("Listener 1 received " + msg)
    }
    
    
  }()
  
  go func() {
    log.Println("Starting listener 2")
    
    myChan := make(chan string)
    controller.Sub("topic", myChan)
    
    for{ 
      msg := <-myChan
      log.Println("Listener 2 received " + msg)
    }
  }()
  
  go func() {
    log.Println("Starting listener 3")
    
    myChan := make(chan string)
    controller.Sub("topic", myChan)
    
    for{ 
      msg := <-myChan
      log.Println("Listener 2 received " + msg)
    }
  }()
  
  time.Sleep(1 * time.Second)
  
  log.Printf("Length: ", len(controller.chans["topic"]))
  
  go func() {
    log.Println("Starting publisher 1")
    controller.Pub("topic", "jasim")
  }()
  
  time.Sleep(1 * time.Second)
  
  
  go func() {
    log.Println("Starting publisher 2")
    controller.Pub("topic", "qazi")
  }()
  
  time.Sleep(10 * time.Second)
  
}


func(c* Controller)Sub(topic string, ch chan string) {

  if _, ok := c.chans[topic]; !ok {
    log.Printf("topic '%s' does not exist. Creating...", topic)
      c.chans[topic] = make([]chan string, 0)
  }
  
  c.chans[topic] = append(c.chans[topic], ch)
  //log.Printf("append Length: ", len(c.c["topic"]))
}


func(c* Controller)Pub(topic string, msg string) {
  
  if _, ok := c.chans[topic]; !ok {
    log.Printf("topic '%s' does not exist. Exiting...", topic)
      return
  }
  
  for i, ch := range c.chans[topic] {
    log.Printf("Pushing to ch %d", i)
    ch <- fmt.Sprintf("%s%s", msg , strconv.Itoa(i))
  }
}

func (c* Controller) init() {
  c.chans = make(map[string][]chan string)
  c.capacity = 10
}

