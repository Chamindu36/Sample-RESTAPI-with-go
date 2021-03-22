package service
import (
	"../entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)


type MockRepository struct {
	mock.Mock
}


func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
		args:= mock.Called()
		result := args.Get(0)
		return result.(*entity.Post),args.Error(1)
}

func (mock *MockRepository) FindAll() ([] entity.Post,error){
	args:= mock.Called()
	result := args.Get(0)
	return result.([] entity.Post),args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err:= testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty",err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{Id: 1, Title: "", Text: "B"}
	testService := NewPostService(nil)
	err:= testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty",err.Error())
}

func TestFindAll (t *testing.T) {
	mockRepo := new(MockRepository)
	var identifier int64 = 1
	//Setup Expectations
	post := entity.Post{Id: 1, Title: "A", Text: "B"}
	mockRepo.On("FindAll").Return([]entity.Post{post},nil)

	testService := NewPostService(mockRepo)
	result,_ := testService.FindAll()

	//Mock assertion: Behavioral
	mockRepo.AssertExpectations(t)
	// Data assertion
	assert.Equal(t, identifier, result[0].Id)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)

}

func TestCreate (t *testing.T) {
	mockRepo := new(MockRepository)
	var identifier int64 = 5577006791947779410
	post := entity.Post{Id: 5577006791947779410, Title: "A", Text: "B"}

	mockRepo.On("Save").Return(&post,nil)
	testService := NewPostService(mockRepo)
	result,err := testService.Create(&post)

	//Mock assertion: Behavioral
	mockRepo.AssertExpectations(t)
	// Data assertion
	assert.Equal(t, identifier, result.Id)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t,err)
}