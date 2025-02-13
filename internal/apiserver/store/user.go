package store

import (
	"context"
	v1 "j-iam/internal/pkg/model/apiserver/v1"

	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

// UserStore 抽象出存储层的数据操作接口
type UserStore interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, usernames []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
}
