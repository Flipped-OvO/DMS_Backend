package request

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Checker interface {
	Check() error
}

//func GetClaims(ctx *gin.Context) (claims *auth.CustomClaims, err error) {
//	claims = auth.GetClaims(ctx)
//	if claims == nil {
//		err = errors.New("token invalid")
//	}
//	return
//}

func ParseParamFail(ctx *gin.Context, r interface{}, fs ...func() error) bool {
	if r != nil {
		err := ctx.ShouldBind(r)
		if err != nil {
			err = HTTPError{Status: http.StatusBadRequest, Msg: err.Error()}

			return Fail(ctx, err)
		}

		if checker, ok := r.(Checker); ok {
			err := checker.Check()
			if err != nil {
				err = HTTPError{Status: http.StatusBadRequest, Msg: err.Error()}

				return Fail(ctx, err)
			}
		}
	}

	for _, f := range fs {
		err := f()
		if err != nil {
			err = HTTPError{Status: http.StatusBadRequest, Msg: err.Error()}

			return Fail(ctx, err)
		}
	}

	return false
}

func Fail(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	log.Println(err.Error())
	status := http.StatusInternalServerError
	if v, ok := err.(HTTPError); ok {
		status = v.Status
	}

	ctx.JSON(200, gin.H{"code": status, "msg": err.Error()})

	return true
}

func Success(ctx *gin.Context, data interface{}, extra ...interface{}) {
	if data == nil {
		ctx.Status(http.StatusNoContent)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"data":  data,
		"extra": extra,
	})
}

func ErrorAuth(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusForbidden)
}

func ParseQueryInt(ctx *gin.Context, key string) (number int64, err error) {
	return strconv.ParseInt(ctx.DefaultQuery(key, "-1"), 10, 64)
}

func ParseFormInt(ctx *gin.Context, key string) (number int64, err error) {
	return strconv.ParseInt(ctx.DefaultPostForm(key, ""), 10, 64)
}
