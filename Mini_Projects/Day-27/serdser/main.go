package main

import (

    "encoding/json"
    "fmt"
	"bytes"
	"encoding/gob"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    // Create an instance of the Person struct and Marshalling Go Data 
	p1 := Person{Name: "John Doe", Age: 30}

  jsonData, err := json.Marshal(p1)
  if err != nil {
	fmt.Println("ERROR",err)
	return
  }
  fmt.Println("Serialized Json data : ",string(jsonData))


  // UnMarshalling Json Data

  var p2 Person
  jsonData2 := []byte(`{"Name": "John Doe", "Age": 30}`)
  err = json.Unmarshal(jsonData2,&p2)
  if err != nil {
	fmt.Println("ERROR",err)
	return
  }
  fmt.Println("Deserialized Json Data",p2)

  
    // Create a new buffer to write the serialized data to
    var b bytes.Buffer

    // Create a new gob encoder and use it to encode the person struct
    enc := gob.NewEncoder(&b)
    if err := enc.Encode(p1); err != nil {
        fmt.Println("Error encoding struct:", err)
        return
    }

    // The serialized data can now be found in the buffer
    serializedData := b.Bytes()
    fmt.Println("Serialized data:", serializedData)

	    // Create a new buffer from the serialized data
		b1 := bytes.NewBuffer(serializedData)

		// Create a new gob decoder and use it to decode the person struct
		var person Person
		dec := gob.NewDecoder(b1)
		if err := dec.Decode(&person); err != nil {
			fmt.Println("Error decoding struct:", err)
			return
		}
	
		// The person struct has now been deserialized
		fmt.Println("Deserialized struct:", person)
}