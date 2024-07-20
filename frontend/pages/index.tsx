import { useEffect, useState } from 'react';
import Link from 'next/link';
import api from '../utils/api';
import { useRouter } from 'next/router';

interface Post {
  id: number;
  title: string;
  published_at: string;
  updated_at: string;
}

const HomePage: React.FC = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const router = useRouter();

  useEffect(() => {
    const checkLoginStatus = async () => {
      try {
        const response = await api.get('/check-auth');
        setIsLoggedIn(response.data.authenticated);
      } catch (error) {
        setIsLoggedIn(false);
      }
    };

    const fetchPosts = async () => {
      try {
        const response = await api.get('/posts');
        setPosts(response.data);
      } catch (error) {
        console.error('Error fetching posts:', error);
      }
    };

    checkLoginStatus();
    fetchPosts();

    // 5分ごとにログイン状態を確認
    const interval = setInterval(() => {
      checkLoginStatus();
    }, 300000); // 300,000ミリ秒 = 5分

    return () => clearInterval(interval);
  }, []);

  const handleLogout = async () => {
    try {
      await api.post('/logout');
      setIsLoggedIn(false);
      router.push('/');
    } catch (error) {
      console.error('Error logging out:', error);
    }
  };

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
      <div className="mt-4">
        {!isLoggedIn ? (
          <>
            <Link href="/login">
              <a className="text-blue-500 hover:underline mr-4">Login</a>
            </Link>
            <Link href="/register">
              <a className="text-blue-500 hover:underline">Register</a>
            </Link>
          </>
        ) : (
          <button onClick={handleLogout} className="text-blue-500 hover:underline">Logout</button>
        )}
      </div>
    </div>
  );
};

export default HomePage;
