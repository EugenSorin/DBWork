package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
    var nrC int
    fmt.Println("Program de transfer a datelor din SIMEC, realizat în GO")
    db, err := sql.Open("mysql", "eugen:quit;#427@(192.168.1.9)/Evidenta")
    if err != nil {
        fmt.Printf("Eroare de conectare: %v", err); fmt.Println()
    } else {
        fmt.Println("Spor la treabă!...")
    }
    err = db.QueryRow("SELECT COUNT(*) FROM beneficiari WHERE Renuntat = 'DA' AND IsGoodCNP(CodFiscal)").Scan(&nrC)
    switch {
    case err == sql.ErrNoRows:
        fmt.Println("Nu există clienți")
    case err != nil:
        fmt.Printf("Eroare de execuție: %v", err); fmt.Println()
    default:
        fmt.Println("Număr de clienți renunțați:", nrC)
    }
    err = db.QueryRow("SELECT COUNT(*) FROM Contracte.facturi").Scan(&nrC)
    switch {
    case err == sql.ErrNoRows:
        fmt.Println("Nu există clienți")
    case err != nil:
        fmt.Printf("Eroare de execuție: %v", err); fmt.Println()
    default:
        fmt.Println("Număr de facturi emise:", nrC)
    }
}
