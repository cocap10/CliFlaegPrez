package main

import (
	"fmt"
	"github.com/cocap10/flaeg"
	"os"
)

//Prez the Configuration Structure used by Flaeg
type Prez struct {
	Page     uint `short:"p" description:"Page number"`
	NextPage bool `short:"n" description:"Next page"`
}

func (p *Prez) error() {
	fmt.Printf("Oups, page %d unknown ...", p.Page)
}

func (p *Prez) printLogo() {
	const logo = `
FFFFFFFFFFFFFFFFFFF lllll                                              
F:::::::::::::::::F l:::l                                              
F:::::::::::::::::F l:::l                                              
FF:::FFFFFFFFF::::F l:::l                                              
  F::F       FFFFFF  l::l    æææææææææææ    ææææææææ       gggggg  gggg
  F::F               l::l    æ::::::::::æ  æ::::::::æ     g::::::gg:::g
  F:::FFFFFFFFFF     l::l    ææææææææ:::::æ:::ææææ:::ææ  g::::::::::::g
  F::::::::::::F     l::l            æ:::æ:::æ    æ:::æ g::::gggg::::gg
  F::::::::::::F     l::l     æææææææ::::æ::::ææææ::::æ g:::g    g:::g 
  F:::FFFFFFFFFF     l::l    æ:::::::::::æ:::::::::::æ  g:::g    g:::g 
  F::F               l::l   æ:::ææææ:::::æ:::ææææææææ   g:::g    g:::g 
  F::F               l::l  æ:::æ    æ::::æ::::æ         g::::g   g:::g 
FF::::FF            l::::l æ:::æ    æ::::æ:::::æ        g:::::gggg:::g 
F:::::FF            l::::l æ::::ææææ:::::æ::::::æææææ    g:::::::::::g 
F:::::FF            l::::l  æ::::::::::æ   æ::::::::æ     g::::::::::g 
FFFFFFFF            llllll   ææææææææææ     æææææææææ      gggggg::::g 
                                                                 g:::g 
                                                        gggg     g:::g 
                                                        g:::gg  gg:::g 
                                                         g::::gg:::::g 
                                                          g:::::::::g  
                                                           gg:::::gg   
                                                             ggggg   `

	fmt.Println(logo)
}

// New creats ans init a new Prez, return as reference
func New() *Prez {
	return &Prez{
		Page:     0,
		NextPage: false,
	}
}

func main() {
	prezCmd := &flaeg.Command{
		Name:                  "Prez",
		Description:           "This go program intrduce flaeg library in CLI using flaeg library ;)",
		Config:                New(),
		DefaultPointersConfig: New(),
		Run: func(InitalizedConfig interface{}) error {
			//Type assertions
			p, ok := InitalizedConfig.(*Prez)
			if !ok {
				return fmt.Errorf("Cannot convert the config into Configuration")
			}
			switch p.Page {
			case 0:
				p.printLogo()
			default:
				p.error()
				p.printLogo()
			}
			return nil
		},
	}
	f := flaeg.New(prezCmd, os.Args[1:])
	f.Run()

}
