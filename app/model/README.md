# Guideline Model

1. I suggest file `name.go` should be match with `table name`
2. The struct name should be `singular`
3. The method name should be clearly when you read at the first time
   1. example `GetProductByID` `GetProducts`
4. I suggest 2 way to using ORM
   1. With Simple ORM
   2. With Bun ability (tags)
   ![tags](https://drive.google.com/file/d/1BOwtm72ub6fVfqs-n26GpGUW-M0EynNP/view?usp=sharing)

If you're using 4.1, you can manage data at `controller`
example at `controler/product/product.go`


If you're using 4.2, you should be completely add tags and return interface
![4.2](https://drive.google.com/file/d/1JLn-JJe3oHq8yvZdmKTYbeQsOkzKzNWK/view?usp=sharing)
example at `model/products.go`

