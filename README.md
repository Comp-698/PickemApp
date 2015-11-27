PickemApp
==========

PickemApp is a web based project for the Comp 698 UNHM class that allows users to bet on the outcomes of various sporting events. 

some super handy commands to keep *ahem* handy

steps to get rolling from almost any computer:

Download golang, Install it, and then setup your GOPATH
https://golang.org/doc/install
https://golang.org/doc/install#testing (setting up your GOPATH)

this assumes you have already forked the https://github.com/Comp-698/PickemApp

and use your Github user, in place of mine - which is Gargame11
tchadwick@artemis:~/workspace/go$ git clone https://github.com/Gargame11/PickemApp

tchadwick@artemis:~/workspace/go/PickemApp$ git remote add upstream https://github.com/Comp-698/PickemApp.git

run this often, both to make sure you are up to date and to make sure your ongoing changes don't break
tchadwick@artemis:~/workspace/go/PickemApp$ git pull upstream master; git push;

tchadwick@artemis:~/workspace/go/PickemApp$ go get -u github.com/Comp-698/PickemLib

tchadwick@artemis:~/workspace/go/PickemApp$ go build pickem.go

tchadwick@artemis:~/workspace/go/PickemApp$ go run pickem.go

Develop.
