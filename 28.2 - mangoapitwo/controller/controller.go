package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"main.go/model"
)

const connectingApiTwo = "mongodb+srv://mongodb:Sumitraj@cluster0.igcr8.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "apitwo"
const colName = "employeelist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectingApiTwo)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Done!")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("instance is created..")
}

//insert one data

func insertEmployee(emp model.Employee) {
	inserted, err := collection.InsertOne(context.Background(), emp)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("employee Created with id: ", inserted.InsertedID)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var emp model.Employee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	insertEmployee(emp)
	json.NewEncoder(w).Encode(emp)
}

func deleteOneEmployee(empId string) {
	id, _ := primitive.ObjectIDFromHex(empId)
	filter := bson.M{"_empid": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Employee Deleted Count: ", deleteCount)
}

func DeleteAEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneEmployee(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func getAllEmployees() []primitive.M {
	curser, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var employees []primitive.M

	for curser.Next(context.Background()) {
		var emp bson.M
		err := curser.Decode(&emp)

		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, emp)
	}
	defer curser.Close(context.Background())
	return employees
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allemp := getAllEmployees()
	json.NewEncoder(w).Encode(allemp)
}
