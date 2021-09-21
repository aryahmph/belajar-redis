package repository

import (
	"belajar-redis/helper"
	"belajar-redis/model/web"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type StudentCacheImpl struct {
}

func NewStudentCacheImpl() *StudentCacheImpl {
	return &StudentCacheImpl{}
}

func (s *StudentCacheImpl) SetAll(ctx context.Context, client *redis.Client, response []web.StudentResponse) {
	bytes, err := json.Marshal(response)
	helper.PanicIfError(err)

	err = client.Set(ctx, "studentResponses", string(bytes), 0).Err()
	helper.PanicIfError(err)
	log.Println("set all cache")
}

func (s *StudentCacheImpl) FindAll(ctx context.Context, client *redis.Client) ([]web.StudentResponse, error) {
	result, err := client.Get(ctx, "studentResponses").Result()
	if err != nil {
		return nil, err
	}
	var studentResponses []web.StudentResponse
	err = json.Unmarshal([]byte(result), &studentResponses)
	helper.PanicIfError(err)

	log.Println("find all cache")
	return studentResponses, nil
}

func (s *StudentCacheImpl) SetById(ctx context.Context, client *redis.Client, response web.StudentResponse) {
	bytes, err := json.Marshal(response)
	helper.PanicIfError(err)

	key := fmt.Sprintf("%s%d", "studentResponse", response.Id)
	err = client.Set(ctx, key, string(bytes), 0).Err()
	helper.PanicIfError(err)

	log.Println("set by id cache")
}

func (s *StudentCacheImpl) FindById(ctx context.Context, client *redis.Client, studentId uint) (web.StudentResponse, error) {
	key := fmt.Sprintf("%s%d", "studentResponse", studentId)
	result, err := client.Get(ctx, key).Result()
	if err != nil {
		return web.StudentResponse{}, err
	}
	studentResponse := web.StudentResponse{}
	err = json.Unmarshal([]byte(result), &studentResponse)
	helper.PanicIfError(err)

	log.Println("find by id cache")
	return studentResponse, nil
}

func (s *StudentCacheImpl) Delete(ctx context.Context, client *redis.Client, studentId uint) {
	key := fmt.Sprintf("%s%d", "studentResponse", studentId)
	err := client.Set(ctx, key, "", time.Millisecond).Err()
	helper.PanicIfError(err)

	err = client.Set(ctx, "studentResponses", "", time.Millisecond).Err()
	helper.PanicIfError(err)

	log.Println("delete cache")
}
