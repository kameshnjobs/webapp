package main
import (
    "fmt"
    "html/template"
    "log"
    "os/exec"
    "net/http"
    "strings"
    "time"
)

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
func main() {

    http.HandleFunc("/",create) // setting router rule
    http.HandleFunc("/healthcheck",healthcheck)
 
    err := http.ListenAndServe(":9090", nil) // setting listening port
    // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
