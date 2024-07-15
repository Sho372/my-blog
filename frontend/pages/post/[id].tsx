import { useEffect, useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/router';
import api from '../../utils/api';

interface Post {
  id: number;
  title: string;
  content: string;
  published_at: string;
  updated_at: string;
}

const PostPage: React.FC = () => {
  const router = useRouter();
  const { id } = router.query;
  const [post, setPost] = useState<Post | null>(null);

  useEffect(() => {
    if (id) {
      const fetchPost = async () => {
        try {
          const response = await api.get(`/posts/${id}`);
          setPost(response.data);
        } catch (error) {
          console.error('Error fetching post:', error);
        }
      };

      fetchPost();
    }
  }, [id]);

  const handleEdit = () => {
    router.push(`/post/edit?id=${post?.id}`);
  };

  const handleDelete = async () => {
    try {
      await api.delete(`/posts/${post?.id}`);
      router.push('/');
    } catch (error) {
      console.error('Error deleting post:', error);
    }
  };

  if (!post) return <div>Loading...</div>;

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">{post.title}</h1>
      <div>
        <span>Published: {post.published_at}</span>
        <span> | Updated: {post.updated_at}</span>
      </div>
      <p className="mt-4">{post.content}</p>
      <button onClick={handleEdit} className="bg-yellow-500 text-white p-2 mr-2">Edit</button>
      <button onClick={handleDelete} className="bg-red-500 text-white p-2">Delete</button>
      <Link href="/">
        <a className="text-blue-500 hover:underline mt-4 inline-block">Go back to home</a>
      </Link>
    </div>
  );
};

export default PostPage;
