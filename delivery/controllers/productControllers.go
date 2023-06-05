package controllers

import (
	deliverymodels "70_Off/delivery/deliveryModels"
	"70_Off/domain/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

// Add Product
//
//	@Summary		Add Product
//	@Description	Admin Can Add Product
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			admin	body	deliverymodels.AddProductInputs	true	"Product Data"
//	@Router			/v1/admin/addProduct [post]
func (pc *ProductController) AddProduct(c *gin.Context) {
	adminIDCookie, err := c.Cookie("adminID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "admin ID not found"})
		return
	}

	adminID, err := strconv.Atoi(adminIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	var input deliverymodels.AddProductInputs

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error binding data": err.Error()})
		return
	}

	// product, err := pc.productService.ProductGetByName(input.Name)
	// if err != nil {
	// 	c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	// }
	// if product != nil {
	// 	c.JSON(http.StatusConflict, gin.H{"Product already exist, try prodct update": err.Error()})
	// 	return
	// }

	err1 := pc.productService.ProductCreate(uint(adminID), &input)
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})

	}

	c.JSON(http.StatusOK, gin.H{"message": "Product Created successfully"})

}

// Add Product Details
//
//	@Summary		Add Product Details
//	@Description	Admin Can Add Product Details
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			admin	body	deliverymodels.AddProductsDeatils	true	"Product Data"
//	@Router			/v1/admin/addProductDetails [post]
func (pc *ProductController) AddProductDetails(c *gin.Context) {
	adminIDCookie, err := c.Cookie("adminID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "admin ID not found"})
		return
	}

	adminID, err := strconv.Atoi(adminIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	var input deliverymodels.AddProductsDeatils

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error binding data": err.Error()})
		return
	}

	err1 := pc.productService.CreateProductDeatails(uint(adminID), &input)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "ProductDeatrails created successfully"})

}

// Edit Product
//
//	@Summary		Edit Product
//	@Description	Admin Can Edit Product
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			admin	body	deliverymodels.UpdateProductInput	true	"Product Data"
//	@Router			/v1/admin/updateProduct [put]
func (pc *ProductController) UpdateProduct(c *gin.Context) {

	adminIDCookie, err := c.Cookie("adminID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "admin ID not found"})
		return
	}

	adminID, err := strconv.Atoi(adminIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	var input deliverymodels.UpdateProductInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err1 := pc.productService.ProductUpdate(uint(adminID), &input)
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err1.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// Edit Product Details
//
//	@Summary		Edit Product Details
//	@Description	Admin Can Edit Product Details
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			admin	body	deliverymodels.UpdateProductsDeatils	true	"Product Details"
//	@Router			/v1/admin/updateProductDetails [put]
func (pc *ProductController) UpdateProductsDeatils(c *gin.Context) {
	adminIDCookie, err := c.Cookie("adminID")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "admin ID not found"})
		return
	}

	adminID, err := strconv.Atoi(adminIDCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}

	var input deliverymodels.UpdateProductsDeatils
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err1 := pc.productService.ProductDetailsUpdate(uint(adminID), &input)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err1.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func (pc *ProductController) ListProducts(c *gin.Context) {

	products, err := pc.productService.ProductListAll()

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// List Products
//
//	@Summary		List Products Pagination
//	@Description	List Products with Pagination
//	@Tags			user
//	@param			page	query		string	false	"Page"
//	@param			limit	query		string	false	"Limit"
//	@Success		200		{object}	deliverymodels.ViewProductDetail
//	@Router			/v1/productsPagination [get]
func (pc *ProductController) ListProductsPagination(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	// page := c.Param("page")
	// limit := c.Param("limit")

	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum < 1 || limitNum > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	offset := (pageNum - 1) * limitNum

	products, totalCount, err := pc.productService.ProductListPaginated(offset, limitNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"page":          pageNum,
		"limit":         limitNum,
		"total_records": totalCount,
		"products":      products,
	}

	c.JSON(http.StatusOK, response)
}

// Soft Delete
//
//	@Summary		Soft Delete
//	@Description	Product Soft Delete
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
//	@Param			productID	path		string	true	"ProductID "
//	@Success		200			{string}	message
//	@Router			/v1/admin/delete/{productId} [put]
func (pc *ProductController) ProductSoftDelete(c *gin.Context) {

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid productID"})
		return
	}

	err1 := pc.productService.ProductDeleteToggle(uint(productID))
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})

}

func (pc *ProductController) DeleteProduct(c *gin.Context) {

	productID, err := strconv.Atoi(c.Param("productID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid productID"})
		return
	}

	err1 := pc.productService.ProductDelete(uint(productID))
	if err1 != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})

}

// func (pc *ProductController) AddProduct(c *gin.Context) {
// 	adminID, err := strconv.Atoi(c.Param("adminID"))
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid adminID"})
// 		return
// 	}

// 	var product entity.Product
// 	if err := c.BindJSON(&product); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	product.AdminID = uint(adminID)

// 	sizes := c.PostFormArray("sizes")
// 	quantities := c.PostFormArray("quantities")

// 	if len(sizes) != len(quantities) {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Sizes and quantities mismatch"})
// 		return
// 	}

// 	for i := 0; i < len(sizes); i++ {
// 		size := sizes[i]
// 		quantity, err := strconv.Atoi(quantities[i])
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
// 			return
// 		}

// 		// Create a new size entry with the given size and quantity
// 		sizeEntry := &entity.Size{
// 			SizeName: size,
// 		}

// 		switch size {
// 		case "small":
// 			sizeEntry.Smalls = append(sizeEntry.Smalls, entity.Small{
// 				Quantity: uint(quantity),
// 			})
// 		case "medium":
// 			sizeEntry.Mediums = append(sizeEntry.Mediums, entity.Medium{
// 				Quantity: uint(quantity),
// 			})
// 		case "large":
// 			sizeEntry.Larges = append(sizeEntry.Larges, entity.Large{
// 				Quantity: uint(quantity),
// 			})
// 		default:
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid size"})
// 			return
// 		}

// 		// Add the size entry to the product's sizes
// 		product.Sizes = append(product.Sizes, *sizeEntry)
// 	}

// 	addedProduct, err := pc.productService.Create(&product)
// 	if err != nil {
// 		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, addedProduct)
// }
