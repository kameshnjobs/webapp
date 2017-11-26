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

