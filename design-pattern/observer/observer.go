package main

/*
 * 观察者模式，顾名思义，是对象或者线程对某一对象或者线程进行持续性的关注，一旦被关注的对象或线程发生任何改变，能够及时的通知观察者，已做出正确的应对。
 * 观察者模式在开发中非常常见，诸如某些管理系统，或者电子商务系统，创造观察者线程对需要关注的线程进行关注，一旦某种事件发生，则可以通知观察者做出反应。
 * 比如某工厂的生产管理系统，一旦某些重要指标超过警戒线，需要及时的通知观察者线程，可以依据不同的问题状况构建不同的观察者，不同的观察者可以以不同的方
 * 式或渠道通知用户，比如发送短信，发送邮件，推送手机APP消息等，避免生产事故的发生。
 */

import "fmt"

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}

func main() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode")
	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}
