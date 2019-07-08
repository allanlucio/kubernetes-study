package main

import "testing"

func TestGreeting(t *testing.T){
  resultado := Greeting("Code.education Rocks!")
  if resultado != "<b>Code.education Rocks!</b>"{
    t.Errorf("Resultado esperado esperada: %s","<b>Code.education Rocks!</b>")
  }
}