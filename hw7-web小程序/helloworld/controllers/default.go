package controllers

import (
	"github.com/astaxie/beego"
	//"encoding/json"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
    tmp := c.GetString("name")
    age := c.GetString("age")
    out := "Hello!" + tmp + "\n" + "your age is " + age + "!\n";
    c.Ctx.WriteString(out)  

}

//通过子struct的方法来重写post函数，实现自己的逻辑。
func (c *MainController) Post() {

    // var req map[string]interface{}
    // var res map[string]interface{}

    //请求信息
    // bufReq := c.Ctx.Input.RequestBody
    // c.Ctx.WriteString(string(bufReq))
    // bufReq := c.GetString("name")
    // json.Unmarshal(bufReq, &req)
    // //返回响应结果
    // lenReq := len(bufReq)
    // if lenReq > 1024 {
    //     beego.Info("req:", "\n"+string(bufReq[:64]), "...", string(bufReq[lenReq-64:]), "\n")
    // } else {
    //     beego.Info("req:", "\n"+string(bufReq), "\n")
    // }

    //模拟处理客户端发来的请求消息
    // c.HandleReq(bufReq)

    // bufRes, _ := json.MarshalIndent(res, "", "  ")
    // lenRes := len(bufRes)
    // if lenRes > 1024 {
    //     beego.Info("res:", "\n"+string(bufRes[:32]), "...", string(bufRes[lenRes-64:]), "\n")
    // } else {
    //     beego.Info("res:", "\n"+string(bufRes), "\n")
    // }

    // c.Data["json"] = res
    // c.ServeJson(false)
}

//-- 处理客户端发来的请求消息
//-- 参数：json数据，
// func (c *MainController) HandleReq(jsonData []byte) map[string]interface{} {

//     var res map[string]interface{}

//     json.Unmarshal(jsonData, &res)
//     res["ver"] = 2
//     res["cmd"] = "sys.pong"

//     return res
// }

func (c *MainController) HandleReq(jsonData string){
	c.Ctx.WriteString("Welcome!");
	c.Ctx.WriteString(jsonData);
}
