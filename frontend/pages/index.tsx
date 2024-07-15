import { useEffect, useState } from 'react';
import Link from 'next/link';
import api from '../utils/api';

interface Post {
  id: number;
  title: string;
  published_at: string;
  updated_at: string;
}

const HomePage: React.FC = () => {
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await api.get('/posts');
        setPosts(response.data);
      } catch (error) {
        console.error('Error fetching posts:', error);
      }
    };

    fetchPosts();
  }, []);

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">Latest Posts</h1>
      <ul>
        {posts.map(post => (
          <li key={post.id} className="mb-2">
            <Link href={`/post/${post.id}`}>
              <a className="text-blue-500 hover:underline">{post.title}</a>
            </Link>
            <div>
              <span>Published: {post.published_at}</span>
              <span> | Updated: {post.updated_at}</span>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default HomePage;
