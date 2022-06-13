package response

//import (
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//)
//
//type baseRsp struct {
//	statusCode int          // 返回的状态码
//	ctx        *gin.Context // 记录上下文信息
//}
//
//type Rsp struct {
//	BaseStatus *baseRsp    `json:"-"`
//	Status     int         `json:"status"` // 0代表成功,-1代表失败
//	Msg        string      `json:"msg"`    // 失败时携带错误信息
//	Data       interface{} `json:"data"`   // 成功且动作为查询时返回数据
//}
//
//type QueryResult struct {
//	Items interface{} `json:"items"` // 实际数据
//	Total int64       `json:"total"` // 查询时用于标注为数据总量,也可复用其它用途
//}
//
//func NewRsp(ctx *gin.Context) *Rsp {
//	return &Rsp{
//		BaseStatus: &baseRsp{
//			statusCode: http.StatusOK,
//			ctx:        ctx,
//		},
//	}
//}
//
//func (r *Rsp) SetStatusCode(statusCode int) {
//	r.BaseStatus.statusCode = statusCode
//}
//
//func (r *Rsp) SetData(data interface{}) {
//	r.Data = data
//}
//
//func (r *Rsp) SetError(err error) {
//	r.Status = -1
//	r.Msg = err.Error()
//}
//
//func (r *Rsp) RspSuccess(data interface{}) {
//	r.SetData(data)
//	r.BaseStatus.ctx.JSON(r.BaseStatus.statusCode, r)
//}
//
//func (r *Rsp) RspError(statusCode int, err error) {
//	r.SetStatusCode(statusCode)
//	r.SetError(err)
//	r.BaseStatus.ctx.Abort()
//	r.BaseStatus.ctx.JSON(r.BaseStatus.statusCode, r)
//}
