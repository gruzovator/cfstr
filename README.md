# Case folded string types

String based types that force string case (uppercase or lowercase).
Types implement Stringer, json.Marshaler and json.Unmarshaler interfaces.

Type instances can be created:
* from string with NewUpperCase or NewLowerCase function 
* from json object bytes 

## Example 1.

Force upper case for _MyStruct_ _Label_ field

```go
func ExampleUpperCase() {
	type MyStruct struct {
		Label cfstr.UpperCase
	}

	a := MyStruct{Label: cfstr.NewUpperCase("Test")}
	fmt.Printf("a.Label=%s\n", a.Label)

	jsonData, err:=json.Marshal(a)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("json for a: %s\n", string(jsonData))

	var b MyStruct
	if err := json.Unmarshal([]byte(`{"Label": "Test"}`), &b); err != nil {
		panic(err)
	}
	fmt.Printf("b.Label=%s\n", b.Label)

	// Output:
	// a.Label=TEST
	// json for a: {"Label":"TEST"}
	// b.Label=TEST
}
```


## Example 2.

Force lower case and define distinct type for _Residency_: 

```go
func ExampleLowerCase() {
	// define destinct type based on cfstr.LowerCase
	type Residency struct {
		cfstr.LowerCase
	}
	NewResidency := func(s string) Residency {
		return Residency{cfstr.NewLowerCase(s)}
	}

	r1 := NewResidency("EN")
	fmt.Printf("r1=%s\n", r1)

	jsonData, err := json.Marshal(r1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("r1 json: %s\n", string(jsonData))

	var r2 Residency
	if err := json.Unmarshal([]byte(`"RU"`), &r2); err != nil {
		panic(err)
	}
	fmt.Printf("r2=%s\n", r2)

	// Output:
	// r1=en
	// r1 json: "en"
	// r2=ru
}
```