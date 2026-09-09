# Project #1: p1-go

* Author: Cameron Kadre
* Class: CS354 Section #
* Semester: Fall 2026

## Overview



## Reflection

This project was challenging, but also pretty fun, and I do like the 
basic structure of go. After a while of struggling with go, it suddenly
got a lot easier (in some ways). One thing that really helped me was 
discovering (by reading the class slides with more attention) that 
lowercase type / method names denote private visibility. It was very 
frustrating before that to have no idea why I kept getting the error
that my methods were unimported, when I could very clearly see at the 
import statement that they should be. 

One thing that I still struggled with my implementation of was using
Stringer. I honestly couldn't quite figure out how that worked, for 
some reason, and so I honestly just gave up and created ToString 
methods for each type. I'm not sure if that was intended or not, 
but it's what I ended up with. Another thing was my implementation
of the Bank class. The map that was hinted in the project specs was
way too frustrating, so I just used a slice. Because of the time
I'd already spent fighting with that, I was pretty lazy and just 
had the length of the Bank slice hardcoded into the constructor. 

## Compiling and Using

Run the code using go run main.go, or by running the test script
run-test.sh

## Results

N/A

## Sources used

The GO language official documentation

----------