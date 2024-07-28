import { AppProps } from 'next/app';
import Layout from '../components/Layout';
import '../styles/globals.css';
import { useEffect, useState } from 'react';
import api from '../utils/api';
import { useRouter } from 'next/router';

const MyApp = ({ Component, pageProps }: AppProps) => {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
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
  }, []);

  useEffect(() => {
    const handleRouteChange = () => {
      setLoading(true);
      checkLoginStatus();
    };

    router.events.on('routeChangeStart', handleRouteChange);
    router.events.on('routeChangeComplete', () => setLoading(false));
    router.events.on('routeChangeError', () => setLoading(false));

    return () => {
      router.events.off('routeChangeStart', handleRouteChange);
      router.events.off('routeChangeComplete', () => setLoading(false));
      router.events.off('routeChangeError', () => setLoading(false));
    };
  }, [router]);

  return (
    <Layout isLoggedIn={isLoggedIn} setIsLoggedIn={setIsLoggedIn}>
      {loading ? <div>Loading...</div> : <Component {...pageProps} isLoggedIn={isLoggedIn} setIsLoggedIn={setIsLoggedIn} />}
    </Layout>
  );
};

export default MyApp;
