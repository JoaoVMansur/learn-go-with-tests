package main

import (
	"testing"
)


func TestWallet(t *testing.T){

    t.Run("Deposit", func(t *testing.T) {

        wallet := Wallet{}

        wallet.Deposit(Bitcoin(10))

        assertBalance(t, wallet,Bitcoin(10))
    })
    t.Run("Withdraw", func(t *testing.T){

        wallet := Wallet{balance: Bitcoin(100)}

        err := wallet.Withdraw(Bitcoin(80))
        
        assertNoError(t, err)
        assertBalance(t, wallet, Bitcoin(20))
    })
        
    t.Run("Withdraw with insuffivient funds", func(t *testing.T){
        
        wallet := Wallet{balance: Bitcoin(20)}

        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, Bitcoin(20))
        
    })

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin){
    t.Helper() 
    got  := wallet.Balance()
    if got != want{
        t.Errorf("got %s wanted %s", got, want)
    }
}

func assertError(t testing.TB, got error, want error){
    t.Helper()
    if got == nil{
        t.Fatal("wanted an error but didn't get one")
    }
    if got != want{

        t.Errorf("got %q wanted %q", got, want)

    }
}

func assertNoError(t testing.TB, got error){
    t.Helper()

    if got != nil{
        t.Fatal("got an error but it wasn't supposed to")
    }


}


