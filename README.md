<pre>
cd src 
//here we have to insert all files
 go run Runner.go SimpleMath.go
 //or better 
 go run *.go
 //install 
 export GO=<absolute-path>
 go install *.go
 //build (it is better than install) from src
 cd src 
 go build -o bin/Runner *.go 
 //build from root 
 go build -o bin/Runner src/*.go
 //run 
 bin/Runner
</pre>