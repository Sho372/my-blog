import React from 'react';
import Link from 'next/link';

const NavBar: React.FC = () => {
  return (
    <nav className="bg-blue-500 p-4 text-white">
      <div className="container mx-auto flex justify-between">
        <Link href="/">
          <a className="text-xl font-bold">Blog</a>
        </Link>
        <div>
          <Link href="/register">
            <a className="mr-4">Register</a>
          </Link>
          <Link href="/login">
            <a>Login</a>
          </Link>
        </div>
      </div>
    </nav>
  );
};

export default NavBar;
