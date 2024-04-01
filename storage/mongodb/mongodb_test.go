package mongodb

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

type User struct {
	Id       string `json:"id" bson:"_id"` // 自定义的 Id 字段
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// TestStorage_InsertAndFindOne 需要通过 Makefile 提前启动本地 MongoDB 服务
func TestStorage_InsertAndFindOne(t *testing.T) {
	config := Config{
		Username:   "root",
		Password:   "root",
		Database:   "doutok",
		Collection: "default",
	}
	store := New(config)

	// 创建用户
	_, err := store.InsertOne(context.Background(), User{
		Id:       "123", // 手动指定自定义的 Id 值
		Username: "zhangsan",
		Password: "123456",
		Email:    "123@.com",
	})
	require.NoError(t, err)

	// 查询用户
	result := store.FindOne(context.Background(), map[string]interface{}{"_id": "123"})

	// 解码查询结果
	var user map[string]interface{}
	err = result.Decode(&user)
	require.NoError(t, err)
	require.Equal(t, "123@.com", user["email"])

	// 修改用户
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"email": "321@.com",
		},
	}
	_, err = store.UpdateOne(context.Background(),
		map[string]interface{}{"_id": "123"},
		update)
	require.NoError(t, err)

	// 查询用户
	result = store.FindOne(context.Background(), map[string]interface{}{"_id": "123"})

	// 解码查询结果
	err = result.Decode(&user)
	require.NoError(t, err)
	require.Equal(t, "321@.com", user["email"])

	// 删除用户
	count, err := store.DeleteOne(context.Background(), map[string]interface{}{"_id": "123"})
	require.NoError(t, err)
	require.Equal(t, int64(1), count)

	// 查询用户
	result = store.FindOne(context.Background(), map[string]interface{}{"_id": "123"})
	user2 := User{}
	err = result.Decode(&user2)
	require.Equal(t, "mongo: no documents in result", err.Error())
}
