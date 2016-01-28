package smpp34

import (
	"fmt"
	"strings"
	"testing"
)

const (
	exampleSubmitSm = "00000051000000040000000000000001000000736f757263652d6164647265737300000064657374696e6174696f6e2d61646472657373000000000000000000000f636563692065737420756e20736d73"
)

var (
	header = &Header{
		Id:       1,
		Status:   2,
		Sequence: 3,
	}
)

func TestBind(t *testing.T) {
	b, _ := NewBind(
		&Header{Id: 1, Sequence: 1},
		[]byte{},
	)
	b.SetField(INTERFACE_VERSION, 0x34)
	b.SetField(SYSTEM_ID, "system_id")
	b.SetField(PASSWORD, "password")
	p := Pdu(b)
	assertPdu(t, p, "0000002800000001000000000000000173797374656d5f69640070617373776f7264000034000000", "Bind")
}

func assertPdu(t *testing.T, p Pdu, expected string, name string) {
	actual := fmt.Sprintf("%x", p.Writer())
	if strings.Compare(actual, expected) != 0 {
		t.Fatalf("%s should be %s but was %s", name, expected, actual)
	}
}

func TestBindResp(t *testing.T) {
	b, _ := NewBindResp(
		header,
		[]byte{},
	)

	b.SetField(SYSTEM_ID, "system_id")
	b.SetTLVField(0x0210, 1, []byte{0x34}) // sc_interface_version TLV
	p := Pdu(b)
	assertPdu(t, p, "0000001f00000001000000020000000373797374656d5f6964000210000134", "BindResp")
}

func TestUnBind(t *testing.T) {
	b, _ := NewUnbind(header)
	p := Pdu(b)
	assertPdu(t, p, "00000010000000060000000200000003", "Unbind")
}

func TestUnBindResp(t *testing.T) {
	b, _ := NewUnbindResp(header)
	p := Pdu(b)
	assertPdu(t, p, "00000010800000060000000200000003", "UnbindResp")
}

func TestEnquireLink(t *testing.T) {
	b, _ := NewEnquireLink(header)

	p := Pdu(b)
	assertPdu(t, p, "00000010000000150000000200000003", "EnquireLink")
}

func TestEnquireLinkResp(t *testing.T) {
	b, _ := NewEnquireLinkResp(header)

	p := Pdu(b)
	assertPdu(t, p, "00000010800000150000000200000003", "EnquireLinkResp")
}

func TestGenericNack(t *testing.T) {
	b, _ := NewGenericNack(header)

	p := Pdu(b)
	assertPdu(t, p, "00000010800000000000000200000003", "GenericNack")
}

func TestSubmitSM(t *testing.T) {
	b, _ := NewSubmitSm(
		&Header{
			Id:       SUBMIT_SM,
			Sequence: 1,
		},
		[]byte{},
	)
	b.SetField(SOURCE_ADDR, "source-address")
	b.SetField(DESTINATION_ADDR, "destination-address")
	b.SetField(SHORT_MESSAGE, "ceci est un sms")
	p := Pdu(b)
	assertPdu(t, p, "00000051000000040000000000000001000000736f757263652d6164647265737300000064657374696e6174696f6e2d61646472657373000000000000000000000f636563692065737420756e20736d73", "SubmitSm")
}

func TestSubmitSMResp(t *testing.T) {
	b, _ := NewSubmitSmResp(
		&Header{
			Id:       SUBMIT_SM,
			Sequence: 1,
		},
		[]byte{},
	)
	p := Pdu(b)
	assertPdu(t, p, "0000001180000004000000000000000100", "SubmitSmResp")
}

func TestDeliverSM(t *testing.T) {
	b, _ := NewDeliverSm(
		&Header{
			Id:       SUBMIT_SM,
			Sequence: 1,
		},
		[]byte{},
	)
	b.SetField(SOURCE_ADDR, "source-address")
	b.SetField(DESTINATION_ADDR, "destination-address")
	b.SetField(SHORT_MESSAGE, "ceci est un sms")
	p := Pdu(b)
	assertPdu(t, p, "00000051000000050000000000000001000000736f757263652d6164647265737300000064657374696e6174696f6e2d61646472657373000000000000000000000f636563692065737420756e20736d73", "DeliverSm")
}

func TestDeliverSMResp(t *testing.T) {
	b, _ := NewDeliverSmResp(
		&Header{
			Id:       SUBMIT_SM,
			Sequence: 1,
		},
		[]byte{},
	)
	p := Pdu(b)
	assertPdu(t, p, "0000001180000005000000000000000100", "DeliverSmResp")
}
func TestQuerySM(t *testing.T) {
	b, _ := NewQuerySm(
		&Header{
			Id:       SUBMIT_SM,
			Sequence: 1,
		},
		[]byte{},
	)
	b.SetField(SOURCE_ADDR, "source-address")
	b.SetField(DESTINATION_ADDR, "destination-address")
	b.SetField(SHORT_MESSAGE, "ceci est un sms")
	p := Pdu(b)
	assertPdu(t, p, "00000022000000040000000000000001000000736f757263652d6164647265737300", "QuerySm")
}

func TestQuerySMResp(t *testing.T) {
	b, _ := NewQuerySmResp(
		&Header{
			Id:       SUBMIT_SM,
			Sequence: 1,
		},
		[]byte{},
	)
	b.SetField(SOURCE_ADDR, "source-address")
	b.SetField(DESTINATION_ADDR, "destination-address")
	b.SetField(SHORT_MESSAGE, "ceci est un sms")
	p := Pdu(b)
	assertPdu(t, p, "0000001400000004000000000000000100000000", "QuerySmResp")
}
