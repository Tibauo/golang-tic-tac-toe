# tic tac toe

Play tic tac toe
## Rules
[https://en.wikipedia.org/wiki/Tic-tac-toe](https://en.wikipedia.org/wiki/Tic-tac-toe)

## How to build
```
[Tibauo]$ go build --o tic-tac-toe
[Tibauo}$ ls
morpion  morpion.go  src  tic-tac-toe
```

## How to run
```
[Tibauo]$ ./tic-tac-toe 
Game
User 1 =  X 
User 2 =  O

       j
     0 1 2
  0 |_|_|_|
i 1 |_|_|_|
  2 |_|_|_|

Player 1 : 
i : 1
j : 2
       j   
     0 1 2 
  0 |_|_|_|
i 1 |_|_|X|
  2 |_|_|_|
```
## Win
```
Player 1 : 
i : 0
j : 1
You Win player :  1
       j   
     0 1 2 
  0 |O|X|O|
i 1 |_|X|X|
  2 |_|X|O|
```

## Debug
```
go test
You Win player :  1
You Win player :  1
You Win player :  1
You Win player :  1
You Win player :  1
You Win player :  1
You Win player :  1
PASS
ok  	morpion	0.002s
```
