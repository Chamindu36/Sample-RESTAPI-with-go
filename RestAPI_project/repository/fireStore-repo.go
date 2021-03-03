package repository

import (
	"../entity"
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"google.golang.org/api/option"
	"log"
)

type repo struct{}

// New FireStoreRepository
func NewFireStoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId            string = "chamma-rest"
	collectionIdentifier string = "posts"
	jsonPath             string = "/home/chamindu/Desktop/Choreo/TestProject/chamma-rest-firebase-adminsdk-dxomi-0774d571c1.json"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionIdentifier).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	var postList []entity.Post
	iterator := client.Collection("posts").Documents(ctx)
	allEntries, _ := iterator.GetAll()
	for i := 0; i < len(allEntries); i++ {
		entry := allEntries[i]
		post := entity.Post{
			Id:    entry.Data()["Id"].(int64),
			Title: entry.Data()["Title"].(string),
			Text:  entry.Data()["Text"].(string),
		}
		postList = append(postList, post)
	}
	return postList, nil
}

func (*repo) FindOne(title string) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	iterator := client.Collection("posts").Documents(ctx)
	allEntries, _ := iterator.GetAll()
	for i := 0; i < len(allEntries); i++ {
		entry := allEntries[i]
		post := entity.Post{
			Id:    entry.Data()["Id"].(int64),
			Title: entry.Data()["Title"].(string),
			Text:  entry.Data()["Text"].(string),
		}
		if post.Title == title {
			return &post,nil
		}
	}
	return nil, errors.New("The requested post with the title is not existed")
}
