package handlers

import (
	"final/cmd/echo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func HandleGetAllLists(c echo.Context) error {

	result, err := models.GetAllLists()

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, result)

}

func HandleCreateList(c echo.Context) error {

	l := new(models.List)
	if err := c.Bind(&l); err != nil {
		return err
	}
	res, err := models.CreateList(l)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, &res)

}

func HandleDeleteList(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	models.DeleteList(id)

	return c.NoContent(200)
}

func HandleGetTasksByListId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	result, err := models.GetTasksByListId(id)

	return c.JSON(http.StatusOK, result)
}

func HandleCreateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	t := new(models.Task)

	if err := c.Bind(&t); err != nil {
		return err
	}
	res := models.CreateTask(id, t)

	return c.JSON(http.StatusOK, &res)

}

func HandleUpdateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	t := new(models.Task)
	if err := c.Bind(&t); err != nil {
		return err
	}
	result, err := models.UpdateTask(id, t.Completed)

	return c.JSON(http.StatusOK, result)
}

func HandleDeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}
	models.DeleteTask(id)

	return c.NoContent(200)
}
