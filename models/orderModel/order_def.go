//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/12/2 2:11 下午
//
//============================================================

package orderModel

import "github.com/pssauron/gocore/libs"

type SEOrder struct {
	OrderID   libs.Int    `json:"orderId" db:"OrderID" primary:"true"` //订单主键
	OrderDesc libs.String `json:"orderDesc" db:"OrderDesc"`            //订单描述
	OrderNum  libs.String `json:"orderNum" db:"OrderNum"`              //订单编码
	CustNum   libs.String `json:"custNum" db:"CustNum"`                //客户编码
	OrderDate libs.String `json:"orderDate" db:"OrderDate"`            //下单日期
	Status    libs.String `json:"status" db:"Status"`                  //审核状态
	Total     libs.Float  `json:"total" db:"Total"`                    //合计
}

type SEOrderEntry struct {
	EntryID  libs.Int    `json:"entryId" db:"EntryID" primary:"true"` //订单明细编码
	OrderNum libs.String `json:"orderNum" db:"OrderNum"`              //订单编码
	ProdNum  libs.String `json:"prodNum" db:"ProdNum"`                //产品编码
	Price    libs.Float  `json:"price" db:"Price"`                    //产品单价
	Qty      libs.Float  `json:"qty" db:"Qty"`                        //数量
	Total    libs.Float  `json:"total" db:"Total"`                    //小计
}

func (SEOrder) TableName() string {
	return "SEOrder"
}

func (SEOrderEntry) TableName() string {
	return "SEOrderEntry"
}

type SEOrderAgg struct {
	*SEOrder
	Entry []SEOrderEntry `json:"entry"`
}
