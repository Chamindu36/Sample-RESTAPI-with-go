package repository

import (
	"../entity"
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"log"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

// New PostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId            string = "chamma-rest"
	collectionIdentifier string = "posts"
	jsonPath string = "/home/chamindu/Desktop/Choreo/TestProject/chamma-rest-firebase-adminsdk-dxomi-0774d571c1.json"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx,projectId, option.WithCredentialsFile(jsonPath))
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
	return  post,nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx,projectId, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
		}

	defer client.Close()
	var posts [] entity.Post
	iterator := client.Collection("posts").Documents(ctx)
	allones,_:=iterator.GetAll()
	fmt.Println(allones)
	for i := 0; i < len(allones); i++ {
		doc := allones[i]
		post:=entity.Post{
			Id: doc.Data()["Id"].(int64),
			Title: doc.Data()["Title"].(string),
			Text: doc.Data()["Text"].(string),
		}
		posts = append(posts,post)
	}
	return posts, nil
}

