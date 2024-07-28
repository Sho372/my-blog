import { useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/router';
import api from '../utils/api';

interface LoginPageProps {
  isLoggedIn: boolean;
  setIsLoggedIn: (loggedIn: boolean) => void;
}

const LoginPage: React.FC<LoginPageProps> = ({ setIsLoggedIn }) => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await api.post('/login', {
        email,
        password,
      });
      console.log('User logged in:', response.data);
      // checkAuthを実行してログイン状態を更新
      const authResponse = await api.get('/check-auth');
      setIsLoggedIn(authResponse.data.authenticated);
      router.push('/');
    } catch (error) {
      console.error('Error logging in user:', error);
    }
  };

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">Login</h1>
      <form onSubmit={handleLogin}>
        <div className="mb-4">
          <label className="block mb-2">Email</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="border p-2 w-full"
          />
        </div>
        <div className="mb-4">
          <label className="block mb-2">Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="border p-2 w-full"
          />
        </div>
        <button type="submit" className="bg-blue-500 text-white p-2">Login</button>
      </form>
      <Link href="/">
        <a className="text-blue-500 hover:underline mt-4 inline-block">Go back to home</a>
      </Link>
    </div>
  );
};

export default LoginPage;
