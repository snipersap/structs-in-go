# ReadMe
This repo is the project *structs-in-go* as described in the course  [Go: The Complete Developer's Guide (Golang)](https://udemy.com/course/go-the-complete-developers-guide/). 

This repo creates 2 structs, `person` and `contactInfo`, where `contactInfo` is used as an element inside `person`. In the `main` function, the following is then defined: 
1. Different ways to initialize the variables of type struct.
2. The recommened way to initialize is to use the element names instead of depending on the order of element definition inside the struct.
3. Assign values to nested struct.
4. Update structs inside a function, even though go passes all arguements by value.
5. Go shortcuts available, to reduce the use of `value at` operator `*` to assign values to structs inside functions.  

## How to test the repo?
Run the files `main.go` using the command 
`go run main.go`.   

Run the tests with 
`go test`

### Test Result:
PASS  
ok      structs-in-go   0.106s  

In case of confusion, check the commits. 

## Expected output
Last updated 13.09.2023  
>Alex: {Alex Hardy {h@h.com 99888}}  
Sandy: {Sandy Homie {sandy.homie@gmail.com 788999}}  
Priti normal: {  { 0}}  
Priti formatted:{firstName: lastName: contact:{email: pincode:0}}  
Priti with contactInfo, formatted:{firstName:Priti lastName:Sharma contact:{email:priti@sharma.com pincode:23099}}  
Update Sandy'S first name {Sandy Homie {sandy.homie@gmail.com 788999}}  
Update Sandy's name with pointer: {Cindy Klara {sandy.homie@gmail.com 788999}}  
Update Sandy's name without passing pointer {Sandeep Kaur {sandy.homie@gmail.com 788999}}  
Arun(init): {Arun Jaitley {a@jaitley.com 67556}}  
Arun, udpated email without using *: {Arun Jaitley {j@arun.com 67556}}  

## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
