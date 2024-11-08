package handlers

import (
	"context"
	"encoding/json"
	"hack3/db"
	"hack3/models"
	"net/http"
	"time"

	"github.com/savsgio/atreugo/v11"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNote(ctx *atreugo.RequestCtx) error {
	var note models.Note
	if err := json.Unmarshal(ctx.Request.Body(), &note); err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Invalid request"}, http.StatusBadRequest)
	}

	note.ID = primitive.NewObjectID().Hex()
	note.Owner_ID = primitive.NewObjectID().Hex()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	collection := db.GetNotesCollection()
	_, err := collection.InsertOne(context.TODO(), note)
	if err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Failed to create note"}, http.StatusInternalServerError)
	}

	return ctx.JSONResponse(note, http.StatusCreated)
}

func GetNote(ctx *atreugo.RequestCtx) error {
	noteID := ctx.UserValue("id").(string)

	var note models.Note
	collection := db.GetNotesCollection()
	err := collection.FindOne(context.TODO(), bson.M{"_id": noteID}).Decode(&note)
	if err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Note not found"}, http.StatusNotFound)
	}

	return ctx.JSONResponse(note, http.StatusOK)
}

func UpdateNote(ctx *atreugo.RequestCtx) error {
	noteID := ctx.UserValue("id").(string)
	// userID := ctx.UserValue("user_id").(string)

	var updateData models.Note
	if err := json.Unmarshal(ctx.Request.Body(), &updateData); err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Invalid request"}, http.StatusBadRequest)
	}

	collection := db.GetNotesCollection()
	var note models.Note
	err := collection.FindOne(context.TODO(), bson.M{"_id": noteID}).Decode(&note)
	if err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Note not found"}, http.StatusNotFound)
	}

	if !hasPermission(note.SharedWith, note.Owner_ID, "edit") {
		return ctx.JSONResponse(map[string]string{"error": "Forbidden"}, http.StatusForbidden)
	}

	updateData.UpdatedAt = time.Now()
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": noteID}, bson.M{"$set": updateData})
	if err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Failed to update note"}, http.StatusInternalServerError)
	}

	return ctx.JSONResponse(updateData, http.StatusOK)
}

func hasPermission(sharedWith []models.Share, userID string, permission string) bool {
	for _, share := range sharedWith {
		if share.UserID == userID && share.Permission == permission {
			return true
		}
	}
	return false
}

func DeleteNote(ctx *atreugo.RequestCtx) error {
	noteID := ctx.UserValue("id").(string)
	// userID := ctx.UserValue("user_id").(string)

	collection := db.GetNotesCollection()
	var note models.Note
	err := collection.FindOne(context.TODO(), bson.M{"_id": noteID}).Decode(&note)
	if err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Note not found"}, http.StatusNotFound)
	}

	if !hasPermission(note.SharedWith, note.Owner_ID, "edit") {
		return ctx.JSONResponse(map[string]string{"error": "Forbidden"}, http.StatusForbidden)
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": noteID})
	if err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Failed to delete note"}, http.StatusInternalServerError)
	}

	return ctx.JSONResponse(map[string]string{"message": "Note deleted"}, http.StatusOK)
}

func AddPlan(ctx *atreugo.RequestCtx) error {

	var plan models.Plan
	if err := json.Unmarshal(ctx.Request.Body(), &plan); err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Invalid request"}, http.StatusBadRequest)
	}

	plan.ID = primitive.NewObjectID().Hex()

	collection := db.GetNotesCollection()
	_, err := collection.InsertOne(context.TODO(), plan)
	if err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Failed to create note"}, http.StatusInternalServerError)
	}

	return ctx.JSONResponse(plan, http.StatusCreated)
}

func GetPlan(ctx *atreugo.RequestCtx) error {
	uid := ctx.UserValue("uid").(string)

	var plan models.Plan
	collection := db.GetNotesCollection()
	err := collection.FindOne(context.TODO(), bson.M{"_id": uid}).Decode(&plan)
	if err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Plan not found"}, http.StatusNotFound)
	}

	return ctx.JSONResponse(plan, http.StatusOK)
}
func GetAllNotes(ctx *atreugo.RequestCtx) error {
	userID := ctx.UserValue("user_id").(string)

	collection := db.GetNotesCollection()
	filter := bson.M{
		"$or": []bson.M{
			{"owner": userID},
			{"shared_with.user_id": userID},
		},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Failed to get notes"}, http.StatusInternalServerError)
	}

	var notes []models.Note
	if err := cursor.All(context.TODO(), &notes); err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Failed to get notes"}, http.StatusInternalServerError)
	}

	return ctx.JSONResponse(notes, http.StatusOK)
}

func CreateNoteForUser(ctx *atreugo.RequestCtx) error {
	uid := ctx.UserValue("uid").(string)

	var note models.Note
	if err := json.Unmarshal(ctx.Request.Body(), &note); err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Invalid request"}, http.StatusBadRequest)
	}

	note.ID = primitive.NewObjectID().Hex()
	note.Owner_ID = uid
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	collection := db.GetNotesCollection()
	_, err := collection.InsertOne(context.TODO(), note)
	if err != nil {
		return ctx.JSONResponse(map[string]string{"error": "Failed to create note"}, http.StatusInternalServerError)
	}

	return ctx.JSONResponse(note, http.StatusCreated)
}
