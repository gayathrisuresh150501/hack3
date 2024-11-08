package handlers

import (
	"context"
	"encoding/json"
	"hack3/config"
	"hack3/models"
	"net/http"

	"github.com/savsgio/atreugo/v11"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateNote(ctx *atreugo.RequestCtx) error {
	var note models.Note
	if err := json.Unmarshal(ctx.PostBody(), &note); err != nil {
		return ctx.ErrorResponse(err, http.StatusBadRequest)
	}

	collection := config.Client.Database("noteit").Collection("notes")
	_, err := collection.InsertOne(context.TODO(), note)
	if err != nil {
		return ctx.ErrorResponse(err, http.StatusInternalServerError)
	}

	ctx.SetStatusCode(http.StatusCreated)
	return nil
}

func GetNotes(ctx *atreugo.RequestCtx) error {
	ownerID := string(ctx.QueryArgs().Peek("owner_id"))
	collection := config.Client.Database("noteit").Collection("notes")

	cursor, err := collection.Find(context.TODO(), bson.M{"owner_id": ownerID})
	if err != nil {
		return ctx.ErrorResponse(err, http.StatusInternalServerError)
	}
	defer cursor.Close(context.TODO())

	var notes []models.Note
	if err := cursor.All(context.TODO(), &notes); err != nil {
		return ctx.ErrorResponse(err, http.StatusInternalServerError)
	}

	response, err := json.Marshal(notes)
	if err != nil {
		return ctx.ErrorResponse(err, http.StatusInternalServerError)
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(response)
	return nil
}
