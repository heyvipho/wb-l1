package main

import "fmt"

type Operator interface {
	Call(number string) string
}

type Client struct{}

func (v *Client) Call(o Operator, number string) string {
	return o.Call(number)
}

type AMobile struct{}

func (v *AMobile) CallA(number string) string {
	return fmt.Sprintf("You called %v through AMobile", number)
}

type AdapterAMobile struct {
	AMobile
}

func (v AdapterAMobile) Call(number string) string {
	return v.CallA(number)
}

type BMobile struct{}

func (v *BMobile) CallB(number string) string {
	return fmt.Sprintf("You called %v through BMobile", number)
}

type AdapterBMobile struct {
	BMobile
}

func (v AdapterBMobile) Call(number string) string {
	return v.CallB(number)
}

func main() {
	client := Client{}

	aMobile := AMobile{}
	fmt.Println(client.Call(AdapterAMobile{aMobile}, "112"))

	bMobile := BMobile{}
	fmt.Println(client.Call(AdapterBMobile{bMobile}, "112"))
}
