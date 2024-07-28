import { ReactNode, useEffect, useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/router';
import api from '../utils/api';

interface LayoutProps {
  children: ReactNode;
  isLoggedIn: boolean;
  setIsLoggedIn: (loggedIn: boolean) => void;
}

const Layout: React.FC<LayoutProps> = ({ children, isLoggedIn, setIsLoggedIn }) => {
  const [loading, setLoading] = useState<boolean>(true);
  const router = useRouter();

  const checkLoginStatus = async () => {
    try {
      const response = await api.get('/check-auth');
      setIsLoggedIn(response.data.authenticated);
    } catch (error) {
      setIsLoggedIn(false);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    checkLoginStatus();

    const interval = setInterval(() => {
      checkLoginStatus();
    }, 300000); // 5 minutes

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
      <header className="flex justify-between items-center py-4 bg-blue-500 text-white px-4">
        <h1 className="text-2xl font-bold">Blog</h1>
        <div>
          {!loading && (
            <>
              {!isLoggedIn ? (
                <>
                  <Link href="/register">
                    <a className="text-white hover:underline mr-4">Register</a>
                  </Link>
                  <Link href="/login">
                    <a className="text-white hover:underline">Login</a>
                  </Link>
                </>
              ) : (
                <>
                  <button onClick={handleLogout} className="text-white hover:underline mr-4">Logout</button>
                  <Link href="/post/new">
                    <a className="text-white hover:underline">Add New Post</a>
                  </Link>
                </>
              )}
            </>
          )}
        </div>
      </header>
      <main className="p-4">{loading ? <div>Loading...</div> : children}</main>
    </div>
  );
};

export default Layout;
