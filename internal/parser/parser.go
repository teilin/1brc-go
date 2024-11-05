package parser

import(

)

const (
  NEW_LINE = byte('\n')
  VALUE_SEPARATOR = byte(';')
)

type Parser struct{
  input []byte
}

func (b Parser) Parse() (map[string]string, error) {
  output := make(map[string]string)
  return output, nil
}

func splitByLine(bytes *[]byte) {
  startPos := 0
  endPos := 0
  for _, b := range *bytes{
    if b != NEW_LINE {
      endPos += 1
      continue
    }
    handleLine(bytes[startPos:endendPos])
    startPos = endPos + 1
    endPos = 0
  }
}

func handleLine(input []byte) {}
