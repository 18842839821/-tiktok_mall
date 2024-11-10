package Product

import (
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	mlog "mall/log"
	"mall/model/api"
	"mall/service/product/proto/product"
	"net/http"
)

var ProductClient product.ProductCatalogServiceClient

func Init(engine *gin.Engine) {
	ProductConn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:4379"},
			Key:   "order.rpc",
		},
	})
	ProductClient = product.NewProductCatalogServiceClient(ProductConn.Conn())
	mlog.SetName("ProductAPI")
	group := engine.Group("/Product")
	{
		group.POST("/List", List)
		group.POST("Get", Get)
		group.POST("/Search", Search)
		group.POST("/Create", Create)
		group.POST("/Update", Update)
		group.POST("/delete", Delete)
	}
}

func List(c *gin.Context) {
	req := api.ListProductsReq{}
	if err := c.ShouldBind(&req); err != nil {
		mlog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	resp, _ := ProductClient.ListProducts(c, &product.ListProductsReq{
		Page:     uint32(req.Page),
		PageSize: uint32(req.PageSize),
	})
	c.JSON(http.StatusOK, resp)
}

func Get(c *gin.Context) {
	req := product.GetProductReq{}
	if err := c.ShouldBind(&req); err != nil {
		mlog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	resp, _ := ProductClient.GetProduct(c, &req)
	c.JSON(http.StatusOK, resp)
}

func Search(c *gin.Context) {
	req := product.SearchProductsReq{}
	if err := c.ShouldBindUri(&req); err != nil {
		mlog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	resp, _ := ProductClient.SearchProducts(c, &req)
	c.JSON(http.StatusOK, resp)
}

func Create(c *gin.Context) {
	req := product.CreateProductsReq{}
	if err := c.ShouldBind(&req); err != nil {
		mlog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	resp, _ := ProductClient.CreateProducts(c, &req)
	c.JSON(http.StatusOK, resp)
}

func Update(c *gin.Context) {
	req := product.UpdateProductsReq{}
	if err := c.ShouldBind(&req); err != nil {
		mlog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	resp, _ := ProductClient.UpdateProducts(c, &req)
	c.JSON(http.StatusOK, resp)
}

func Delete(c *gin.Context) {
	req := product.DeleteProductsReq{}
	if err := c.ShouldBind(&req); err != nil {
		mlog.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	resp, _ := ProductClient.DeleteProducts(c, &req)
	c.JSON(http.StatusOK, resp)
}
