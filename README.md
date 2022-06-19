# 2022字节青训营-抖音项目

## 一、项目简介

实现极简版抖音后端服务，能够实现视频流的播放与上传、用户的管理，点赞、评论以及用户之间的关系管理。

## 二、成员分工

| 功能项                           | 说明                                                         |
| -------------------------------- | ------------------------------------------------------------ |
| 视频 Feed 流、视频投稿、个人信息 | 支持所有用户刷抖音，按投稿时间倒序推出，登录用户可以自己拍视频投稿，查看自己的基本信息和投稿列表，注册用户流程简化。 |
| 点赞列表、用户评论               | 登录用户可以对视频点赞，并在视频下进行评论，在个人主页能够查看点赞视频列表。 |
| 关注列表、粉丝列表               | 登录用户可以关注其他用户，能够在个人信息页查看本人的关注数和粉丝数，点击打开关注列表和粉丝列表。 |

+ 小组成员

| 姓名   | 职责 | 学校 | github                                                       |
| ------ | ---- | ---- | ------------------------------------------------------------ |
| 孔鹏程 | 组长 | CUG  | [PCVocaloid (github.com)](https://github.com/PCVocaloid)     |
| 林正清 | 组员 | NEU  | [Usigned (Qing) (github.com)](https://github.com/Usigned)    |
| 岳金钊 | 组员 | CUG  | [iversonll (github.com)](https://github.com/iversonll)       |
| 李兆英 | 组员 | CUG  | [Honyelchak (Honyelchak) (github.com)](https://github.com/Honyelchak) |
| 陈子默 | 组员 | CUG  |                                                              |
| 张茗韦 | 组员 | CUG  | https://github.com/jackchen996                               |

+ 团队分工

|                 | 成员           | 职责                           |
| --------------- | -------------- | ------------------------------ |
| 基础接口        | 林正清、岳金钊 | 视频Feed流、视频投稿、个人信息 |
| 扩展接口 --- I  | 孔鹏程         | 登陆注册、点赞列表、用户评论   |
| 扩展接口 --- II | 张铭韦、李兆英 | 粉丝列表、关注列表、关注与取关 |

## 三、文件目录结构

├── controller             # 控制层
├── dao                      # 数据访问层，定义数据模型与数据库操作
├── entity                   # 实体类，定义接口所需数据结构体
├── pack                    # 打包工具
├── public                  # 公共资源，存储用户上传的视频
├── service                # 逻辑层部分，负责数据访问层与控制层的逻辑
├── utils                     # 工具类
├── go.mod                # 包依赖管理
├── go.sum                # 依赖管理
├── main.go               # 入口文件
└── router.go              # 路由信息

## 四、技术栈

+ 语言：Go

+ 框架：Gin + Gorm
  + Gin：用于web开发，获取用户信息
  + Gorm：用于数据库连接及CRUD，避免SQL注入的风险

+ 数据库：MySQL8.0（数据存储）

  + 封面提取：ffmpeg

+ 通过ffmpeg将视频解码，将首帧视频数据通过Imaging解码为图像作为封面

## 五、数据库设计

![1655607215519](C:\Users\Elichika.DESKTOP-M2GHUM1\AppData\Roaming\Typora\typora-user-images\1655607215519.png)

### 5.1 用户表（MySQL）

```SQL
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '密码',
  `follow_count` bigint DEFAULT NULL COMMENT '关注人数',
  `follower_count` bigint DEFAULT NULL COMMENT '粉丝数',
  `video_count` bigint DEFAULT NULL COMMENT '作品数',
  `like_count` bigint DEFAULT NULL COMMENT '点赞数',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
```

### 5.2 视频表（MySQL）

```SQL
CREATE TABLE `videos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `author_id` bigint DEFAULT NULL COMMENT '作者id',
  `play_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '视频url',
  `cover_url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '封面url',
  `title` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '文案',
  `create_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `favorite_count` bigint DEFAULT NULL COMMENT '点赞数',
  `comment_count` bigint DEFAULT NULL COMMENT '评论数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
```

### 5.3 评论表（MySQL）

```SQL
CREATE TABLE `comments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `video_id` bigint NOT NULL COMMENT '视频id',
  `user_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '用户名',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '评论内容',
  `create_at` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
```

### 5.4 点赞表（MySQL）

```SQL
CREATE TABLE `favorites` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_token` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '用户token',
  `video_id` bigint DEFAULT NULL COMMENT '视频id',
  `create_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
```

### 5.5 关注表（MySQL）

```SQL
CREATE TABLE `attention` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '关系id唯一值',
  `user_id` bigint DEFAULT NULL COMMENT '用户id',
  `to_user_id` bigint DEFAULT NULL COMMENT '被关注用户id',
  `is_follow` tinyint(1) DEFAULT NULL COMMENT '是否已关注',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC;
```

### 5.6 登录状态表（MySQL）

```SQL
CREATE TABLE `login_statuses` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL COMMENT '用户id',
  `token` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '用户token',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;
```

## 六、接口设计

### 6.1 基础接口

#### 6.1.1 视频流 -- /douyin/feed

> 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个

##### 接口类型

GET

##### 请求参数

| 参数名      | 位置  | 类型   | 必填 | 说明                                                         |
| :---------- | :---- | :----- | :--: | :----------------------------------------------------------- |
| latest_time | query | int64  |  否  | 说明：可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间 |
| token       | query | string |  否  | 说明：用户登录状态下设置                                     |

##### 接口定义

```go
// 请求
type FeedRequest struct {
   Token     string `form:"token,,omitempty"`
   LastTime  int64  `json:"last_time,omitempty"`
}
// 响应
type FeedResponse struct {
   entity.Response
   VideoList []entity.Video `json:"video_list"`
   NextTime  int64          `json:"next_time"`
}
```

##### 测试结果



#### 6.1.2 用户注册 -- /douyin/user/register/

> 新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token

##### 接口类型

POST

##### 请求参数

| 参数名   | 位置  | 类型   | 必填 | 说明                           |
| :------- | :---- | :----- | :--: | :----------------------------- |
| username | query | string |  是  | 说明：注册用户名，最长32个字符 |
| password | query | string |  是  | 说明：密码，最长32个字符       |

##### 接口定义

```go
// 请求
type UserLoginRequest struct {
   UserName string  `form:"username" binding:“required,min=1,max=32”`
    Password string `form:"password" binding:“required,min=6,max=32”`
}
// 响应
type UserLoginResponse struct {
   entity.Response
   UserId int64  `json:"user_id,omitempty"`
   Token  string `json:"token"`
}
```

##### 测试结果



#### 6.1.3 用户登录 -- /douyin/user/login/

##### 接口类型

POST

##### 请求参数

| 参数名   | 位置  | 类型   | 必填 | 说明             |
| :------- | :---- | :----- | :--: | :--------------- |
| username | query | string |  是  | 说明：登录用户名 |
| password | query | string |  是  | 说明：登录密码   |

##### 接口定义

```go
// 请求
type UserLoginRequest struct {
   UserName string  `form:"username" binding:“required,min=1,max=32”`
    Password string `form:"password" binding:“required,min=6,max=32”`
}
// 响应
type UserLoginResponse struct {
   entity.Response
   UserId int64  `json:"user_id,omitempty"`
   Token  string `json:"token"`
}
```

##### 测试结果



#### 6.1.4 用户信息 -- /douyin/user/

> 获取用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数

##### 接口类型

GET

##### 请求参数

| 参数名  | 位置  | 类型   | 必填 | 说明                |
| :------ | :---- | :----- | :--: | :------------------ |
| user_id | query | int64  |  是  | 说明：用户id        |
| token   | query | string |  是  | 说明：用户鉴权token |

##### 接口定义

```go
// 请求
type UserInfoRequest struct {
    UserId int64  `form:"user_id" binding:"required"`
    Token  string `form:"token" binding:"required"`
}
// 响应
type UserResponse struct {
   entity.Response
   User entity.User `json:"user"`
}
```

##### 测试结果



#### 6.1.5 视频投稿 -- /douyin/publish/action/

> 登录用户选择视频上传

##### 接口类型

POST

##### 请求参数

| 参数名 | 类型 | 必填 | 说明                |
| :----- | :--- | :--- | :------------------ |
| data   | file | 是   | 说明：视频数据      |
| token  | text | 是   | 说明：用户鉴权token |
| title  | text | 是   | 说明：视频标题      |

##### 接口定义

```go
// 请求
type PublishRequest struct {
   Data  *multipart.FileHeader  `form:"data" binding:"required"`
   Token string `form:"token" binding:"required"`
   TiTle string `form:"title" binding:"required"`
}
// 响应
type PublishResponse struct {
   entity.Response
}
```



##### 测试结果



#### 6.1.6 发布列表 -- /douyin/publish/list/

> 用户的视频发布列表，直接列出用户所有投稿过的视频

##### 接口类型

GET

##### 请求参数

| 参数名  | 位置  | 类型   | 必填 | 说明                |
| :------ | :---- | :----- | :--: | :------------------ |
| token   | query | string |  是  | 说明：用户鉴权token |
| user_id | query | int64  |  是  | 说明：用户id        |

##### 接口定义

```go
// 请求
type VideoListRequest struct {
   UserId int64  'form:"user_id" binding:"required"'
   Token  string 'form:"token" binding:"required"'
}
// 响应
type VideoListResponse struct {
   entity.Response
   VideoList []entity.Video `json:"video_list"`
}
```

##### 测试结果



### 6.2 扩展接口 -- I

#### 6.2.1 赞操作 -- /douyin/favorite/action/

> 登录用户对视频的点赞和取消点赞操作

##### 接口类型

POST

##### 请求参数

| 参数名      | 位置  | 类型   | 必填 | 说明                     |
| :---------- | :---- | :----- | :--: | :----------------------- |
| token       | query | string |  是  | 说明：用户鉴权token      |
| video_id    | query | int64  |  是  | 说明：视频id             |
| action_type | query | int64  |  是  | 说明：1-点赞，2-取消点赞 |

##### 接口定义

```go
// 请求
type FavoriteActionRequest struct {
   UserId     int64  `json:"user_id,omitempty" binding:"required"`
   Token      string `json:"token" binding:"required"`
   VideoId    int64  `json:"video_id" binding:"required"`
   ActionType int64  `json:"action_type" binding:"required，oneof=1 2"`
}
// 响应
type FavoriteActionResponse struct {
   entity.Response
}
```

##### 测试结果



#### 6.2.2 点赞列表 -- /douyin/favorite/list/

> 用户的所有点赞视频

##### 接口类型

GET

##### 请求参数

| 参数名  | 位置  | 类型   | 必填 | 说明                |
| :------ | :---- | :----- | :--: | :------------------ |
| user_id | query | int64  |  是  | 说明：用户id        |
| token   | query | string |  是  | 说明：用户鉴权token |

##### 接口定义

```go
// 请求
type FavoriteListRequest struct {
   UserId int64  `form:"user_id" binding:"required"`
   Token  string `form:"token" binding:"required"`
}
// 响应
type FavoriteListResponse struct {
   entity.Response
   VideoList []entity.Video `json:"video_list"`
}
```

##### 测试结果



#### 6.2.3 评论操作 -- /douyin/comment/action/

> 登录用户对视频进行评论

##### 接口类型

POST

##### 请求参数

| 参数名       | 位置  | 类型   | 必填 | 说明                                                |
| :----------- | :---- | :----- | :--: | :-------------------------------------------------- |
| token        | query | string |  是  | 说明：用户鉴权token                                 |
| video_id     | query | int64  |  是  | 说明：视频id                                        |
| action_type  | query | int64  |  是  | 说明：1-发布评论，2-删除评论                        |
| comment_text | query | string |  否  | 说明：用户填写的评论内容，在action_type=1的时候使用 |
| comment_id   | query | int64  |  否  | 说明：要删除的评论id，在action_type=2的时候使用     |

##### 接口定义

```go
// 请求
type CommentActionRequest struct {
   Token       string `form:"token" binding:"requeired"`
   VideoId     int64  `form:"video_id" binding:"requeired"`
   ActionType  int64  `form:"action_type" binding:"requeired,oneof=1 2"`
   CommentText string `form:"comment_text" binding:"omitempty"`
   CommentId   int64  `form:"comment_id" binding:"omitempty"`
}
// 响应
type CommentActionResponse struct {
   entity.Response
   Comment entity.Comment `json:"comment,omitempty"`
}
```

##### 测试结果



#### 6.2.4 评论列表 -- /douyin/comment/list/

> 查看视频的所有评论，按发布时间倒序

##### 接口类型

GET

##### 请求参数

| 参数名   | 位置  | 类型   | 必填 | 说明                |
| :------- | :---- | :----- | :--: | :------------------ |
| token    | query | string |  是  | 说明：用户鉴权token |
| video_id | query | int64  |  是  | 说明：视频id        |

##### 接口定义

```go
// 请求
type CommentListRequest struct {
    Token   string `form:"token" binding:"requeired"`
    VideoId int64  `form:"video_id" binding:"requeired"`
}
// 响应
type CommentListResponse struct {
   entity.Response
   CommentList []entity.Comment `json:"comment_list,omitempty"`
}
```

##### 测试结果



### 6.3 扩展接口 -- II

#### 6.3.1 关注操作 -- /douyin/relation/action/

> 登录用户关注或者取消关注某个用户

##### 接口类型

POST

##### 请求参数

| 参数名      | 位置  | 类型   | 必填 | 说明                     |
| :---------- | :---- | :----- | :--: | :----------------------- |
| token       | query | string |  是  | 说明：用户鉴权token      |
| to_user_id  | query | int64  |  是  | 说明：对方用户id         |
| action_type | query | int64  |  是  | 说明：1-关注，2-取消关注 |

##### 接口定义

```go
// 请求
type RelatiionRequest struct {
   Token      string `form:"token" binding:"required"`
   ToUserId   int64  `form:"to_user_id" binding:"required"`
   ActionType int64  `form:"action_type" binding:"required,oneof=1 2"`
}
// 响应
type UserListResponse struct {
   entity.Response
   UserList []entity.User `json:"user_list"`
}
```

##### 测试结果



#### 6.3.2 关注列表 -- /douyin/relation/follow/list/

> 获得用户的关注列表

##### 接口类型

GET

##### 请求参数

| 参数名  | 位置  | 类型   | 必填 | 说明                |
| :------ | :---- | :----- | :--: | :------------------ |
| user_id | query | int64  |  是  | 说明：用户id        |
| token   | query | string |  是  | 说明：用户鉴权token |

##### 接口定义

```go
// 请求
type FollowListRequest struct {
   UserId int64  `form:"user_id" binding:"required"`
   Token  string `form:"token" binding:"required"`
}
// 响应
type UserListResponse struct {
   entity.Response
   UserList []entity.User `json:"user_list"`
}
```

##### 测试结果



#### 6.3.3 粉丝列表 -- /douyin/relation/follower/list/

> 获得用户的粉丝列表

##### 接口类型

GET

##### 请求参数

| 参数名  | 位置  | 类型   | 必填 | 说明                |
| :------ | :---- | :----- | :--: | :------------------ |
| user_id | query | int64  |  是  | 说明：用户id        |
| token   | query | string |  是  | 说明：用户鉴权token |

##### 接口定义

```go
// 请求
type FollowListRequest struct {
   UserId int64  `form:"user_id" binding:"required"`
   Token  string `form:"token" binding:"required"`
}
// 响应
type UserListResponse struct {
   entity.Response
   UserList []entity.User `json:"user_list"`
}
```

##### 测试结果

