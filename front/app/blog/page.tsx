"use client";

import { useEffect, useState } from "react";
import Header from "../components/header";
import Footer from "../components/footer";

type Blog = {
  id: number;
  title: string;
  content: string;
  author: string;
  createdAt: string;
};

const BlogList = () => {
  const [blogs, setBlogs] = useState<Blog[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    fetch("http://localhost:8080/blogs")
      .then(async (res) => {
        if (!res.ok) {
          const errorText = await res.text();
          throw new Error(
            `HTTP error! status: ${res.status}, message: ${errorText}`
          );
        }
        return res.json();
      })
      .then((data) => setBlogs(data))
      .catch((err) => setError(err.message));
  }, []);

  // ブログ作成の即時反映
  const handleCreate = async () => {
    setLoading(true);
    const newBlog = {
      title: "New Blog Post",
      content: "This is the content of the blog post.",
      author: "Admin",
    };

    try {
      const res = await fetch("http://localhost:8080/blogs", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(newBlog),
      });

      if (!res.ok) {
        throw new Error("Failed to create blog");
      }

      const createdBlog = await res.json();
      setBlogs((prevBlogs) => [...prevBlogs, createdBlog]);
    } catch (error) {
      setError(error instanceof Error ? error.message : "Unknown error");
    } finally {
      setLoading(false);
    }
  };

  // ブログ削除の即時反映
  const handleDelete = async (blogId: number) => {
    setLoading(true);
    try {
      const res = await fetch(`http://localhost:8080/blogs?id=${blogId}`, {
        method: "DELETE",
      });

      if (!res.ok) {
        throw new Error("Failed to delete blog");
      }

      setBlogs((prevBlogs) => prevBlogs.filter((blog) => blog.id !== blogId));
    } catch (error) {
      setError(error instanceof Error ? error.message : "Unknown error");
    } finally {
      setLoading(false);
    }
  };

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div>
      <Header />
      <h1>Blogs</h1>
      <button
        onClick={handleCreate}
        disabled={loading}
        className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 disabled:opacity-50"
      >
        {loading ? "Creating..." : "Create Blog"}
      </button>
      {blogs.map((blog) => (
        <article key={blog.id} className="p-4 border rounded shadow">
          <h2 className="text-xl font-bold">{blog.title}</h2>
          <p>{blog.content}</p>
          <p className="text-sm text-gray-500">Author: {blog.author}</p>
          <time className="text-xs text-gray-400">
            {new Date(blog.createdAt).toLocaleDateString()}
          </time>
          <button
            onClick={() => handleDelete(blog.id)}
            disabled={loading}
            className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 disabled:opacity-50"
          >
            {loading ? "Deleting..." : "Delete"}
          </button>
        </article>
      ))}
      <Footer />
    </div>
  );
};

export default BlogList;
