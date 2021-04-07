package err

import "clipper_server/models/resp"

var ParamsError = &resp.Response{
    Code:    100501,
    Message: "参数错误",
    Data:    nil,
}

var MysqlError = &resp.Response{
    Code:    100601,
    Message: "MySQL错误",
    Data:    nil,
}

var NoRouteError = &resp.Response{
    Code:    100404,
    Message: "没有找到路由",
    Data:    nil,
}
