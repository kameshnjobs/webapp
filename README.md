# webapp
This Webapp GO application is used to perform healthcheck for the instances and create instance in google cloud

 
WEBAPP
--------

Please use the webapp_ver_5_final.go to execute the webappalication. Its the lates version

How to call the application in the browser
------------------------------------------
35.196.250.124/healthcheck ----------->Through GET method Implements a basic health check, returning HTTP status code 200 and a blank page.

35.196.250.124                ------> After entering the username and password displays ipaddres of the created instance and url will change to 35.196.250.124/v1/instances/create
The username and password is validation in javascript

Since this is a public ipaddress the if the go application is stopped the public address will change. And so the webapp.go is kept running




Installation
-------------

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
package main
import (
    "fmt"
    "html/template"
    "log"
    "os/exec"
    "net/http"
    "strings"
    "time"
        "regexp"
        "reflect"
         "io/ioutil"
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
        var err error
        var err1 error
        fmt.Println("POST.......................")
        r.ParseForm()
        // logic part of log in
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
         name := strings.ToLower(r.PostFormValue("username"))
        password := r.PostFormValue("password")
        tt := time.Now()
        fmt.Println(tt.Format("20060102150405"))
        kk := tt.Format("20060102150405")
        output := strings.Join([]string{"gcloud compute instances create ",name,kk," --image-family rhel-7 --image-project rhel-cloud --zone us-east1-b "}, "")
 fmt.Println("The output command ", output)
        //gcloud compute instances create qq--image-family rhel-7 --image-project rhel-cloud--zone us-central1-a
        out, err1 := exec.Command("bash", "-c", output).Output()
        if err1 != nil {
                log.Fatal(err1)
            }
        str := fmt.Sprintf("%s", out)
     //   fmt.Fprintf(w,str)
        fmt.Printf("The output is %s\n", str)




//fmt.Printf("The output is %s\n", out)
tr := fmt.Sprintf("%s", str)
fmt.Println(reflect.TypeOf(tr))
// result = strings.Split(tr, " ")
//fmt.Println(result)
re := regexp.MustCompile(`[ABCDEFGHIJKLMNOPQRSTUVWXYZ-abcdefghijklmnopqrstuvwxyz - ]`)
    // Split based on pattern.
    // ... Second argument means "no limit."
    result1 := re.Split(tr, -1)
fmt.Println(reflect.TypeOf(result1))
tr22 := fmt.Sprintf("%s", result1)
fmt.Println(reflect.TypeOf(tr22))
fmt.Println("printing ....", tr22)
 err = ioutil.WriteFile("output.csv",[]byte(tr22), 0644)
  if err != nil {
        panic(err)
     }
licom := `cat output.csv | sed 's/\|/ /'|awk '{print $7}' `
outt,errrr := exec.Command("bash", "-c",licom).Output()
 if errrr != nil {
    log.Fatal(errrr)
   }
   fmt.Printf("The output is %s\n", outt)
publicip := fmt.Sprintf("%s", outt)
fmt.Fprintf(w,"The instances public ip is :-")
fmt.Fprintf(w,publicip)
fmt.Fprintf(w,"The root password is Crealytics@123.")
fmt.Println(publicip)
c1 :="gcloud --quiet compute ssh kameshnj@"+name+kk+" --command="+`"echo -e 'Crealytics@123\nCrealytics@123' | sudo -S passwd && echo -e "Crealytics@123";sudo sed -re 's/^(PermitRootLogin)([[:s
pace:]]+)no/\1\2yes/' -i. /etc/ssh/sshd_config ; sudo sed -re 's/^(PasswordAuthentication)([[:space:]]+)no/\1\2yes/' -i. /etc/ssh/sshd_config`+" &&sudo service sshd restart"+` "` + " --zone us-
east1-b"
c2 := "gcloud --quiet compute ssh " +"root@"+name+kk+" --command="+`"sudo useradd ` + name+ "&&echo -e "+"'"+password+"\n"+password+"'"+" |  passwd "+ name+`&&sudo echo '`+name+`    ALL=(ALL)  
     ALL' >> /etc/sudoers `+`"`+ " --zone us-east1-b"
fmt.Println(c1)
fmt.Println(c2)




out, errr := exec.Command("bash", "-c", c1).Output()
    if errr != nil {
          log.Fatal(errr)
}
out, errrrr := exec.Command("bash", "-c", c2).Output()
    if errrrr != nil {
          log.Fatal(errrrr)
}
       }
}

Healtcheck function describes the healthcheck test1
---------------------------------------------------

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("           "))
}
func main() {
    http.HandleFunc("/",create) // setting router rule
    http.HandleFunc("/healthcheck",ServeHTTP)
    err := http.ListenAndServe(":80", nil) // setting listening port
    // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

main function sets the routing rule and calls the create function and healthcheck 
---------------------------------------------------------------------------------
func main() {
    http.HandleFunc("/",create) // setting router rule
    http.HandleFunc("/healthcheck",ServeHTTP)
    err := http.ListenAndServe(":80", nil) // setting listening port
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


GITHUB link
----------
https://github.com/kameshnjobs/webapp.git
