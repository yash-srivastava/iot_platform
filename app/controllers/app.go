package controllers

import (
	"github.com/StabbyCutyou/buffstreams"
	"github.com/iot_platform/conf"
	"github.com/iot_platform/lib/sender"
	"github.com/iot_platform/lib/utils"
	"github.com/revel/revel"
	"net/http"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Send_3000() revel.Result {

	response := sender.Response{Success: false, Message: "Something went wrong"}
	packet := sender.Packet_3000{}
	err := c.Params.BindJSON(&packet)
	if err == nil {
		response = sender.HandlePacket(0x3000, packet)
	}
	return c.RenderJSON(response)
}

func (c App) Send_8000() revel.Result {

	response := sender.Response{Success: false, Message: "Something went wrong"}
	packet := sender.Packet_8000{}
	err := c.Params.BindJSON(&packet)
	if err == nil {
		response = sender.HandlePacket(0x8000, packet)
	}
	return c.RenderJSON(response)
}

func (c App) IsSguConnected() revel.Result {
	response := sender.Response{Success: false, Message: "Something went wrong"}

	params := make(map[string]uint64)
	err := c.Params.BindJSON(&params)
	if err == nil {
		response.Data = conf.SGU_TCP_CONNECTION.Has(utils.ToStr(params["sguid"]))
		response.Success = true
		response.Message = ""
	}
	return c.RenderJSON(response)
}

func (c App) GetConnectedScus() revel.Result {
	response := sender.Response{Success: false, Message: "Something went wrong"}

	params := make(map[string]uint64)
	err := c.Params.BindJSON(&params)
	if err == nil {

		response.Data, _ = conf.SGU_SCU_LIST.Get(utils.ToStr(params["sguid"]))
		response.Success = true
		response.Message = ""
	}
	return c.RenderJSON(response)
}

func (c App) GetConnectedSgus() revel.Result {
	data := conf.Sgu{}
	var sgus []uint64
	if buffstreams.TcpSguMap != nil {
		for _, v := range buffstreams.TcpSguMap.Items() {
			if v != nil {
				sgus = append(sgus, utils.ToUint64(v))
			}
		}
	}

	data.SguIds = sgus
	response := sender.Response{}
	response.Data = data
	response.Success = true

	return c.RenderJSON(response)
}
type MyHtml string

func (r MyHtml) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusAccepted, "text/html")
}

func (c App) Get() revel.Result {

	revel.INFO.Println(c.Params)
	/*response := sender.Response{Success: true, Message: "Something went wrong"}
	packet := make(map[string]interface{})
	err := c.Params.BindJSON(&packet)
	if err == nil {
		revel.INFO.Println(packet)
	}*/
	c.Response.Status=202
	return MyHtml("")
}
