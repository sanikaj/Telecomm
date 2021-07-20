package main

import (
    "database/sql"
    "fmt"
    "encoding/json"
    "log"
    "time"
    "context"
    "net/http"  
    "github.com/rs/cors"
    _ "github.com/go-sql-driver/mysql"
   
)

const (
    host     = "localhost"
    port     = 3306
    user     = "root"
    password = "password"
    dbname   = "db_telecomm"
  )


  type myData struct {
	name string
    email  string
    phonenumbers string
    customerid int64
    insertion string
}

type Customer struct {
    Customerid  int `json:"CustomerId"`
    Name string    `json:"Name"`
    Email  string `json:"Email"`
    CreateDate string `json:"CreatedDate"`
    Phones string `json:"Phones"`
    
}

func retrieveAllRecords()(map[int]Customer) {
    var (
        customer_id int
        full_name string
        email string
        create_date string
        phones string
    )

    db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db_telecomm")
                    if err != nil {
                        panic(err)
                    }

                sqlStatement_Customer :=`
                SELECT c.customer_id
                     , c.full_name
                     , c.email
                     , c.create_date
                     , GROUP_CONCAT(t.phone_number) AS phones
                  FROM Customers AS c
                INNER
                  JOIN Telephone AS t
                    ON t.customer_id= c.customer_id   
                GROUP
                    BY c.customer_id;
                `
                rows, err := db.Query(sqlStatement_Customer)  
                var count int
                count=0
                var  m = make(map[int]Customer)
                
                for rows.Next() {
                        err := rows.Scan(&customer_id, &full_name, &email, &create_date, &phones)
                        if err != nil {
                            log.Fatal("Error in retrieval of records",err)
                        }
                        //create a map to store values
                        m[count] = Customer {Customerid: customer_id, Name: full_name, Email: email, CreateDate: create_date, Phones: phones}
                        count++
                    }
                
                    defer db.Close()   
                    return m
}


func insertIntoDB(data myData)(int64) {
     
               
                db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db_telecomm")
                    if err != nil {
                        panic(err)
                    }

   
                // INSERT INTO THE CUSTOMER TABLE 
                sqlStatement_Customer := `
                INSERT INTO Customers (
                    full_name, 
                    email,
                    create_date)
                VALUES (?,?,curdate());
                `
            
             
                ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
                defer cancelfunc()
                stmt, err := db.PrepareContext(ctx, sqlStatement_Customer)
                    if err != nil {
                    log.Printf("Error Customers: %s when preparing SQL statement", err)
                    
                    }
                    defer stmt.Close()
                
                    res, err := stmt.ExecContext(ctx, data.name,data.email)  
                    if err != nil {  
                        fmt.Printf("Error Customer: %s when inserting row into customers table", err)
                         }
                    rows, err := res.RowsAffected()  
                    if err != nil {  
                        fmt.Printf("Error Customer: %s when finding rows affected", err)
                    }
                    log.Printf("%d telephone created ", rows)  
                     
                    id, err2 := res.LastInsertId()
                    if err2 != nil {  
                        fmt.Printf("Error Customer %s when retrieving id from Customer table", err2)
                    } else {
                        defer db.Close()   
                        fmt.Printf("Insertion completed with id %v !", id)

                            // new query 
                        sqlStatement_Telephone :=  `INSERT INTO Telephone(phone_number,customer_id)
                                                    VALUES (?,
                                                            (SELECT customer_id FROM Customers WHERE full_name = ? LIMIT 1));`

                        stmt, err := db.PrepareContext(ctx, sqlStatement_Telephone)
                        if err != nil {
                            log.Printf("Error Customers: %s when preparing SQL statement", err)
                            
                        }
                        defer stmt.Close()
                        
                        res, err := stmt.ExecContext(ctx, data.phonenumbers, data.name)  
                        if err != nil {  
                         fmt.Printf("Error Telephone: %s when inserting row into customers table", err)
                        }
                        rows, err := res.RowsAffected()  
                        if err != nil {  
                            fmt.Printf("Error Telephone: %s when finding rows affected", err)
                        }
                            
                        return id
                    }
                    
               
        defer db.Close()     
        fmt.Printf("Insertion into Customers Table completed !")
               
               
  return 0
             
}



func main() {
    var data myData;
    
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        req.ParseForm()
       
        for key,value := range req.Form{
        
            fmt.Printf("%s = %s\n", key,value)
            
            if(key == "email" ) {
                data.email = value[0]
            } else if (key == "name") {
                data.name = value[0]
            } else if (key == "phonenumbers") {
                data.phonenumbers = value[0]
            } else if (key == "insertion") {
                data.insertion = value[0];
            }
            fmt.Printf("%s\n %s\n %s\n", data.name,data.email)
              
        }

      
       if(data.insertion != "" && data.insertion == "true") {
        id := insertIntoDB(data)

            if(id != 0) {
               json.NewEncoder(w).Encode(true)
            } else {
                json.NewEncoder(w).Encode(false)
            }
       } else {

        
        results := retrieveAllRecords()
        fmt.Println(results)
        json.NewEncoder(w).Encode(results)
        
       }
          
    })
    /*Handle CORS*/
    handler := cors.Default().Handler(mux)
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
  
    
}