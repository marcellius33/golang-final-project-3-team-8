package repositories

import (
	"kanbanboard/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category *models.Category) (*models.Category, error)
	GetCategories(categories *[]models.Category) (*[]models.Category, error)
	UpdateCategory(categoryId uint, category *models.Category) (*models.Category, error)
	DeleteCategory(id uint) error
	FindCategoryById(categoryId uint) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (c *categoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	return category, c.db.Create(category).Error
}

func (c *categoryRepository) GetCategories(categories *[]models.Category) (*[]models.Category, error) {
	err := c.db.Preload("Task").Find(&categories).Error
	return categories, err
}

func (c *categoryRepository) UpdateCategory(categoryId uint, updateCategory *models.Category) (*models.Category, error) {
	category := updateCategory
	err := c.db.Model(&category).Where("id=?", categoryId).Updates(updateCategory).Error
	return category, err
}

func (c *categoryRepository) DeleteCategory(categoryId uint) error {
	err := c.db.Where("id=?", categoryId).Delete(&models.Category{}).Error
	return err
}

func (c *categoryRepository) FindCategoryById(categoryId uint) (*models.Category, error) {
	category := models.Category{}
	err := c.db.Where("id=?", categoryId).First(&category).Error
	return &category, err
}
