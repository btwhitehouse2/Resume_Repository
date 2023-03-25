#!/bin/bash 

module load golang/1.19
go mod tidy

cd grammars 

source generate.sh 

cd ..
 
go build -o golitec golite/golite 