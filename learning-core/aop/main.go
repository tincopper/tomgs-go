package main

import (
	"fmt"
	"github.com/gogap/aop"
)

type Auth struct {
}

func (p *Auth) Login(userName, password string) bool {
	if userName == "zeal" && password == "gogap" {
		return true
	}
	return false
}

// use join point to get Args from real method
func (p *Auth) Before(jp aop.JoinPointer) {
	username := ""
	jp.Args().MapTo(func(u, p string) {
		username = u
	})

	fmt.Printf("Before Login: %s\n", username)
}

// the args is same as Login
func (p *Auth) After(username, password string) {
	fmt.Printf("After Login: %s %s\n", username, password)
}

// use join point to around the real func of login
func (p *Auth) Around(pjp aop.ProceedingJoinPointer) {
	fmt.Println("@Begin Around")

	ret := pjp.Proceed("fakeName", "fakePassword")
	ret.MapTo(func(loginResult bool) {
		fmt.Println("@Proceed Result is", loginResult)
	})

	fmt.Println("@End Around")
}

func main() {
	beanFactory := aop.NewClassicBeanFactory()
	beanFactory.RegisterBean("auth", new(Auth))

	aspect := aop.NewAspect("aspect_1", "auth")
	aspect.SetBeanFactory(beanFactory)

	pointcut := aop.NewPointcut("pointcut_1").Execution(`Login()`)
	aspect.AddPointcut(pointcut)

	aspect.AddAdvice(&aop.Advice{Ordering: aop.Before, Method: "Before", PointcutRefID: "pointcut_1"})
	aspect.AddAdvice(&aop.Advice{Ordering: aop.After, Method: "After", PointcutRefID: "pointcut_1"})
	aspect.AddAdvice(&aop.Advice{Ordering: aop.Around, Method: "Around", PointcutRefID: "pointcut_1"})

	gogapAop := aop.NewAOP()
	gogapAop.SetBeanFactory(beanFactory)
	gogapAop.AddAspect(aspect)

	proxy, err := gogapAop.GetProxy("auth")
	if err != nil {
		fmt.Println("login err:", err)
	}
	login := proxy.Method(new(Auth).Login).(func(string, string) bool)("zeal", "gogap")

	fmt.Println("login result:", login)
}
