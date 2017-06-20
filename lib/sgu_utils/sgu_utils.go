package sgu_utils

import (
	"github.com/StabbyCutyou/buffstreams"

	"github.com/iot_platform/lib/parser"
	"github.com/revel/revel"
	"github.com/iot_platform/lib/utils"
)


func ParseInputPackets(conn *buffstreams.Client)  {
	string_data := utils.ConvertBytesToString(conn.Data)
	revel.INFO.Println(string_data)
	parser.Wrap(conn)
}

