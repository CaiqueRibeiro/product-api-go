package database

import (
	"testing"

	"github.com/CaiqueRibeiro/product-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserSuiteTest struct {
	suite.Suite
	db *gorm.DB
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuiteTest))
}

// roda antes de tudo (beforeAll do jest)
func (t *UserSuiteTest) SetupSuite() {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.FailNow(err.Error())
	}
	t.db = db
	t.db.AutoMigrate(&entity.User{})
}

// Roda depois de cada teste (afterEach do jest)
func (t *UserSuiteTest) TearDownTest() {
	t.db.Where("1 = 1").Delete(&entity.User{}) // Deleta todos os registros da tabela
}

// Roda depois de tudo (afterAll do jest)
func (t *UserSuiteTest) TearDownSuite() {
	dbInstance, err := t.db.DB()
	if err != nil {
		t.FailNow(err.Error())
	}
	dbInstance.Close()
}

func (t *UserSuiteTest) TestCreate() {
	repo := NewUserRepository(t.db)
	user, _ := entity.NewUser("Caique", "caique@gmail.com", "12345")
	err := repo.Create(user)

	assert.NoError(t.T(), err)

	var foundUser entity.User
	err = t.db.First(&foundUser, "id = ?", user.ID.String()).Error

	assert.NoError(t.T(), err)

	assert.Equal(t.T(), user.ID, foundUser.ID)
	assert.Equal(t.T(), user.Name, foundUser.Name)
	assert.Equal(t.T(), user.Email, foundUser.Email)
	assert.Equal(t.T(), user.Password, foundUser.Password)
}

func (t *UserSuiteTest) TestFindByEmail() {
	repo := NewUserRepository(t.db)
	user, _ := entity.NewUser("Caique", "caique@gmail.com", "12345")
	repo.Create(user)

	foundUser, err := repo.FindByEmail(user.Email)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), user.ID, foundUser.ID)
	assert.Equal(t.T(), user.Name, foundUser.Name)
	assert.Equal(t.T(), user.Email, foundUser.Email)
}
func (t *UserSuiteTest) TestFindByID() {
	repo := NewUserRepository(t.db)
	user, _ := entity.NewUser("Caique", "caique@gmail.com", "12345")
	repo.Create(user)

	foundUser, err := repo.FindByID(user.ID.String())

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), user.ID, foundUser.ID)
	assert.Equal(t.T(), user.Name, foundUser.Name)
	assert.Equal(t.T(), user.Email, foundUser.Email)
}

func (t *UserSuiteTest) TestFindAll() {
	repo := NewUserRepository(t.db)

	user1, _ := entity.NewUser("Caique", "caique@gmail.com", "12345")
	user2, _ := entity.NewUser("Zezinho", "zezinho@gmail.com", "111")
	user3, _ := entity.NewUser("Marcela", "marcela@gmail.com", "442323")

	repo.Create(user1)
	repo.Create(user2)
	repo.Create(user3)

	users, err := repo.FindAll(1, 10, "asc")

	assert.NoError(t.T(), err)
	assert.Len(t.T(), users, 3)
}

func (t *UserSuiteTest) TestUpdate() {
	repo := NewUserRepository(t.db)
	user, _ := entity.NewUser("Caique", "caique@gmail.com", "12345")
	repo.Create(user)

	user.Name = "Caique Ribeiro"

	err := repo.Update(user)

	assert.NoError(t.T(), err)

	var foundUser entity.User
	err = t.db.First(&foundUser, "id = ?", user.ID.String()).Error

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), user.ID, foundUser.ID)
	assert.Equal(t.T(), user.Name, foundUser.Name)
	assert.Equal(t.T(), user.Email, foundUser.Email)
}

func (t *UserSuiteTest) TestDelete() {
	repo := NewUserRepository(t.db)
	user, _ := entity.NewUser("Caique", "caique@gmail.com", "12345")
	repo.Create(user)

	err := repo.Delete(user.ID.String())

	assert.NoError(t.T(), err)

	var foundUser entity.User
	err = t.db.First(&foundUser, "id = ?", user.ID.String()).Error

	assert.Error(t.T(), err)
}
