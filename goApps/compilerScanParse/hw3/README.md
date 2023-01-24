How To Implement Parsor on linux server

Begin with the following code to get set up:

% scp -r hw3-btwhitehouse2 bwhitehouse@linux.cs.uchicago.edu:hw3-btwhitehouse2 
enter password:

Logon to linux machine: % ssh bwhitehouse@linux.cs.uchicago.edu enter password: $ module load golang/1.19 $ go version go version go1.19 linux/amd64

I broke this homework into 4 seperate packages: rdp, token, scanner and driver. RDP (Recursive Decent Parser): starts with the Cal grammar and builds top down with a right recursive grammar will return a Bool that Driver will print or an error if one is encountered  Token: stores the information for the Token Type Driver: is the main package for this program, you will need to call it to operate the program will print False if no errors are found. Scanner: has the meat or spaghetti code for creating tokens in this program.

Now cd to hw3*/hw3/driver You will need to run the following:$ go run driver.go simple.cal In order for this work you must have the .cal file or whatever file you are reading in the same folder As driver or have a good root pointing to that fil