package database

import (
	"testing"
	"time"

	"github.com/CaiqueRibeiro/product-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductSuiteTest struct {
	suite.Suite
	db *gorm.DB
}

func TestProductSuite(t *testing.T) {
	suite.Run(t, new(ProductSuiteTest))
}

// roda antes de tudo (beforeAll do jest)
func (t *ProductSuiteTest) SetupSuite() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.FailNow(err.Error())
	}
	t.db = db
	t.db.AutoMigrate(&entity.Product{})
}

// Roda depois de cada teste (afterEach do jest)
func (t *ProductSuiteTest) TearDownTest() {
	t.db.Where("1 = 1").Delete(&entity.Product{}) // Deleta todos os registros da tabela
}

// Roda depois de tudo (afterAll do jest)
func (t *ProductSuiteTest) TearDownSuite() {
	dbInstance, err := t.db.DB()
	if err != nil {
		t.FailNow(err.Error())
	}
	dbInstance.Close()
}

func (t *ProductSuiteTest) TestCreate() {
	repo := NewProductRepository(t.db)
	product, _ := entity.NewProduct("Produto 1", 10, time.Now())
	err := repo.Create(product)

	assert.NoError(t.T(), err)

	var foundProduct entity.Product
	err = t.db.First(&foundProduct, "id = ?", product.ID.String()).Error

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), product.ID, foundProduct.ID)
	assert.Equal(t.T(), product.Name, foundProduct.Name)
	assert.Equal(t.T(), product.Price, foundProduct.Price)
}

func (t *ProductSuiteTest) TestFindAll() {
	repo := NewProductRepository(t.db)

	product1, _ := entity.NewProduct("Produto 1", 10, time.Now())
	product2, _ := entity.NewProduct("Produto 2", 20, time.Now())
	product3, _ := entity.NewProduct("Produto 3", 30, time.Now())

	repo.Create(product1)
	repo.Create(product2)
	repo.Create(product3)

	products, err := repo.FindAll(1, 10, "asc")

	assert.NoError(t.T(), err)
	assert.Len(t.T(), products, 3)
}

func (t *ProductSuiteTest) TestFindByID() {
	repo := NewProductRepository(t.db)
	product, _ := entity.NewProduct("Produto 1", 10, time.Now())
	repo.Create(product)

	foundProduct, err := repo.FindByID(product.ID.String())

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), product.ID, foundProduct.ID)
	assert.Equal(t.T(), product.Name, foundProduct.Name)
	assert.Equal(t.T(), product.Price, foundProduct.Price)
}

func (t *ProductSuiteTest) TestUpdate() {
	repo := NewProductRepository(t.db)
	product, _ := entity.NewProduct("Produto 1", 10, time.Now())
	repo.Create(product)

	product.Name = "Produto 2"
	product.Price = 20

	err := repo.Update(product)

	assert.NoError(t.T(), err)

	var foundProduct entity.Product
	err = t.db.First(&foundProduct, "id = ?", product.ID.String()).Error

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), product.ID, foundProduct.ID)
	assert.Equal(t.T(), product.Name, foundProduct.Name)
	assert.Equal(t.T(), product.Price, foundProduct.Price)
}
