// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     blobschema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type RedPanda struct {
	Date string `json:"Date"`

	Open float64 `json:"Open"`

	Close float64 `json:"Close"`

	High float64 `json:"High"`

	Low float64 `json:"Low"`

	Volume int32 `json:"Volume"`

	Dividends float64 `json:"Dividends"`

	StockSplits float64 `json:"StockSplits"`

	CapitalGains float64 `json:"CapitalGains"`
}

const RedPandaAvroCRC64Fingerprint = "\x1f0\xb3\x1c\x90\xb8\x01\xba"

func NewRedPanda() RedPanda {
	r := RedPanda{}
	r.Close = 0
	r.High = 0
	r.Low = 0
	r.Volume = 0
	r.Dividends = 0
	r.StockSplits = 0
	r.CapitalGains = 0
	return r
}

func DeserializeRedPanda(r io.Reader) (RedPanda, error) {
	t := NewRedPanda()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRedPandaFromSchema(r io.Reader, schema string) (RedPanda, error) {
	t := NewRedPanda()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRedPanda(r RedPanda, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Date, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Open, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Close, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.High, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Low, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Volume, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Dividends, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.StockSplits, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.CapitalGains, w)
	if err != nil {
		return err
	}
	return err
}

func (r RedPanda) Serialize(w io.Writer) error {
	return writeRedPanda(r, w)
}

func (r RedPanda) Schema() string {
	return "{\"fields\":[{\"name\":\"Date\",\"type\":\"string\"},{\"name\":\"Open\",\"type\":\"double\"},{\"default\":0,\"name\":\"Close\",\"type\":\"double\"},{\"default\":0,\"name\":\"High\",\"type\":\"double\"},{\"default\":0,\"name\":\"Low\",\"type\":\"double\"},{\"default\":0,\"name\":\"Volume\",\"type\":\"int\"},{\"default\":0,\"name\":\"Dividends\",\"type\":\"double\"},{\"default\":0,\"name\":\"StockSplits\",\"type\":\"double\"},{\"default\":0,\"name\":\"CapitalGains\",\"type\":\"double\"}],\"name\":\"RedPanda\",\"type\":\"record\"}"
}

func (r RedPanda) SchemaName() string {
	return "RedPanda"
}

func (_ RedPanda) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RedPanda) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RedPanda) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RedPanda) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RedPanda) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RedPanda) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RedPanda) SetString(v string)   { panic("Unsupported operation") }
func (_ RedPanda) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RedPanda) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.Date}

		return w

	case 1:
		w := types.Double{Target: &r.Open}

		return w

	case 2:
		w := types.Double{Target: &r.Close}

		return w

	case 3:
		w := types.Double{Target: &r.High}

		return w

	case 4:
		w := types.Double{Target: &r.Low}

		return w

	case 5:
		w := types.Int{Target: &r.Volume}

		return w

	case 6:
		w := types.Double{Target: &r.Dividends}

		return w

	case 7:
		w := types.Double{Target: &r.StockSplits}

		return w

	case 8:
		w := types.Double{Target: &r.CapitalGains}

		return w

	}
	panic("Unknown field index")
}

func (r *RedPanda) SetDefault(i int) {
	switch i {
	case 2:
		r.Close = 0
		return
	case 3:
		r.High = 0
		return
	case 4:
		r.Low = 0
		return
	case 5:
		r.Volume = 0
		return
	case 6:
		r.Dividends = 0
		return
	case 7:
		r.StockSplits = 0
		return
	case 8:
		r.CapitalGains = 0
		return
	}
	panic("Unknown field index")
}

func (r *RedPanda) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ RedPanda) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RedPanda) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RedPanda) HintSize(int)                     { panic("Unsupported operation") }
func (_ RedPanda) Finalize()                        {}

func (_ RedPanda) AvroCRC64Fingerprint() []byte {
	return []byte(RedPandaAvroCRC64Fingerprint)
}

func (r RedPanda) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Date"], err = json.Marshal(r.Date)
	if err != nil {
		return nil, err
	}
	output["Open"], err = json.Marshal(r.Open)
	if err != nil {
		return nil, err
	}
	output["Close"], err = json.Marshal(r.Close)
	if err != nil {
		return nil, err
	}
	output["High"], err = json.Marshal(r.High)
	if err != nil {
		return nil, err
	}
	output["Low"], err = json.Marshal(r.Low)
	if err != nil {
		return nil, err
	}
	output["Volume"], err = json.Marshal(r.Volume)
	if err != nil {
		return nil, err
	}
	output["Dividends"], err = json.Marshal(r.Dividends)
	if err != nil {
		return nil, err
	}
	output["StockSplits"], err = json.Marshal(r.StockSplits)
	if err != nil {
		return nil, err
	}
	output["CapitalGains"], err = json.Marshal(r.CapitalGains)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RedPanda) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Date"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Date); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Date")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Open"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Open); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Open")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Close"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Close); err != nil {
			return err
		}
	} else {
		r.Close = 0
	}
	val = func() json.RawMessage {
		if v, ok := fields["High"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.High); err != nil {
			return err
		}
	} else {
		r.High = 0
	}
	val = func() json.RawMessage {
		if v, ok := fields["Low"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Low); err != nil {
			return err
		}
	} else {
		r.Low = 0
	}
	val = func() json.RawMessage {
		if v, ok := fields["Volume"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Volume); err != nil {
			return err
		}
	} else {
		r.Volume = 0
	}
	val = func() json.RawMessage {
		if v, ok := fields["Dividends"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Dividends); err != nil {
			return err
		}
	} else {
		r.Dividends = 0
	}
	val = func() json.RawMessage {
		if v, ok := fields["StockSplits"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StockSplits); err != nil {
			return err
		}
	} else {
		r.StockSplits = 0
	}
	val = func() json.RawMessage {
		if v, ok := fields["CapitalGains"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.CapitalGains); err != nil {
			return err
		}
	} else {
		r.CapitalGains = 0
	}
	return nil
}
