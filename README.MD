# BASIC REST API BUILD WITH GO STANDARD LIB
utilizing go 1.22 newest feature for building rest api with pure http/net standard library. 

## api explanation 
simulating basic user tweeting something , then we can like or commenting that tweet. I'll add some authentication to make it more fancier

## Project structure Explanation
 - bin is used for containing compiled application binaries
 - cmd/api containing application-specific code , his will include the code for running the server, reading and writing HTTP requests, and managing authentication
 - internal , it will contain the code for interacting with databases, data validation and so on
 - The remote directory will contain the configuration files and setup scripts for the production server

