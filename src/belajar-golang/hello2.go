package main

import "fmt"

func main()  {
  // Seleksi Kondisi if, else if, else
  var point = 10

  if point == 10 {
    fmt.Println("Lulus dengan nilai Sempurna !")
  } else if point > 5 {
    fmt.Println("lulus")
  } else if point == 4 {
    fmt.Println("Hampir Lulus")
  } else {
    fmt.Println("Tidak lulus, nilai anda %d\n", point)
  }


  // Perulangan for
  for i := 2; i < 10; i++ {
    fmt.Println("Angka", i)
  }

  // Perulangan for dengan argumen hanya Kondisi
  var i = 0
  for i < 7 {
    fmt.Println("Angka", i)
    i++
  }

  // Perulangan for tanpa argumen
  var b = 0

  for {
    fmt.Println("Angka tanpa argumen", b)

    b++
    if b == 5 {
      break
    }
  }

  // break dan continue
  for i := 1; i <= 10; i++ {
    if i % 2 == 1 {
      continue
    }

    if i > 10 {
      break
    }

    fmt.Println("AngkaLooping Genap", i)
  }

  // Perulangan bersarang
  for i := 0; i < 5; i++ {
    for j := i; j < 5; j++ {
        fmt.Print(j, " ")
    }

    fmt.Println()
  }

  // Perulangan laber pada break dan continue
  outerLoop:
  for i := 0; i < 5; i++ {
      for j := 0; j < 5; j++ {
          if i == 3 {
              break outerLoop
          }
          fmt.Print("matriks [", i, "][", j, "]", "\n")
      }
  }

}
