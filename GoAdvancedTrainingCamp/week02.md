# Week02

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

```go
// References Code [http](https://github.com/go-kratos/beer-shop/blob/main/app/user/service/internal/data/user.go#41)
// 统一错误处理，分配唯一错误码，在Dao层返回相应错误结构的的错误码，底层SentinelError可以使用wrapf的方式记录
func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
    po, err := r.data.db.User.Get(ctx, id)
    switch {
    case err == sql.ErrNoRows:
        return nil, errors.Wrapf(errors.New(10001, "USER_NOT_FOUND"), err.Error())
    case err != nil:
        return nil, errors.Wrapf(errors.New(10002, "SQL_ERR"), err.Error())
    default:
        return &biz.User{Id: po.ID, Username: po.Username}, err
    }
}
```
