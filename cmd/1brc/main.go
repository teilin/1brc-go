package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

type Measure struct 
  Min float64
  Max float64
  Sum float64
  Count int
}

func main() {
  file, err := os.Open("../1brc/measurements2.txt")
  if err != nil {
    fmt.Println("Error reading file")
  }
  defer file.Close()

  mem := make(map[string]Measure)

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    entry := scanner.Text()
    splits := strings.Split(entry, ";")
    agg, ok := mem[splits[0]]
    temp := 0.0
    if s, err := strconv.ParseFloat(splits[1], 64); err == nil {
      temp = s
    }
    if ok {
      if agg.Min > temp {
        agg.Min = temp
      }
      if agg.Max < temp {
        agg.Max = temp
      }
      agg.Sum += temp
      agg.Count += 1
      mem[splits[0]] = agg
    } else {
      mem[splits[0]] = Measure{
        Min: temp,
        Max: temp,
        Sum: temp,
        Count: 1,
      }
    }
  }

  for city, m := range mem {
    fmt.Println(fmt.Sprintf("%s : %f %f %f", city, m.Min, m.Max, m.Sum / float64(m.Count)))
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Failed")
  }
}
