 package main

import (
	"fmt" 
	"net/http"
	"os"    
    "time" 
    "strconv"  
    "os/exec"
    "runtime"
)
//***********************************************************
//
//
//***********************************************************
type msStruct struct { 
    flag  chan string
} 
  
/* curl "http://127.0.0.1:9080/?start=yes&num=3" */
func (mss *msStruct) commMasterReceiveCommand(w http.ResponseWriter, r *http.Request) { 
  
    r.ParseForm()
    if r.Method == "GET" { 
        if clo := r.FormValue("close"); clo == "yes"{   
            num := r.FormValue("num")     
            c ,_ :=   strconv.Atoi(num)
            mss.supervisorctl_Command(0, c) 
        }
        if clo := r.FormValue("start"); clo == "yes"{   
            num := r.FormValue("num")         
            c ,_ :=   strconv.Atoi(num)
            mss.supervisorctl_Command(1, c)  
        }  
    }else if r.Method == "POST" {                

    }
 
    fmt.Fprintf(w, "OK")  
}  
//***********************************************************
//
//
//***********************************************************
func (mss *msStruct) supervisorctl_Command(commandFlag int, count int){
    
    if(commandFlag==0){
        for i := 0; i < count; i++{   
            _, err := exec.Command("supervisorctl", "stop", "slave:"+ strconv.Itoa(i)).Output()
            if err != nil {
                fmt.Printf("%s", err)
            }
            time.Sleep(time.Millisecond * 200)   
        }  
    }
    if(commandFlag==1){
        for i := 0; i < count; i++{   
            _, err := exec.Command("supervisorctl", "start", "slave:"+ strconv.Itoa(i)).Output()
            if err != nil {
                fmt.Printf("%s", err)
            }
            time.Sleep(time.Millisecond * 800)    
        }  
    }
}

func (mss *msStruct) getStartMaster( ) {
    
    tick  := time.NewTicker(time.Millisecond * 1000)
    flag := true  
    for {
        select {
            case <- tick.C:
                mss.flag <- "2"  
                tick.Stop()
                flag = false 
        }        
        if(!flag){
            break
        }  
    }
}

func (mss *msStruct) runMaster(){
 
    for {
        select {
            case  flag := <-mss.flag:                 

                if( flag=="2" ){  
                /***********
                Write code you would like to run in this master.....
                ************/     
                } 
        }
        runtime.Gosched()
    }

}

func (mss *msStruct) getStartSlave( ) {
    
    tick  := time.NewTicker(time.Millisecond * 1000)
    flag := true  
    for {
        select {
            case <- tick.C:
                mss.getStart()
                tick.Stop()
                flag = false 
        }        
        if(!flag){
            break
        }  
    }
}
 
func (mss *msStruct) runSlave(){ 

    for {   
        select {
            case  flag := <-mss.flag: 
                if( flag=="1" ){    
                /***********
                Write code you would like to run in this slave.....
                ************/                
                     fmt.Printf("gNum => %d \n", gNum)
                }                  
        }
        runtime.Gosched()
    }

}

func (mss *msStruct) runMS() {  
    
        if(gNum == -1){           
            go mss.getStartMaster()           
            mss.runMaster()        
        }else{      
            go mss.getStartSlave()  
            mss.runSlave()             
        } 
}   
 
func (mss *msStruct) getUpdate(){
    mss.flag <- "2"
}

func (mss *msStruct) getStart(){
    mss.flag <- "1"
}

func (mss *msStruct) getStop(){
    mss.flag <- "0"
}
//***********************************************************
//
//
//***********************************************************
var gNum int
var gPort int

func main() {
    
    runtime.GOMAXPROCS(2)

    arg_num := len(os.Args)
    gNum = -1
    gPort = 9080
 
    if(arg_num == 1){
    // gNum : -1 
        var mss msStruct
        mss.flag = make(chan string)       
    
        go mss.runMS()
            
        http.HandleFunc("/", mss.commMasterReceiveCommand)
        http.ListenAndServe(":"+ strconv.Itoa(gPort), nil)       
        
    }else if(arg_num == 2){
//slave:0                        STOPPED   Not started
//slave:1                        STOPPED   Not started 
        gNum ,_ = strconv.Atoi(os.Args[1])      
      
        var mss msStruct
        mss.flag = make(chan string)
       
        go mss.runMS()
  
        http.ListenAndServe(":"+ strconv.Itoa(gPort+gNum+1), nil)
       
    }

    fmt.Println("Run MS....." )
}
 