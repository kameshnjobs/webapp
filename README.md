# webapp
This Webapp is used to perform healthcheck and create instance in google cloud
Description

WEBAPP

Installation

Install go with the below commands

Step 1 — Install Go 1.9

Login to your Redhat or its derivative system using ssh and upgrade to apply latest security updates there.

# yum update
Now download the Go language binary archive file using following link. To find and download latest version available or 32 bit version go to official download page.

# wget ttps://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
Now extract the downloaded archive and install it to the desired location on your system. For this tutorial, I am installing it under /usr/local directory. You can also put this under the home directory (for shared hosting) or other location.

# tar -xzf go1.9.2.linux-amd64.tar.gz
# mv go /usr/local
Step 2 — Setup Go Environment

Now you need to set up Go language environment variables for your project. Commonly you need to set 3 environment variables as GOROOT, GOPATH and PATH.

GOROOT is the location where Go package is installed on your system.

# export GOROOT=/usr/local/go
GOPATH is the location of your work directory. For example my project directory is ~/Projects/Proj1 .

# export GOPATH=$HOME/Projects/Proj1
Now set the PATH variable to access go binary system wide.

# export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
All above environment will be set for your current session only. To make it permanent add above commands in ~/.bash_profile file.

Step 3 — Verify Installation

At this step, you have successfully installed and configured go language on your system. First, use the following command to check Go version.

# go version

go version go1.9.2 linux/amd64
Now also verify all configured environment variables using following command.

# go env

----------------------

Copying the webapp.go
---------------------

Copy the webapp.go in the /home/<username home directory>
Copy index.html  in the /home/<username home directory>
 
 
 
  
Webapp.go code

Packges which needs to be imported
import (
    "fmt"
    "html/template"
    "log"
    "os/exec"
    "net/http"
    "strings"
    "time"
)

Create Function parses index.html and POST method takes the username, current time and creates instances
----------------------------------------------------------------------------------------------------------

func create(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Inside create method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("index.html")
      t.Execute(w, nil)
        fmt.Println("GET.............................")
    } else {
        fmt.Println("POST.......................")
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
        name := strings.ToLower(r.PostFormValue("username"))
        tt := time.Now()
        fmt.Println(tt.Format("20060102150405"))
        kk := tt.Format("20060102150405")
        output := strings.Join([]string{"gcloud compute instances create ",name,kk," --image-family rhel-7 --image-project rhel-cloud --zone us-central1-a"}, "")
       


 fmt.Println("The output command ", output)
        //gcloud compute instances create qq--image-family rhel-7 --image-project rhel-cloud--zone us-central1-a
        out, err1 := exec.Command("bash", "-c", output).Output()
        if err1 != nil {
                log.Fatal(err1)
            }
        str := fmt.Sprintf("%s", out)
        fmt.Fprintf(w,str)
        fmt.Printf("The output is %s\n", out)
        }

}

Healtcheck function describes the healthcheck test1
---------------------------------------------------


func healthcheck(w http.ResponseWriter, r *http.Request) {
    fmt.Println("inside healcheck method method:", r.Method) //get request method
    if r.Method == "GET" {
    out1, err2 := exec.Command("bash", "-c", "gcloud compute health-checks describe testing1").Output()
        if err2 != nil {
                log.Fatal(err2)
            }
        str1 := fmt.Sprintf("%s", out1)
        fmt.Fprintf(w,str1)
        fmt.Printf("The output is %s\n", out1)
    }
}


main function sets the routing rule and calls the create function and healthcheck 
---------------------------------------------------------------------------------
func main() {
    http.HandleFunc("/",create) // setting router rule
    http.HandleFunc("/healthcheck",healthcheck)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
 
-------------------------------------------------------------- 
Author
-------


Name: Kamesh Narayanababu
Email: kameshnjobs@gmail.com

 
License
--------

Free License
