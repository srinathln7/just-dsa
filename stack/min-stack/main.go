package main 

type MinStack struct {
  Stack []Element
}

 type Element struct {
     val int 
     min int 
 }


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    var min int
    if len(this.Stack) == 0 {
        min = val
    } else { 
        topElement := this.Stack[len(this.Stack)-1]   
        if val <= topElement.min {
            min = val
        } else {
            min = topElement.min
        }
    }
    this.Stack = append(this.Stack, Element{val: val, min: min})
}


func (this *MinStack) Pop()  {
    this.Stack = this.Stack[:len(this.Stack)-1]
}


func (this *MinStack) Top() int {
    topElement := this.Stack[len(this.Stack)-1]
    return topElement.val
}


func (this *MinStack) GetMin() int {
    topElement := this.Stack[len(this.Stack)-1]
    return topElement.min
}


