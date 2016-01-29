package smpp34

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

var (
	header = &Header{
		Id:       1,
		Status:   2,
		Sequence: 3,
	}
)

func BuildPduTest(t *testing.T) { TestingT(t) }

type BuildPduSuite struct{}

var _ = Suite(&BuildPduSuite{})

func (s *BuildPduSuite) TestBind(c *C) {
	b, _ := NewBind(
		&Header{Id: 1, Sequence: 1},
		[]byte{},
	)
	b.SetField(INTERFACE_VERSION, 0x34)
	b.SetField(SYSTEM_ID, "system_id")
	b.SetField(PASSWORD, "password")
	p := Pdu(b)

	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "0000002800000001000000000000000173797374656d5f69640070617373776f7264000034000000")
}

func (s *BuildPduSuite) TestBindResp(c *C) {
	b, _ := NewBindResp(
		header,
		[]byte{},
	)

	b.SetField(SYSTEM_ID, "system_id")
	b.SetTLVField(0x0210, 1, []byte{0x34}) // sc_interface_version TLV
	p := Pdu(b)
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "0000001f00000001000000020000000373797374656d5f6964000210000134")
}

func (s *BuildPduSuite) TestUnBind(c *C) {
	b, _ := NewUnbind(header)
	p := Pdu(b)
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "00000010000000060000000200000003")
}

func (s *BuildPduSuite) TestUnBindResp(c *C) {
	b, _ := NewUnbindResp(header)
	p := Pdu(b)
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "00000010800000060000000200000003")
}

func (s *BuildPduSuite) TestEnquireLink(c *C) {
	b, _ := NewEnquireLink(header)

	p := Pdu(b)
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "00000010000000150000000200000003")
}

func (s *BuildPduSuite) TestEnquireLinkResp(c *C) {
	b, _ := NewEnquireLinkResp(header)

	p := Pdu(b)
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "00000010800000150000000200000003")
}

func (s *BuildPduSuite) TestGenericNack(c *C) {
	b, _ := NewGenericNack(header)

	p := Pdu(b)
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "00000010800000000000000200000003")
}

func (s *BuildPduSuite) TestSubmitSM(c *C) {
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
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "00000051000000040000000000000001000000736f757263652d6164647265737300000064657374696e6174696f6e2d61646472657373000000000000000000000f636563692065737420756e20736d73")
}

func (s *BuildPduSuite) TestSubmitSMResp(c *C) {
	b, _ := NewSubmitSmResp(
		&Header{
			Id:       SUBMIT_SM,
			Sequence: 1,
		},
		[]byte{},
	)
	p := Pdu(b)
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "0000001180000004000000000000000100")
}

func (s *BuildPduSuite) TestDeliverSM(c *C) {
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
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "00000051000000050000000000000001000000736f757263652d6164647265737300000064657374696e6174696f6e2d61646472657373000000000000000000000f636563692065737420756e20736d73")
}

func (s *BuildPduSuite) TestDeliverSMResp(c *C) {
	b, _ := NewDeliverSmResp(
		&Header{
			Id:       SUBMIT_SM,
			Sequence: 1,
		},
		[]byte{},
	)
	p := Pdu(b)
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "0000001180000005000000000000000100")
}
func (s *BuildPduSuite) TestQuerySM(c *C) {
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
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "00000022000000040000000000000001000000736f757263652d6164647265737300")
}

func (s *BuildPduSuite) TestQuerySMResp(c *C) {
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
	c.Assert(fmt.Sprintf("%x", p.Writer()), Equals, "0000001400000004000000000000000100000000")
}
