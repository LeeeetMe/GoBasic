package repository

import (
	"IrisProduct/common"
	"IrisProduct/datamodels"
	"database/sql"
	"fmt"
)

type IProduct interface {
	SQLConn() error
	SearchProduct(int64) (*datamodels.Product, error)
	Insert(*datamodels.Product) (int64, error)
	Update(int64, *datamodels.Product) error
	Delete(int64) bool
	SelectAll() ([]*datamodels.Product, error)
}

type ProductManager struct {
	Table   string
	SqlConn *sql.DB
}

func NewProductManager(t string, s *sql.DB) IProduct {
	p := &ProductManager{
		Table:   t,
		SqlConn: s,
	}
	return p
}

func (p *ProductManager) SQLConn() (err error) {
	if p.SqlConn == nil {
		p.SqlConn, err = common.NewMysqlConn()
		if err != nil {
			return
		}
	}
	if p.Table == "" {
		p.Table = "product"
	}
	return nil
}

func (p *ProductManager) Insert(product *datamodels.Product) (id int64, err error) {
	if err = p.SQLConn(); err != nil {
		return
	}
	stmtIn, err := p.SqlConn.Prepare("insert into product (name,image,url,num) value (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	r, err := stmtIn.Exec(product.Name, product.Image, product.Url, product.Num)
	if err != nil {
		return 0, err
	}
	id, _ = r.LastInsertId()
	return
}
func (p *ProductManager) Update(id int64, product *datamodels.Product) (err error) {
	if err = p.SQLConn(); err != nil {
		return
	}
	stmtIn, err := p.SqlConn.Prepare("update product set name=?,image=?,url=?,num=? where id=?")
	if err != nil {
		return
	}
	_, err = stmtIn.Exec(product.Name, product.Image, product.Url, product.Num, id)
	return
}

func (p *ProductManager) Delete(id int64) bool {
	if err := p.SQLConn(); err != nil {
		return false
	}
	stmtDel, err := p.SqlConn.Prepare("delete from product where id=?")
	if err != nil {
		return false
	}
	_, err = stmtDel.Exec(id)
	if err != nil {
		return false
	}
	return true
}

func (p *ProductManager) SearchProduct(id int64) (product *datamodels.Product, err error) {
	stmtOut, err := p.SqlConn.Prepare("select num,name,image,url from product where id = ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query(id)
	if err != nil {
		return
	}
	x := common.GetResultRow(rows)
	fmt.Println("---------------")
	fmt.Println(x)
	product = &datamodels.Product{}
	common.DataToStructByTagSql(x, product)
	//stmtOut.QueryRow(id).Scan(&product.Num, &product.Name, &product.Image, &product.Url)
	fmt.Println(product)
	defer stmtOut.Close()
	return
}

func (p *ProductManager) SelectAll() (products []*datamodels.Product, err error) {
	if err := p.SQLConn(); err != nil {
		return nil, err
	}
	stmtOut, err := p.SqlConn.Prepare("select * from ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query(p.Table)
	if err != nil {
		return nil, err
	}
	result := common.GetResultRows(rows)
	for _, value := range result {
		tmpProduct := &datamodels.Product{}
		common.DataToStructByTagSql(value, tmpProduct)
		products = append(products, tmpProduct)
	}
	return
}
