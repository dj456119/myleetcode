/*
 * @Descripttion:
 * @version:
 * @Author: cm.d
 * @Date: 2021-11-02 21:09:29
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-02 21:36:34
 */
/*
 * @Descripttion:我们提供了一个类：

public class Foo {
  public void first() { print("first"); }
  public void second() { print("second"); }
  public void third() { print("third"); }
}
三个不同的线程 A、B、C 将会共用一个 Foo 实例。

一个将会调用 first() 方法
一个将会调用 second() 方法
还有一个将会调用 third() 方法
请设计修改程序，以确保 second() 方法在 first() 方法之后被执行，third() 方法在 second() 方法之后被执行。

 * @version:
 * @Author: cm.d
 * @Date: 2021-11-02 21:09:29
 * @LastEditors: cm.d
 * @LastEditTime: 2021-11-02 21:09:29
*/
package myleetcode

import "fmt"

type Foo struct {
}

var chanAToB chan int
var chanBToC chan int

func (Foo) first(printFirst func()) {
	printFirst()
	chanAToB <- 0
}

func (Foo) second(printSecond func()) {
	<-chanAToB
	printSecond()
	chanBToC <- 0
}

func (Foo) third(printThird func()) {
	<-chanBToC
	printThird()
}

func Exe() {
	chanAToB = make(chan int, 1)
	chanBToC = make(chan int, 1)
	f := Foo{}
	funA := func() {
		f.first(func() { fmt.Print("first") })
	}
	funB := func() {
		f.second(func() { fmt.Print("second") })
	}
	funC := func() {
		f.third(func() { fmt.Print("third") })
	}

	go funA()
	go funB()
	go funC()
}
