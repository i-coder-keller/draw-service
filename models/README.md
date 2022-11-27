# 集合列表

## 用户集合
```json
{
    "account": "账号",
    "password": "123456",
    "nickname": "昵称",
    "avatar": "头像",
    "email": "123@qq.com",
    "created_at": 1,
    "updated_at": 1,
    "status": 1 // 账号状态 1 在线 2离线
}
```

## 消息集合
```json
{
    "user_identity": "用户唯一标识",
    "project_identity": "项目唯一标识",
    "data": "发送的数据",
    "created_at": 1, // 创建时间
    "updated_at": 1 // 更新时间
}
```

## 项目集合
```json
{
    "owner_identity": "项目所有人唯一标识",
    "participants": [{"user_identity": "参与人唯一标识", "status": 1}],
    "data": [],
    "name": "项目名称",
    "number": "房间号",
    "info": "简介",
    "created_at": 1,
    "updated_at": 1,
}
```