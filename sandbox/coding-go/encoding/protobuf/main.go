package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	pbExample "github.com/butuzov/sandbox/protobuf/example"
	pbValid "github.com/butuzov/sandbox/protobuf/validation"

	"github.com/davecgh/go-spew/spew"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
)

type Any struct {
	message string
}

func main() {
	example := pbExample.Example{

		// 1: Nested (and repeated) message
		Messages: []*pbExample.Example_Message{
			{Text: "Привіт Світ!"},
			{Text: "¡Hola, mundo!"},
			{Text: "こんにちは世界"},
		},

		// 2: Using Standalone Enums
		Color: pbExample.Colors_RED,

		// 3: Using Nested Enums
		Shape: pbExample.Example_CUBE,

		// 4... skiping some reserved.

		// 6: Referenced other subtype
		Other: &pbExample.Main_Derivative{
			Value: "hej...",
		},

		// Any: nil,
		Localization: map[string]string{
			"de_DE": "Deutschland",
		},
	}

	//== Any ================================================================

	// Any Type marshaling an d unmarshaling isn't supported by api v2
	// var (
	// 	a1 = pb.Variation{Text: "hi"}
	// 	a2 = pb.Alternative{Text: "hi"}
	// )

	// if any, err := ptypes.MarshalAny(a1); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	example.Any = append(example.Any, any)
	// }
	// if any, err := ptypes.MarshalAny(a2); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	example.Any = append(example.Any, any)
	// }

	fmt.Fprintln(os.Stdout, "Proto Message")
	fmt.Fprintln(os.Stdout, strings.Repeat("=", 60))
	spew.Fdump(os.Stdout, example)
	fmt.Fprintln(os.Stdout, strings.Repeat("^", 60))
	fmt.Println()

	b, err := proto.Marshal(&example)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, "Binary Representation")
	fmt.Fprintln(os.Stdout, strings.Repeat("=", 60))
	spew.Fdump(os.Stdout, b)
	fmt.Fprintln(os.Stdout, strings.Repeat("^", 60))
	fmt.Println()

	sample := pbExample.Example{}
	fmt.Fprintln(os.Stdout, "Proto Message (Empty)")
	fmt.Fprintln(os.Stdout, strings.Repeat("=", 60))
	spew.Fdump(os.Stdout, sample)
	fmt.Fprintln(os.Stdout, strings.Repeat("^", 60))
	fmt.Println()

	if err := proto.Unmarshal(b, &sample); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, "Proto Message (Unmarshaled)")
	fmt.Fprintln(os.Stdout, strings.Repeat("=", 60))
	spew.Fdump(os.Stdout, sample)
	fmt.Fprintf(os.Stdout, "\n%#v\n", sample)
	fmt.Fprintln(os.Stdout, strings.Repeat("^", 60))
	fmt.Println()

	//== OneOf ==============================================================

	fmt.Fprintln(os.Stdout, "Oneof test")
	fmt.Fprintf(os.Stdout, "Header %s:\n", example.GetHeader())
	fmt.Fprintf(os.Stdout, "SubHeader %s:\n", example.GetSubheader())

	example.Oneof = &pbExample.Example_Header{"Hey!"}

	fmt.Fprintln(os.Stdout, "Setting Header")
	fmt.Fprintf(os.Stdout, "Header %s:\n", example.GetHeader())
	fmt.Fprintf(os.Stdout, "SubHeader %s:\n", example.GetSubheader())

	example.Oneof = &pbExample.Example_Subheader{"Hey!"}
	fmt.Fprintln(os.Stdout, "Setting SubHeader")
	fmt.Fprintf(os.Stdout, "Header %s:\n", example.GetHeader())
	fmt.Fprintf(os.Stdout, "SubHeader %s:\n", example.GetSubheader())

	fmt.Fprintln(os.Stdout, strings.Repeat("^", 60))
	fmt.Println()

	//=== String ============================================================

	fmt.Fprintln(os.Stdout, "Proto Message (Stringer)")
	fmt.Fprintln(os.Stdout, strings.Repeat("=", 60))
	spew.Fdump(os.Stdout, example.String())
	fmt.Fprintln(os.Stdout, strings.Repeat("^", 60))
	fmt.Println()

	//=== Optional =======================================================

	val := "string"
	spew.Fdump(os.Stdout, example.String())
	example.OptString = &val
	spew.Fdump(os.Stdout, example.String())

	//=== jsonp ===============================================================
	fmt.Fprintln(os.Stdout, "Proto Message")
	fmt.Fprintln(os.Stdout, strings.Repeat("=", 60))

	m := jsonpb.Marshaler{
		EnumsAsInts:  false,
		Indent:       "\t",
		OrigName:     true,
		EmitDefaults: true,
	}

	if err := m.Marshal(os.Stdout, &example); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Fprintln(os.Stdout, strings.Repeat("^", 60))
	fmt.Println()

	//=== Reset ===============================================================

	example.Reset()
	fmt.Fprintln(os.Stdout, "Proto Message Reset")
	fmt.Fprintln(os.Stdout, strings.Repeat("=", 60))
	spew.Fdump(os.Stdout, example)
	spew.Fdump(os.Stdout, example.String())
	fmt.Fprintf(os.Stdout, "%#v", sample)
	fmt.Fprintln(os.Stdout, strings.Repeat("^", 60))
	fmt.Println()

	//

	validation := &pbValid.ValidationExample{}
	validation.Id = 1
	fmt.Println(validation.Validate())

	validation.Id = 1000
	fmt.Println(validation.Validate())
}
