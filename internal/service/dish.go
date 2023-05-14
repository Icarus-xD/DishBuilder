package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Icarus-xD/DishBuilder/internal/repository"
	"github.com/Icarus-xD/DishBuilder/pkg/combination"
	"github.com/Icarus-xD/DishBuilder/pkg/str"
	"github.com/gin-gonic/gin"
)

type DishRepository interface {
	GetByType(ctx context.Context, code sql.NullString) ([]repository.GetByTypeRow, error)
}

type DishService struct {
	repo DishRepository
}

type Product struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Dish struct {
	Products []Product `json:"products"`
	Price    float32	 `json:"price"`
}

func NewDishService(repo DishRepository) *DishService {
	return &DishService{
		repo: repo,
	}
}

func (d *DishService) GetByIngredients(ctx *gin.Context) {
	recipe := ctx.Query("recipe")
	if len(recipe) == 0 {
		ctx.JSON(http.StatusBadRequest, errors.New("request should contain 'recipe' query param"))
	}

	recipeChars := str.UniqueChars(recipe)

	dishes := make([]Dish, 0)
	for _, char := range recipeChars {
		ingredient := string(char)

		rows, err := d.repo.GetByType(ctx, sql.NullString{
			String: ingredient,
			Valid: true,
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		rowsLen := len(rows)
		ingredientCount := strings.Count(recipe, ingredient)
		if rowsLen < ingredientCount {
			ctx.JSON(http.StatusBadRequest, fmt.Errorf("too much %s ingredients", ingredient))
		}

		combinations := getDishes(combination.GetCombinations(rows, ingredientCount))
		combinationsLen := len(combinations)

		if len(dishes) == 0 {
			dishes = append(dishes, combinations...)
		} else if combinationsLen == 1 {
			for i, dish := range dishes {
				c := combinations[0]

				dish.Products = append(dish.Products, c.Products...)
				dish.Price += c.Price

				dishes[i] = dish
			}
		} else {
			currDishes := dishes

			for i := 0; i < combinationsLen - 1; i++ {
				dishes = append(dishes, currDishes...)
			}

			combinationPointer := 0
			for i, dish := range dishes {
				c := combinations[combinationPointer]
				combinationPointer++

				dish.Products = append(dish.Products, c.Products...)
				dish.Price += c.Price

				dishes[i] = dish

				if combinationPointer == combinationsLen - 1 {
					combinationPointer = 0
				}
			}
		}
	}

	ctx.JSON(http.StatusOK, dishes)
}

func getDishes(combinations [][]repository.GetByTypeRow) []Dish {
	dishes := make([]Dish, len(combinations))

	for i, combination := range combinations {
		dish := Dish{}

		for _, ingredient := range combination {
			dish.Products = append(dish.Products, Product{
				Type: ingredient.Type.String,
				Value: ingredient.Title,
			})

			price, _ := strconv.ParseFloat(ingredient.Price, 32)
			dish.Price += float32(price)
		}

		dishes[i] = dish
	}

	return dishes
}