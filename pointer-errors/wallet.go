package main

import (
	"errors"
	"fmt"
)



type Bitcoin int

var ErrInsufficientFunds  = errors.New("Withdraw atempt failed. amount greater than balance")

type Wallet struct{
    balance Bitcoin
}


func (w *Wallet) Deposit(amount Bitcoin){
        w.balance += amount
}

func (w *Wallet) Balance() Bitcoin{
    return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
    
    if amount > w.Balance(){
        return ErrInsufficientFunds
    }
    
    w.balance -= amount
    return nil
}

type Stringer interface{
    String() string
}

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}
