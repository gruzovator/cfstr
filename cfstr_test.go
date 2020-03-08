package cfstr_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gruzovator/cfstr"
)

func ExampleUpperCase() {
	type MyStruct struct {
		Label cfstr.UpperCase
	}

	a := MyStruct{Label: cfstr.NewUpperCase("Test")}
	fmt.Printf("a.Label=%s\n", a.Label)

	jsonData, err := json.Marshal(a)
	if err != nil {
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

func TestUpperCase_Equality(t *testing.T) {
	s1 := cfstr.NewUpperCase("СтрОка")
	s2 := cfstr.NewUpperCase("строкА")
	if s1 != s2 {
		t.Errorf("%s is not equal to %s", s1, s2)
	}
}

func TestUpperCase_String(t *testing.T) {
	arg := "Test Строка"
	want := "TEST СТРОКА"
	s := cfstr.NewUpperCase(arg)
	if s.String() != want {
		t.Errorf("want: %s, got: %s", want, s.String())
	}
}

func TestUpperCase_MarshalJSON(t *testing.T) {
	arg := "tEsT"
	want := "\"TEST\""

	s := cfstr.NewUpperCase(arg)
	bb, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("json marshal error: %s", err)
	}
	if string(bb) != want {
		t.Errorf("want: %s, got: %s", want, string(bb))
	}
}

func TestUpperCase_UnmarshalJSON(t *testing.T) {
	jsonData := []byte("\"tEsT\"")
	want := "TEST"

	var s cfstr.UpperCase
	if err := json.Unmarshal(jsonData, &s); err != nil {
		t.Fatalf("json ummarshal error: %s", err)
	}
	if s.String() != want {
		t.Errorf("want: %s, got: %s", want, s.String())
	}
}

func TestLowerCase_Equality(t *testing.T) {
	s1 := cfstr.NewLowerCase("СтрОка")
	s2 := cfstr.NewLowerCase("строкА")
	if s1 != s2 {
		t.Errorf("%s is not equal to %s", s1, s2)
	}
}

func TestLowerCase_String(t *testing.T) {
	arg := "Test Строка"
	want := "test строка"
	s := cfstr.NewLowerCase(arg)
	if s.String() != want {
		t.Errorf("want: %s, got: %s", want, s.String())
	}
}

func TestLowerCase_MarshalJSON(t *testing.T) {
	arg := "TeSt"
	want := "\"test\""

	s := cfstr.NewLowerCase(arg)
	bb, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("json marshal error: %s", err)
	}
	if string(bb) != want {
		t.Errorf("want: %s, got: %s", want, string(bb))
	}
}

func TestLowerCase_UnmarshalJSON(t *testing.T) {
	jsonData := []byte("\"tEsT\"")
	want := "test"

	var s cfstr.LowerCase
	if err := json.Unmarshal(jsonData, &s); err != nil {
		t.Fatalf("json ummarshal error: %s", err)
	}
	if s.String() != want {
		t.Errorf("want: %s, got: %s", want, s.String())
	}
}
