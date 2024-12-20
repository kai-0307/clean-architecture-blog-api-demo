# メモ

## 突貫工事で作ったので、コードが無茶苦茶
## 後ほど全体的に修正する

## api は Clean Architecture っぽい感じで実装

```
psql -h localhost -p 5432 -U $(whoami) -d postgres
```

```
psql -h localhost -p 5432 -U postgres -d postgres
```

psql (14.15 (Homebrew))
Type "help" for help.

```
CREATE ROLE postgres WITH SUPERUSER LOGIN PASSWORD 'password';
```

demo_user ユーザーを作成

```
CREATE ROLE demo_user WITH LOGIN PASSWORD 'password';
```

データベースへの権限を付与 tech_blog データベースにアクセスできるように権限を付与します。

```
GRANT ALL PRIVILEGES ON DATABASE tech_blog TO demo_user;
```

ロール一覧を確認して、demo_user が作成されていることを確認します：

```
\du
```

テーブルへのデータ挿入

```
INSERT INTO public.blogs (title, content, author) VALUES ('Sample Blog', 'This is a test blog content.', 'Admin');
```

確認

```
SELECT * FROM blogs;

```

```
curl -X GET http://localhost:8080/blogs
curl -X POST http://localhost:8080/blogs \
  -H "Content-Type: application/json" \
  -d '{"title":"New Blog Post", "content":"This is the content of the blog post.", "author":"Admin"}'
curl -X DELETE 'http://localhost:8080/blogs?id=5'
```

```
curl -v -X GET http://localhost:8080/blogs
```
